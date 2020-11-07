package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	//每次在一个新的project中引入gin(后面的gorm)等第三方包，都需要执行一次go mod tidy
)

//这个相当于处理器函数的，就是在执行模板
func handler1(c *gin.Context) {
	//下面一行的name是填写模板名（无define时默认是带所属文件夹的文件名，有define的用自定义模板名，
	//遇到不同文件夹下重名文件的，使用define给文件自定义模板文件名），而非相对路径文件名，
	c.HTML(http.StatusOK, "other a", gin.H{
		"title": "<a href='https://www.bilibili.com/'>b站</a>",
	})
}

type myText struct{
	Title string `json:"title" html:"title"` //todo:有html这个标签嘛？咋没作用？
	Txt string `json:"txt" html:"txt"`
}
func handler0(c *gin.Context){
	m := myText{
		Title: "<a href='https://www.bilibili.com/'>b站</a>",
		Txt: "haha",
	}
	c.HTML(200,"c.html",m)//todo:传给前端一个结构体也可以，
									// 但前端的点字段对应地也必须大写，否则不显示
}
func handler11(c *gin.Context) {
	//下面一行的name是填写模板名【而非带路径的文件名】无define时默认是带所属文件夹的文件名，有define的用自定义模板名，
	//遇到不同文件夹下重名文件的，两个文件都要使用define给文件自定义模板文件名，否则前端浏览器不会因为两个重名文件中的
	//一个有define而区分开。
	c.HTML(http.StatusOK, "a.html", gin.H{//todo:由于这个文件名重名，故在调用此处理器时浏览器显示空白页
		"title": "<a href='https://www.bilibili.com/'>b站</a>",
	})
}
func handler2(c *gin.Context) {
	c.HTML(200, "b.html", gin.H{
		"title": "<a href='https://www.bilibili.com/'>b站</a>"})
}
func hanshu(s string) template.HTML {
	return template.HTML(s)//todo:即解除html/template自带的安全模式
}
func main() {
	r := gin.Default()
	//todo:添加自定义函数；如需要，则必须在解析模板之前添加Funcs自定义函数
	r.SetFuncMap(template.FuncMap{"unsafe": hanshu})
	//上句相当于官方文档的t.Funcs(template.FuncMap{"zidingyi":kua})
	//不同之处是这里注册的下面都可默认用，官方文档的需每个handler注册一回

	//todo:解析模板文件r.LoadHTMLFiles("./a.html")
	// 一次解析多个模板文件
	r.LoadHTMLGlob("templatemu2ban3/**/*") //todo:两个星表示其下所有文件夹，一个星表示其下所有文件，这里就解析了template文件夹下所有文件

	r.GET("/index1", handler1) //相当于http.HandlerFunc一样的注册处理器函数
	r.GET("index0",handler0)
	r.GET("/index11", handler11)
	r.GET("/wenjian/index2", handler2) //可以是/wenjian/index2这样的

	r.Run(":8888")
}
