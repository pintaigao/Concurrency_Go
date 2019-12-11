package main

import "rabbitmq/RabbitMQ"

func main()  {
	imoocOne:=RabbitMQ.NewRabbitMQRouting("exImooc","imooc_two")
	imoocOne.RecieveRouting()
}
