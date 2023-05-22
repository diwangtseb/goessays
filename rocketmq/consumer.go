package rocketmq

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/apache/rocketmq-clients/golang"
	"github.com/apache/rocketmq-clients/golang/credentials"
	"github.com/rs/zerolog"
)

const DEFAULT_LENS = 16
const DEFAULT_GROUP_NAME = "default"
const DEFAULT_INVISIBLE_DURATION = time.Second * 20
const DEFAULT_AWAIT_DURATION = time.Second * 5

type HandleConsumerFunc[T any] func(context.Context, *RocketmqConsumer[T], *golang.MessageView) error

type RocketmqConsumer[T any] struct {
	endpoint     string
	accessKey    string
	accessSecret string
	consumer     golang.SimpleConsumer
	service      *T
	logger       *zerolog.Logger
	// maximum number of messages received at one time
	maxMessageNum int32
	//消息的最大处理时长 invisibleDuration should > 20s
	invisibleDuration time.Duration
	// await duration
	awaitDuration time.Duration
	// consumer group
	consumerGroup string
	// filter
	filters   []string
	mapHandle map[Topic]HandleConsumerFunc[T]
}

type ConsumerOpt[T any] func(*RocketmqConsumer[T])

func (co ConsumerOpt[T]) apply(rc *RocketmqConsumer[T]) {
	co(rc)
}

func WithMaxMessageNum[T any](lens int32) ConsumerOpt[T] {
	return func(rc *RocketmqConsumer[T]) {
		rc.maxMessageNum = lens
	}
}

func WithMaxDuration[T any](dur time.Duration) ConsumerOpt[T] {
	return func(rc *RocketmqConsumer[T]) {
		rc.invisibleDuration = dur
	}
}

func WithAwaitDurationn[T any](dur time.Duration) ConsumerOpt[T] {
	return func(rc *RocketmqConsumer[T]) {
		rc.awaitDuration = dur
	}
}

type Access struct {
	Endpoint     string
	AccessKey    string
	AccessSecret string
	Group        string
}

func toGolangFilterExpressions(src []string) map[string]*golang.FilterExpression {
	m := make(map[string]*golang.FilterExpression)
	for _, v := range src {
		m[v] = golang.SUB_ALL
	}
	return m
}

func NewRocketmqConsumer[T any](acc *Access, service *T, opts ...ConsumerOpt[T]) *RocketmqConsumer[T] {
	logger := LogBooter()
	rc := &RocketmqConsumer[T]{
		endpoint:          acc.Endpoint,
		accessKey:         acc.AccessKey,
		accessSecret:      acc.AccessSecret,
		service:           service,
		logger:            &logger,
		maxMessageNum:     DEFAULT_LENS,
		invisibleDuration: DEFAULT_INVISIBLE_DURATION,
		awaitDuration:     DEFAULT_AWAIT_DURATION,
		consumerGroup:     acc.Group,
		filters:           []string{},
		mapHandle:         map[Topic]HandleConsumerFunc[T]{},
	}

	for _, opt := range opts {
		opt.apply(rc)
	}
	return rc
}

// must pay attention to!! if err == nil msg will be ack else error try
func (rc *RocketmqConsumer[T]) Regist(topic Topic, handle HandleConsumerFunc[T]) {
	_, ok := rc.mapHandle[topic]
	if ok {
		return
	}
	rc.mapHandle[topic] = handle
}

func (rc *RocketmqConsumer[T]) handleNormal(ctx context.Context) {
	messageViews, err := rc.consumer.Receive(context.TODO(), rc.maxMessageNum, rc.invisibleDuration)
	if err != nil {
		rc.logger.Err(err).Any("messageViews", messageViews)
	}
	for _, msg := range messageViews {
		topic := msg.GetTopic()
		handle, ok := rc.mapHandle[Topic(topic)]
		if !ok {
			continue
		}
		err := handle(ctx, rc, msg)
		if err != nil {
			rc.logger.Err(err).Any(topic, msg)
			continue
		}

		err = rc.consumer.Ack(ctx, msg)
		if err != nil {
			rc.logger.Err(err)
		}
	}
}

func (rc *RocketmqConsumer[T]) Start() {
	os.Setenv("mq.consoleAppender.enabled", "true")
	golang.ResetLogger()
	for k := range rc.mapHandle {
		rc.filters = append(rc.filters, string(k))
	}
	simpleConsumer, err := golang.NewSimpleConsumer(&golang.Config{
		Endpoint:      rc.endpoint,
		ConsumerGroup: rc.consumerGroup,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    rc.accessKey,
			AccessSecret: rc.accessSecret,
		},
	},
		//设置获取消息时间
		golang.WithAwaitDuration(rc.awaitDuration),
		//设置订阅关系
		golang.WithSubscriptionExpressions(toGolangFilterExpressions(rc.filters)),
	)
	if err != nil {
		panic(err)
	}
	rc.consumer = simpleConsumer
	err = rc.consumer.Start()
	defer rc.close()
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			rc.handleNormal(context.Background())
		}
	}()
}

func (rc *RocketmqConsumer[T]) close() {
	systemSignal := make(chan os.Signal, 1)
	signal.Notify(systemSignal, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-systemSignal
		rc.logger.Info().Msg("secure close")
		rc.Close()
	}()
}

// notice: you can be explicitly turned off by business,default: start defer close()
func (rc *RocketmqConsumer[T]) Close() {
	err := rc.consumer.GracefulStop()
	if err != nil {
		rc.logger.Err(err).Str("rocketmq", "consumer")
	}
}
