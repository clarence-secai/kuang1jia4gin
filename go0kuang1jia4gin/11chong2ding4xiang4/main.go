package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//301是永久重定向，也就是浏览器不会把最终出现在页面当做是/index的
	r.GET("/index", func(c *gin.Context) {
		c.Redirect(301, "https://www.guancha.cn/")
		//todo:这里的状态码必须与实际情况匹配，不能随便填其他不匹配的状态码如200
	})

	//下面的方式是请求的转发，两个r.GET都运行了
	r.GET("/a", func(c *gin.Context) {
		c.Request.URL.Path = "/b" //这个/b就对应下面的r.GET("/b",...
		r.HandleContext(c)        //由于这里用到r,因此这种情况下这里的r.GET("/a",...里的处理器只能用匿名的
	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(200, "重定向成功")
	})

	r.Run(":8888")
}
