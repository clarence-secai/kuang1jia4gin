package main

//todo:
// 在浏览器输入   http://localhost:8888/index/jack/27/33.3
// 返回的是 绑定uri的结果是{  0 0}  因此，ShouldBind没有绑定uri上的路径param参数
// 的功能（tag改成form也不行），要想绑定，必须用ShouldBindUri方法,必须有uri的tag标签

import "github.com/gin-gonic/gin"

type people struct {
	Name  string  `uri:"a" ` //这里uri的tag的内容对应着下面的  "/index/:a/:b/:c"
	Age   int     `uri:"b" `
	Score float64 `uri:"c" `
}

func main() {
	r := gin.Default()
	r.GET("/index/:a/:b/:c", func(c *gin.Context) {
		var x people
		c.ShouldBind(&x)
		//todo:发现改成c.ShouldBindUri的才可成功了,因为ShouldBind是根据请求头的Content-Type来推断的
		// 而单纯的路径param参数的请求中，请求头内没有Content-Type，故无法智能地推断绑定成功。
		c.String(200, "绑定uri的结果是%v", x)
	})

	r.Run(":8888")
}
