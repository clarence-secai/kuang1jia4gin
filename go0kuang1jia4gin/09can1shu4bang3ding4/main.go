//todo:请前往查看其它几个文件夹的绑定参数的示例和解释。这里说的是一个奇怪现象

package main

import "github.com/gin-gonic/gin"

//为避免每次获取客户端参数赋值给结构体这个操作，下面采用参数绑定的函数的方法 参加七米视频P14 15:10
//ShouldBind的反射机制是将c中包含的前端数据的键（键值对的键）去比对tag标签，比对发现一模一样的就把
//前端的值（键值对的值）的数据赋值到相应字段，使得接下来的u中的相应字段拿到了客户端的值

type user struct {
	Name string `form:"mingzi"`
	//这个form就表示对应客户端表单数据之意（其他tag还有json,binding等，详见七米博客）
	//用于下面shouldbind函数内置的反射，c.shouldbind就是
	//将c传来的表单参数的键值对的键去和tag比对，比对发现是一样的就把键值对的值匹赋值给对应的字段
	//这里form冒号后面写的就是和客户端表单里每个input里的name=“mingzi”引号内的值是一模一样对应的

	PassWord string `form:"mima"` //同时写上其他tag书写格式就是   `form:"mima"  json:"password"`
	// json的tag作用是在c.JSON()时，把字段PassWord弄成小写password发送到前端

	Like string //这一个不带任何tag标签   通过浏览器输入query路径参数，无tag时，
	//只要键值对的键大小写与结构体字段完全相同，shouldbind就可成功绑定该字段
}

//上述关于form的tag标签的例子，见bang3ding4form文件夹里的内容，下面的是一个奇怪的现象。
func main() {
	r := gin.Default()

	//下面这种奇怪的情况是
	//浏览器输入框输入http://localhost:8888/login?mingzi=jack&mima=123&Like=xixixi，
	//奇怪的成功借助form的tag完成绑定，返回给客户端的是{"Name":"jack","PassWord":"123","Like":"xixixi"}
	//但输入http://localhost:8888/login?Name=jack&PassWord=125&Like=xixixi 返回给前端的是
	//{"Name":"","PassWord":"","Like":"xixixi"}   并不能成功实现Like之外的字段的绑定，
	//对照bang3ding4param文件夹中main.go里笔记，可见是form的tag标签产生了原本是针对
	//表单现在却发挥在query路径参数上进行绑定的影响

	//----------------
	//最终结合另外几个文件夹的实验，总结可知，shouldbind配合form的tag标签，既可以解决单纯的前端传
	//表单数据（form的tag标签发挥协助映射比对作用），也可以解决单纯的前端传query参数数据（就是上述这
	//个奇怪的现象），也可以解决用postman发送的json格式数据（无需tag,没用上这里的form的tag标签）
	//但对于前端一次请求既包含表单数据又包含query路径参数数据的情况，只会成功绑定表单数据，而不绑定query路径参数
	//----------------
	r.GET("/login", func(c *gin.Context) {
		u := user{}             //申明一个结构体，下一行就把客户端传来的数据放进来
		err := c.ShouldBind(&u) //这里传进来的u必须有&，才能获取c中的参数、用ShouldBind的反射
		//机制将c中包含的前端数据的键（键值对）去比对tag标签来把值（键值对）的数据赋值到相应字段，
		//使得接下来的u不再是申明时候的初始默认值，而是拿到了客户端的值，

		if err != nil {
			c.JSON(200, "c.ShouldBind获取客户端数据赋值进入结构体失败")
			return
		} else {
			c.JSON(200, u)
			//此时这里的u的值已经不再是{Name:"",PassWord:""},而是两个字段都获得了客户端的值
		}
	})
	r.Run(":8888")
}
