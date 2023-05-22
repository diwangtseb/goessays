package main

import (
	"context"

	"github.com/apache/rocketmq-clients/golang"
	"github.com/apache/rocketmq-clients/golang/credentials"
)

func main() {
	producer, err := golang.NewProducer(&golang.Config{
		Endpoint: "localhost:9876",
		Credentials: &credentials.SessionCredentials{
			AccessKey:    "",
			AccessSecret: "",
		},
	})
	err = producer.Start()
	defer func() {
		err = producer.GracefulStop()
		panic(err)
	}()
	if err != nil {
		panic(err)
	}
	_, err = producer.Send(context.TODO(), &golang.Message{
		Topic: "hi",
		Body:  []byte{0x01},
		Tag:   new(string),
	})
	panic(err)
}
