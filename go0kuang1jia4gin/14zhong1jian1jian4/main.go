//todo:多个中间件与处理器的执行顺序问题
// 中间件的本质其实就是一个处理器，写在路由的前面进行执行，来完成以下相对公共的功能。gin框架
//的中间件要求必须是gin.HandlerFunc类型，也即func(*Context)类型【也就是一般处理器的类型】
//重点是gin框架下的中间件，可通过在中间件middleware中使用c.Set() c.next() c.Abort()等来和
//路由后面的handler沟通。
package main
import "github.com/gin-gonic/gin"


func middleware1(c *gin.Context) { //这个中间件可以在main()函数中的任意路由中被使用
	c.JSON(200, "我是中间件1")
}
//gin框架的中间件要求必须是gin.HandleFunc类型，也即func(*Context)类型
//故中间件也可写成如下形式
func middleware0(a interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "我是中间件0")
	}
}
func middleware2(c *gin.Context) { //这个中间件可以在main()函数中的任意路由中被使用
	c.JSON(200, "我是中间件2")
	c.Abort() //表示不予执行该中间件后面的处理器(包括中间件)，本例中就是不运行handler函数
	//，即不显示“我是handler处理器”这句话，Abort常用于通过判断来决定是否执行后面的处理器
}
func AllMidHand1(c *gin.Context) {
	c.JSON(200, "我是后续总中间件1")
}
func AllMidHand2(c *gin.Context) {
	c.JSON(200, "我是后续总中间件2")
}
func handler(c *gin.Context) {
	c.JSON(200, "我是handler处理器")
}
func main() {
	r := gin.Default()

	r.GET("/ha", func(c *gin.Context) {
		//todo:该路由由于是在r.Use()之前，故不会运行AllMidHand1,AllMidHand2中间件
		c.JSON(200, "我是注册后续总中间件之前的匿名处理器")
	})

	//todo:统一注册一个后面都必用的中间件,会在其后的路由中每次都被默认调用，且是先于接下来其他手写调用的中间件之前调用的
	r.Use(middleware0("ha"),AllMidHand1, AllMidHand2) //也可以注册三个中间件，会按顺序执行
	r.GET("/index", middleware1, func(c *gin.Context) {
		c.JSON(200, "我是匿名函数处理器")
	})
	//上面返回客户端的就是	 "我是中间件0""我是后续总中间件1""我是后续总中间件2""我是中间件1""我是匿名函数处理器"

	r.GET("/login", middleware2, handler)
	//这里客户端显示   "我是中间件0""我是后续总中间件1""我是后续总中间件2""我是中间件2"
	//但没有"我是handler处理器" 因为middleware2中有一个r.Abort()

	r.Run(":8888")
}
