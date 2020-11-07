//todo:gin.Context下的Param方法，是用于获取url上的路径参数

package main
import "github.com/gin-gonic/gin"


func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")

	//在浏览器输入栏输入http://localhost:8888/index/2020/04就会显示相应信息
	//注意，输入栏输入的信息和顺序必须和"/index/:year/:month"一一对应且不多不少，否则出错

	//"/index/:year/:month"的  :year  :month相当于服务端对浏览器输入路径参数的相应位置在这里进行的命名
	r.GET("/index/:year/:month", func(c *gin.Context) {
		nian := c.Param("year")
		//鼠标悬浮Param：Param returns the value of the URL param. It is a shortcut for c.Params.ByName(key)
		yue := c.Param("month")
		c.HTML(200, "index.html", gin.H{"str1": nian, "str2": yue})
	})

	//注意，不能像下面这样又写一个，启动运行会报错，因为当在浏览器输入时，下面这个包含了上面的情况，
	//解决方法是把下面的路径改一改，避免两个存在包含关系，比如改为"/index2/:name/:age
	//这样就可以对应客户端输入的url以/index打头的/index/**/**模式。详见七米视频的p13的08:58处
	//r.GET("/:name/:age", func(c *gin.Context) {
	//	name := c.Param("name")
	//	age := c.Param("age")
	//	c.HTML(200, "index.html", gin.H{"str1": name, "str2": age})
	//})

	//在浏览器输入栏输入http://localhost:8888/index3/jack/mary就会显示相应信息
	r.GET("index3/:num1/:num2", func(c *gin.Context) {
		p1:= c.Param("num1")
		p2:= c.Param("num2")
		c.HTML(200,"index.html",gin.H{"str1":p1,"str2":p2}) //前端会显示   年：jack 月:mary
	})
	r.Run(":8888")
}
