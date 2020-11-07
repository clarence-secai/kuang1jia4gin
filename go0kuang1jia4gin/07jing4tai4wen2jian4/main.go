 package main

import "github.com/gin-gonic/gin"

func handler1(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
func main() {
	r := gin.Default()

	//处理静态文件  具体见七米的P9视频25:00以后
	r.Static("/xxx", "./mu2ban3") //替换为的路径以当前main.go为起始参照
	//relativePath的"/xxx"是index.html等静态文件中src和href等中的路径开头，
	//root是将这些路径的开头/xxx替换为./mu2ban3，后续路径不变,类似于官方文档中原来
	//的http.Handle("/tmpfiles/", http.StripPrefix("/tmpfiles/", http.FileServer(http.Dir("/tmp"))))

	//todo:解析模板文件
	r.LoadHTMLFiles("./mu2ban3wen2jian4/index.html")
	//todo:处理器中执行模板文件
	r.GET("/index", handler1)
	r.Run(":8888")
}
