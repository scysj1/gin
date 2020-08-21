package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type HelloController struct {

}

func(hello *HelloController)Router(engine *gin.Engine){
	engine.GET("/hello",hello.Hello)
}


	//接口
 func (hello *HelloController) Hello(context *gin.Context){
		fmt.Println("Hello")
		context.Writer.WriteString("Hello")
 }
