package controllers

import (
	"github.com/kataras/iris/mvc"
	"imooc-iris/services"
	"imooc-iris/repositories"
	"github.com/kataras/iris"

)

type MovieController struct {
    Ctx  iris.Context


}

func (c *MovieController) Get() mvc.View {
	movieRepository:= repositories.NewMovieManager()
	movieService :=services.NewMovieServiceManger(movieRepository)
	MovieResult:=movieService.ShowMovieName()

	return mvc.View{
		Name:"movie/indexs.html",
		Data:MovieResult,
	}
	
}
