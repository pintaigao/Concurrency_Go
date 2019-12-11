package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"imooc-iris/web/controllers"
)

func main()  {
	//1.创建实例
	app:=iris.New()
	//2.注册控制器
	mvc.New(app.Party("/hello")).Handle(new(controllers.MovieController))
	//3.启动服务
	app.Run(
		iris.Addr("localhost:8080"),
	)
}
