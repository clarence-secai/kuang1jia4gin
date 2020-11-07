package main
import "github.com/gin-gonic/gin"

//todo:经验证明填写表单这种方式必须有form的tag标签
type stu struct {
	Name     string `form:"mingzi"` //填写表单向后台发送数据，数据是在请求体中，故请求体的content-type须是multipart/form-data
	PassWord string `form:"mima"`   //经验证明填写表单这种方式必须有form的tag标签，
	// (即使HTML中name=""里填这里字段大小写相同的也不行)
	// 借助form的tag标签，c.ShouldBind可以成功根据tag将表单中填的数据绑定到这里的结构体字段上，
	//返回前端{"Name":"tom","PassWord":"111"}
	Img string    //todo:这里即使加上`form:"Img"`由于Img是query查询参数，故下面运行依然无法成功绑定该字段
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./a.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(200, "a.html", nil)
	})
	r.POST("/login", func(c *gin.Context) {
		u := stu{}              //申明一个结构体变量，下一行就把客户端传来的数据放进来
		err := c.ShouldBind(&u) //这里传进来的u必须有&
		//才能获取c中的参数、用ShouldBind的反射机制去比对form的tag标签来匹配到相应的字段，
		//并将前端数据赋值给相应字段，使得接下来的u不再是申明时候的初始默认值，而是拿到了客户端的值，

		//todo:但尤为要注意的是，前端写表单传数据的情况下，shouldbind绑定只能取到客户端发送的请求体（即表单）中的值，
		// 不能同时取到浏览器输入框方式路径上query的参数，因为ShouldBind是根据请求头的Content-type来推断进行绑定的
		// 而本例中，运行起来前端请求体的Content-Type: multipart/form-data; 因此只会绑定表单字段而不会绑定查询参数
		// 本例中如果在前端表单里填入tom和413188ok提交后，浏览器输入框中是http://localhost:8888/login?Img=lala
		// 返回前端的结果是{"Name":"tom","PassWord":"413188ok"，"Img":""}，即没取到query路径参数

		if err != nil {
			c.JSON(200, "c.ShouldBind获取客户端数据赋值进入结构体失败")
			return
		} else {
			c.JSON(200, u)
		}
	})

	r.Run(":8888")
}
