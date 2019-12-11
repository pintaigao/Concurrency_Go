package main

import "rabbitmq/RabbitMQ"

func main()  {
	imoocOne:=RabbitMQ.NewRabbitMQTopic("exImoocTopic","#")
	imoocOne.RecieveTopic()
}
