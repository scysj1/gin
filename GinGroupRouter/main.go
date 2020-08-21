package main



import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)
type login struct{
	Username string `form:"username"`
	Password string `form:"password"`
}
func RequestInfos() gin.HandlerFunc{
	return func(context *gin.Context){
		path:=context.FullPath()
		method:=context.Request.Method
		fmt.Println("请求Path：",path)
		fmt.Println("请求Path：",method)
		context.Next()//context.Next函数可延迟后面代码执行，可以在请求处理完成后再执行后面代码、
		fmt.Println("状态码",context.Writer.Status())

	}
}
//路由组分类
func main(){

	engine :=gin.Default()
	//engine.Use(RequestInfos())
	routeruser :=engine.Group("/user")
	routeruser.POST("/login",RequestInfos(),func(context *gin.Context){
		fmt.Println("中间键分隔")
		var user login
		if err :=context.ShouldBind(&user);err!=nil{
			log.Fatal(err.Error())
			return
		}
		fmt.Println("username:"+user.Username)
		fmt.Println("password:"+user.Password)
		//context.Writer.WriteString("用户名："+user.Username+"\n密码："+user.Password)
		context.JSON(200,user)
	})

	engine.Run(":8081")


}
