package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main(){
	engine :=gin.Default()
	engine.GET("/Video",func(context *gin.Context){
		fmt.Println("路径：",context.FullPath())
		context.Writer.Write([]byte("Hello, gin\n"))



	})
	if err :=engine.Run(":8091"); err!=nil{
		log.Fatal(err.Error())
	}
}
