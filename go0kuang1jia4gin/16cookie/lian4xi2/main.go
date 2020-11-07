package main

import "github.com/gin-gonic/gin"

func test(c *gin.Context) { //检查是否已经有正确cookie的中间价
	str, _ := c.Cookie("usercookie")
	if str != "choklate" {
		c.JSON(200, gin.H{
			"err":     "还没有正确的cookie",
			"message": "请访问/login",
		})
		c.Abort() //todo:这一句必须写，下面的return不能做到让中间件后面的处理器不运行
		return
	} else {
		c.Next()
	}
}
func main() {
	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		c.SetCookie("usercookie", "choklate", 60, "/",
			"localhost", false, true)
	})

	r.GET("/home", test, func(c *gin.Context) {
		c.String(200, "欢迎来到主页")
	})

	r.Run(":8888")
}
