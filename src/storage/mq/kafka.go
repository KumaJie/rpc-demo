package mq

import (
	"errors"
	"fmt"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
	"rpc-douyin/src/config"
	"rpc-douyin/src/util/log"
)

var host string

func init() {
	host = fmt.Sprintf("%s:%d", config.Cfg.Kafka.Host, config.Cfg.Kafka.Port)
}

func NewSyncProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true
	client, err := sarama.NewSyncProducer([]string{host}, config)
	if err != nil {
		log.Error("Kafka: create sync producer failed", zap.Error(err))
		return nil, err
	}
	return client, nil
}

func NewAsyncProducer() (sarama.AsyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	client, err := sarama.NewAsyncProducer([]string{host}, config)
	if err != nil {
		log.Error("Kafka: create async producer failed", zap.Error(err))
		return nil, err
	}
	return client, nil
}

func NewConsumer() (sarama.Consumer, error) {
	client, err := sarama.NewConsumer([]string{host}, nil)
	if err != nil {
		log.Error("Kafka: create consumer failed", zap.Error(err))
		return nil, err
	}
	return client, nil
}

func NewMessage(topic string, value interface{}) (*sarama.ProducerMessage, error) {
	msg := sarama.ProducerMessage{
		Topic: topic,
	}
	switch value.(type) {
	case string:
		msg.Value = sarama.StringEncoder(value.(string))
	case []byte:
		msg.Value = sarama.ByteEncoder(value.([]byte))
	default:
		return nil, errors.New("unsupport message type")
	}
	return &msg, nil
}
