 package main

import "github.com/gin-gonic/gin"

type st struct {
	Name string `json:"mingzi"`
	Age  int    `json:"nianling"`
}

//postman输入{"mingzi":"jackk","nianling":27} //注意，27不要加引号，否则返回的对应位置的是0
//后，返回前端postman的是绑定结果是{jackk 27}{"mingzi":"jackk","nianling":27}
func main() {
	r := gin.Default()
	r.POST("/index", func(c *gin.Context) {
		var s st
		c.ShouldBindJSON(&s)
		c.String(200, "绑定结果是%v", s)
		c.JSON(200, s)
	})

	r.Run(":8888")
}
