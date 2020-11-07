package main

import "github.com/gin-gonic/gin"

//下面是关于中间件的嵌套的执行顺序的 详情也可见七米视屏的P18 29:55
func AllMidHan1(c *gin.Context) {
	c.JSON(200, "我是后续总中间件1开始")
	c.Next() //调用后续处理器，这个后续处理器包括中间件，因为中间件本质也是处理器
	c.JSON(200, "我是后续总中间件1开始结束")
}
func AllMidHan2(c *gin.Context) {
	c.JSON(200, "我是后续总中间件2开始")
	c.Next() //调用后续处理器，这个后续处理器包括中间件，因为中间件本质也是处理器
	c.JSON(200, "我是后续总中间件2开始结束")
}
func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		c.JSON(200, "我是匿名处理器")
	})
	//上面一行输出的是  “我是匿名处理器” 不会有中间件，因为中间件在下面才开始注册

	r.Use(AllMidHan1, AllMidHan2)
	r.GET("/login", func(c *gin.Context) {
		c.JSON(200, "我是匿名处理器2")
	})
	r.Run(":8888")
	/*
	"我是后续总中间件1开始"
	"我是后续总中间件2开始"
	"我是匿名处理器2"
	"我是后续总中间件2开始结束"
	"我是后续总中间件1开始结束"
	*/
}


