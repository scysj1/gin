package main

import (
	"GIn_learn/Projectlearn/controller"
	"GIn_learn/Projectlearn/tool"
	"github.com/gin-gonic/gin"
)

func main(){

	cfg,err := tool.ParseConfig("./config/app.json")
	if err !=nil{
		panic(err.Error())
	}
	app :=gin.Default()
	rigisterRouter(app)
	app.Run(cfg.AppHost+":"+cfg.AppPort)

}

func rigisterRouter(router *gin.Engine){
	new(controller.HelloController).Router(router)
}
