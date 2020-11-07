package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//head请求，一个像一般的get请求post请求等那样的http请求
	r.HEAD("/index1", func(c *gin.Context) {
		c.JSON(200, "HEAD请求")
	})
	//可同时使用下面这种用一个路径应对多种请求方式,但注意head请求里的relativepath
	//和Any里的不能一样，否则会报错，因为重复了，服务器不知道到底该用哪一个，会掐架
	r.Any("/index", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet: //注意这样写的话不能加引号
			c.JSON(200, "get请求")
		case "POST":
			c.JSON(200, "post请求")
		case "DELETE":
			c.JSON(http.StatusOK, "delete请求")
		default: //这里的default未能包含所有其他未在上面列出的请求，
			// 比如copy请求等会返回404而不是“未知请求“四个字
			c.JSON(200, "未知请求")
		}

	})
	r.NoRoute(func(c *gin.Context) { //todo:针对客户端输入的路径是本服务器内不存在的情况
		c.JSON(404, "您要找的网页不存在")
	})
	//此处状态码没有硬性要求，填200也可，但不建议
	r.Run(":8888")
}
