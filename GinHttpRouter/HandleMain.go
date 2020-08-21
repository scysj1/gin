package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)
//get
type User struct{
	Username string `form:"username" binding:"required"`
	Phone string `form:"phone" binding:"required"`
	Password string `form:"password" binding:"required"`
}
type Userp struct{
	Username string `form:"username"`
	Phone int `form:"phone"`
	Password string `form:"password" `
}
func main(){
	engine :=gin.Default()

	//hello接口
	//http://localhost:8080/hello?name=ShiyangJiao
	engine.Handle("GET","/hello",func(context *gin.Context){
		path :=context.FullPath()
		fmt.Println(path)
		name :=context.DefaultQuery("name","hello")
		fmt.Println("hello, "+name)
	})
	//Video接口
	engine.GET("/Video",func(context *gin.Context) {
		fmt.Println("路径：", context.FullPath())
		context.Writer.Write([]byte("Hello, gin\n"))
	})
	//qurey接口
	engine.Handle("POST","/qurey",func(context *gin.Context){
		path :=context.FullPath()
		fmt.Println(path)
		username :=context.PostForm("username")
		password :=context.PostForm("Password")
		fmt.Println("username: "+username)
		fmt.Println("password: "+password)
	})

	//** 数据绑定**//
	//http://localhost:8080/s1?username=david&password=jsy556699&phone=13300000000
	engine.GET("/s1",func(context *gin.Context){
		fmt.Println(context.FullPath())
		var user User
		err :=context.ShouldBindQuery(&user)
		if err !=nil{
			log.Fatal(err.Error())
		}
		fmt.Println(user.Username)
		fmt.Println(user.Password)
		fmt.Println(user.Phone)
		context.Writer.Write([]byte("hello,"+user.Username))
	})

	//http://localhost:8080/s2
	//POST从表单获取参数
	engine.POST("/s2",func(context *gin.Context){
		fmt.Println(context.FullPath())
		var user Userp
		if err :=context.ShouldBind(&user) ; err!=nil{
			log.Fatal(err.Error())
			return
		}
		fmt.Println(user.Username)
		fmt.Println(user.Password)
		fmt.Println(user.Password)
		context.Writer.Write([]byte(user.Username+"login"))
	})
	//POST从json获取参数
	engine.POST("/s3",func(context *gin.Context){
		fmt.Println(context.FullPath())
		var user Userp
		if err :=context.BindJSON(&user) ; err!=nil{
			log.Fatal(err.Error())
			return
		}
		fmt.Println(user.Username)
		fmt.Println(user.Password)
		fmt.Println(user.Phone)
		context.Writer.Write([]byte(user.Username+"login"))
	})

	//**数据返回**//
	//write返回String
	engine.GET("/return1",func(context *gin.Context){
		fullpath :="路径"+context.FullPath()
		fmt.Println(fullpath)
		context.Writer.Write([]byte(fullpath))
	})

	//writestring返回string
	engine.GET("/return",func(context *gin.Context){
		fullpath :="路径"+context.FullPath()
		fmt.Println(fullpath)
		context.Writer.WriteString(fullpath)
	})

	engine.GET("/returnjson",func(context *gin.Context){
		fullpath :="路径:"+context.FullPath()
		fmt.Println(fullpath)
		context.JSON(200,map[string]interface{}{
			"code":1,
			"message":"ok",
			"data":fullpath,
		})
	})
	//结构体返回
	engine.GET("/returnjsonstruct",func(context *gin.Context){
		fullpath :="路径:"+context.FullPath()
		fmt.Println(fullpath)
		resp := result{Code:1,Message:"OK",Data:fullpath}
		context.JSON(200,resp)

	})


	engine.GET("/return/html",func(context *gin.Context){
		fullpath :="路径:"+context.FullPath()
		fmt.Println(fullpath)


	})




	engine.Run()
}
type result struct {
	Code int
	Message string
	Data interface{}

}
