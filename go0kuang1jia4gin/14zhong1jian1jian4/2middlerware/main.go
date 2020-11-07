package main
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func middlerware(c *gin.Context) {
	fmt.Println("中间件开始运行")
	c.JSON(200, "这里是middleware中间件")
	c.Set("name", "jack") //设置值，供后续处理器使用
	c.Next()              //用于调用后续的处理器    与之相反的c.Abort()是阻止调用后续处理器的
	fmt.Println("中间件结束运行")
}
func main() {
	r := gin.Default()
	r.GET("/index", middlerware, func(c *gin.Context) { //将中间件处理器写在其他处理器前面
		fmt.Println("匿名函数开始运行")
		a := c.MustGet("name") //todo:获取由中间件通过c.Set()设置的值
		//这里还有c.Get() c.GetString() 等等用Get + 各种类型  详情可以点击滚轮去看
		//其中c.Get()的返回值是interface和bool,bool是判断是否有返回值的，即之前是否有
		//c.Set()的相应key的值的。而MustGet会在拿不到c.Set()设置的key的值时直接panic
		c.JSON(http.StatusOK, "这里是匿名函数处理器")
		c.JSON(200, a)
		fmt.Println("匿名函数结束运行")
	})
	r.Run(":8888")
}
//浏览器和cmd运行结果如下
/*
"这里是middleware中间件""这里是匿名函数处理器""jack"
中间件开始运行
匿名函数开始运行
匿名函数结束运行
中间件结束运行
*/