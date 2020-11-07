package main

//1、总结，通过url传递query查询参数，无tag时，只要键值对的键大小写与结构体字段完全相同shouldbind就
//可成功绑定该字段,否则不能成功绑定。
//todo:当然也可使用form类型tag标签，来和query参数的key保持一致，从而完成绑定【体会r.form和r.postform在官方
// 文档中的包含关系】但query参数和表单参数同时存在，tag都用form时，绑定会有问题。参见bang3ding4form文件夹下内容
//2、tag标签json，不是用来辅助shoulbind进行路径参数的绑定的。tag标签是json时，具体作用见bang3ding4json/main.go
import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type user struct {
	Name string

	PassWord string `json:"ps"` //验证发现json的tag标签不是用来绑定前端query路径
	               // 参数的；实际上它是专门针对前端传来json数据或后台传出json数据的
}

//对上面这种，发现即使不加任何tag,在浏览器输入框输入http://localhost:8888/login?Name=jack&PassWord=123
//ShouldBind也可以成功绑定，返回给前端的是{"Name":"jack","PassWord":"123"}，在cmd终端打印的是{jack 123}
//但浏览器输入http://localhost:8888/login?name=jack&PassWord=123
//就不行，返回给客户端的是{"Name":"","PassWord":"123"}    在cmd终端打印的是{  123}
//todo:可见，当url的query查询参数的key与结构体字段名完全一致时，不要任何tag标签ShouldBind也可以成功绑定参数

//当在PassWord字段加上`json:"ps"`，浏览器输入http://localhost:8888/login?Name=jack&ps=123
//返回的是{"Name":"jack","ps":""}，而不是{"Name":"jack","PassWord":""} 在cmd终端中打印{jack  }
//可见，当前端通过浏览器输入框发送query参数请求，json的tag标签并不用来辅助shouldbind绑定query参数【除非是前段发来json数
//据，才会体现参数绑定】，只是向前端用c.JSON返回数据时由于c.JSON是返回json类型，导致json的tag替代
//结构体字段,(可通过b.html来验证）即之前学的json用来改变前端只接受小写字段而后端结构体不大写没法被公开使用的问题

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./b.html")

	//在浏览器输入http://localhost:8888/login?Name=jack&PassWord=123  注意是Name、PassWord不是name、passWord
	r.GET("/login", func(c *gin.Context) {
		u := user{}
		err := c.ShouldBind(&u)
		//todo:这里有点奇怪，ShouldBind是根据请求体的Content-type来推断进行绑定的，
		// 但输入上面的网址发送请求并没有Content-type
		if err != nil {
			c.JSON(200, "c.ShouldBind获取客户端数据赋值进入结构体失败")
			return
		} else {
			fmt.Println("绑定后为u=", u)
			//c.JSON(200,u)
			c.HTML(200, "b.html", u)

		}
	})

	r.Run(":8888")
}
