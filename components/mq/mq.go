package mq

import (
	"github.com/google/uuid"
	"github.com/nsqio/go-nsq"
	"github.com/xlizy/common-go/base/const/threadlocal"
	"github.com/xlizy/common-go/utils/crypto"
	"github.com/xlizy/common-go/utils/zlog"
)

type RootConfig struct {
	Config nsqConfig `yaml:"nsq"`
}

type nsqConfig struct {
	NSQLookupdAddr     string `yaml:"lookupdAddr"`
	NSQDAddr           string `yaml:"nsqdAddr"`
	DefaultTopicPrefix string `yaml:"defaultTopicPrefix"`
}

type Consumer struct {
	Topic       string
	Channel     string
	Concurrency int
	Handler     func(msg string) error
}

var _msgKey = "e7uvNbck5NPSS6z0iglw"
var _producer *nsq.Producer
var _rc *RootConfig

type MsgHandler struct {
	Handler func(msg string) error
}

func (h MsgHandler) HandleMessage(message *nsq.Message) error {
	if len(message.Body) == 0 {
		return nil
	}
	threadlocal.SetTraceId(uuid.New().String())
	msg := string(message.Body)
	msg = crypto.AesDecryptECB(msg, _msgKey)
	zlog.Info("接收到MQ消息:{}", msg)
	return h.Handler(msg)
}

func NewConfig() *RootConfig {
	return &RootConfig{}
}

func InitNsq(rc *RootConfig) {
	_rc = rc
	if _rc.Config.DefaultTopicPrefix == "" {
		_rc.Config.DefaultTopicPrefix = "default"
	}
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(rc.Config.NSQDAddr, config)
	if err != nil {
		zlog.Error("创建NSQ生产端异常:{}", err.Error())
		panic(err)
	} else {
		_producer = producer
	}
}

func BuildConsumer(consumers []Consumer) {
	for _, consumer := range consumers {
		config := nsq.NewConfig()
		if consumer.Channel == "" {
			consumer.Channel = "default"
		}
		c, err := nsq.NewConsumer(_rc.Config.DefaultTopicPrefix+"-"+consumer.Topic, consumer.Channel, config)
		if err != nil {
			zlog.Error("创建nsq消费端异常:Topic={},Channel={},{}", consumer.Topic, consumer.Channel, err.Error())
		}
		if consumer.Concurrency == 0 {
			consumer.Concurrency = 1
		}
		c.AddConcurrentHandlers(&MsgHandler{Handler: consumer.Handler}, consumer.Concurrency)
		err = c.ConnectToNSQLookupd(_rc.Config.NSQLookupdAddr)
		if err != nil {
			zlog.Error("连接NSQLookupd异常:{}", err.Error())
		}
	}
}

func SendMsg(topic, msg string) error {
	zlog.Info("发送到MQ消息:{}", msg)
	msg = crypto.AesEncryptECB(msg, _msgKey)
	err := _producer.Publish(_rc.Config.DefaultTopicPrefix+"-"+topic, []byte(msg))
	if err != nil {
		zlog.Error("nsq推送消息异常:{}", err.Error())
		return err
	} else {
		return nil
	}
}
