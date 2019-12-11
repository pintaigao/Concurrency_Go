package main

import (
	"github.com/kataras/iris"
)

func main() {
	//1.创建iris 实例
	app := iris.New()

	//2.设置模板
	app.StaticWeb("/public", "./fronted/web/public")
	//3.访问生成好的html静态文件
	app.StaticWeb("/html", "./fronted/web/htmlProductShow")

	app.Run(
		iris.Addr("0.0.0.0:80"),
		iris.WithoutVersionChecker,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
