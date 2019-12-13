package RabbitMQ

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

const MQURL = "amqp://imoocuser:imoocuser@192.168.2.10:5672/imooc"

type RabbitMQ struct {
	// conn 是变量名，*amqp.Connection 是类型
	conn    *amqp.Connection
	channel *amqp.Channel

	// 队列名称
	QueueName string
	// 交换机
	Exchange string
	// key
	key string
	// 连接信息
	Mqurl string
}

// 创建RabbitMQ的实例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	return &RabbitMQ{QueueName: queueName, Exchange: exchange, key: key, Mqurl: MQURL}
}

// 类方法，相当于class中的方法，目的是断开channel和connecttion
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

//错误处理
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

// 创建简单模式下RabbitMQ实例
func NewRabbitMQSimple(queuName string) *RabbitMQ {
	rabbitmq := NewRabbitMQ(queuName, "", "")
	var err error
	// 创建rabbitmq连接
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "创建连接错误！")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "获取channel失败")
	return rabbitmq
}
