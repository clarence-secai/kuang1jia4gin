package main

import "github.com/gin-gonic/gin"

func handler(c *gin.Context) {
	c.JSON(200, "haha")//obj其实没有特定的类型要求,都会被c.JSON序列化后发给前端
}
func main() {
	r := gin.Default()
	r.GET("/json1", func(c *gin.Context) {
		c.JSON(200, gin.H{"name": "tom"})
		//gin.H这个gin框架里定义好的类型，其本质是map[string]interface{}
	})

	r.GET("/json2", handler)

	type stu struct {
		Name    string
		Age     int64
		Message string `json:"message"` //这样就可在以c.JSON方式发送到客户端时显示小写message，否则
		//如直接小写因不可导出导致在客户端不显示message这个字段
	}
	r.GET("/json3", func(c *gin.Context) {
		c.JSON(200, stu{Name: "jack", Age: 27, Message: "努力学习"})
	}) //c.JSON()会将第二个参数obj序列化后返回前端，这里返回
	  // 的是{"Name":"jack","Age":27,"message":"努力学习"}message字段是小写的
	r.Run(":8888")
}
