package rocketmq

import (
	"context"
	"os"

	"github.com/apache/rocketmq-clients/golang"
	"github.com/apache/rocketmq-clients/golang/credentials"
	"github.com/bytedance/sonic"
	"github.com/lithammer/shortuuid"
	"github.com/rs/zerolog"
)

type RocketmqProducer struct {
	producer golang.Producer
	logger   *zerolog.Logger
}

func NewRocketmqProducer(endpoint string, accessKey string, accessSecret string) *RocketmqProducer {
	os.Setenv("mq.consoleAppender.enabled", "true")
	golang.ResetLogger()
	producer, err := golang.NewProducer(&golang.Config{
		Endpoint: endpoint,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    accessKey,
			AccessSecret: accessSecret,
		},
	})
	if err != nil {
		panic(err)
	}

	logger := LogBooter()
	return &RocketmqProducer{
		producer: producer,
		logger:   &logger,
	}
}

func (rocket *RocketmqProducer) Start() {
	if err := rocket.producer.Start(); err != nil {
		panic(err)
	}
}

type SendParam struct {
	Topic Topic
	Key   []string
	Tag   string
	Param interface{}
}

func hasKey(keys ...string) bool {
	return len(keys) > 0
}
func generateUuid() string {
	return shortuuid.New()
}

func (rocket *RocketmqProducer) Send(ctx context.Context, param SendParam) error {
	defer rocket.logger.Info().Any("param", param)
	body, err := sonic.Marshal(param.Param)
	if err != nil {
		rocket.logger.Err(err)
		return err
	}
	msg := &golang.Message{
		Topic: string(param.Topic),
		Body:  body,
		Tag:   &param.Tag,
	}
	if !hasKey(param.Key...) {
		param.Key = []string{generateUuid()}
	}
	msg.SetKeys(param.Key...)
	_, err = rocket.producer.Send(ctx, msg)
	if err != nil {
		rocket.logger.Err(err)
		return err
	}
	return nil
}

func (rocket *RocketmqProducer) Close() {
	if err := rocket.producer.GracefulStop(); err != nil {
		rocket.logger.Err(err)
	}
}
