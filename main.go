package main
import (
	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context)  {
	c.JSON(200,gin.H{
		"message":"helo Golang",
	})
}
func main() {
	r:=gin.Default()//返回默认的路由引擎
	//指定当用户使用Get请求访问/hello时，执行sayHello这个函数
	r.GET("/hello",sayHello)
	r.Run(":9090")
}
