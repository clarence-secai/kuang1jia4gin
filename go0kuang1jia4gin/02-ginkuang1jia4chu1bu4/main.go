//todo:gin框架和文档中原始方式不能混用，gin框架本身兼容文档中的原始方式
// 以下的gin框架的代码全部正常运行
package main

import (
	"github.com/gin-gonic/gin"
)
type stu struct {
	Name string
	Age  int64
}
type pest struct {
	Color  string
	Weight int64
}

func myhandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"neirong": "haha",
	})
}

func main() {
	engine := gin.Default()
	engine.GET("/index", func(c *gin.Context) {
		c.JSON(200, "GET的内容")
	})
	engine.POST("/index", myhandler)//也可以这样

	engine.PUT("/index", func(context *gin.Context) {
		context.JSON(200, gin.H{"key": "value from PUT"})
	})
	engine.GET("/index2", func(c *gin.Context) {
		s1 := stu{Name: "jack", Age: 27}
		p1 := pest{Color: "black", Weight: 20}
		m := map[string]interface{}{"no1": s1, "no2": p1}
		c.JSON(200,m)
	})

	engine.LoadHTMLFiles("./a.html")
	engine.GET("/index3", func(c *gin.Context) {
		s1 := stu{Name: "jack", Age: 27}
		p1 := pest{Color: "black", Weight: 20}
		m := map[string]interface{}{"no1": s1, "no2": p1}
		c.HTML(200,"a.html",m)
	})
	engine.Run(":9999")


	//todo:上下两个监听不能同时都运行，否则虽不会报错，但程序卡在前一个监
	// 听，后一个监听不会运行并且，由于两个监听不能都运行，因此各自的处理器
	// 也没法同时工作，即实际上不能混用

	////下面展示的是如何一次向前端发送多个结构体变量
	//stu1 := stu{Name: "jack", Age: 27}
	//pest1 := pest{Color: "black", Weight: 20}
	//me := map[string]interface{}{"no1": stu1, "no2": pest1}
	////前端在调用时，不是{{."no1".Name}},而是{{.no1.Name}}
	//http.HandleFunc("/family", func(rw http.ResponseWriter, r *http.Request) {
	//	t, _ := template.ParseFiles("./a.html")
	//	t.Execute(rw, me)
	//})
	//http.ListenAndServe(":8888", nil)


}
