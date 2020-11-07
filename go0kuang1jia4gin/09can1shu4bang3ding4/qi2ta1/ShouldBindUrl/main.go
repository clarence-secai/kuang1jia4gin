//todo:必须是uri标签，绑定uri路径参数也只能用ShouldBindUri
package main
import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type st struct {
	A string `uri:"a"`
	B string `uri:"b"`
	C string `uri:"c"`
	//这里的tag的类型uri对应ShouldBindUri这种情况，tag后面的内容对应"/index/:a/:b/:c"
}

func main() {
	r := gin.Default()
	r.GET("/index/:a/:b/:c", func(c *gin.Context) {
		//m := make(map[string][]string,2)
		var m st
		c.ShouldBindUri(&m)
		fmt.Println("m是", m)   //成功输出   m是 {jack mary bob}
		c.String(200, "%v", m) //成功向前端返回{jack mary bob}
	})

	r.Run(":8888")
}
