package RabbitMQ

import "github.com/streadway/amqp"

const MQURL = "amqp://imoocuser:imoocuser@192.168.2.10:5672/imooc"

type RabbitMQ struct {
	// conn 是变量名，*amqp.Connection 是类型
	conn    *amqp.Connection
	channel *amqp.Channel
}
