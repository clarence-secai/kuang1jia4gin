//todo:ShouldBindQuery只能成功绑定query路径参数
package main
import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type kid struct {
	Name      string `form:"mingzi"`
	PassWord  string `form:"mima"`
	SmallName string `form:"nick"` //todo:此处这个form的tag标签必须有，前端html的query键值对的键也必
	// 须跟这个tag保持一样，尤其不能这里已经有`form:"nick"`，前端html里的query键值对的键
	//又去用SmallName  否则连query参数也绑定不上
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./a.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(200, "a.html", nil)
	})
	r.POST("/login", func(c *gin.Context) {
		var k kid
		c.ShouldBindQuery(&k)
		fmt.Println("k=", k)
		fmt.Println("k的两个字段是", k.Name, k.SmallName)
		c.JSON(200, k)
	})
	//cmd控制台的结果是
	//k= {  huahua}
	//huahua
	//todo:这两行是运行结果，可见ShouldBindQuery只能成功绑定query查询参数
	r.Run(":8888")
}
