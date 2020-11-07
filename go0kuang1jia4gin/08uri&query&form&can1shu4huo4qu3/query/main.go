//todo:gin.Context下的query系列方法，是用于获取url上的查询字符串参数，系列方法如下：
//Query(key string) string
//DefaultQuery(key string, defaultValue string) string
//GetQuery(key string) (string, bool)
//QueryArray(key string) []string
//GetQueryArray(key string) ([]string, bool)
//QueryMap(key string) map[string]string
//GetQueryMap(key string) (map[string]string, bool)
//以下可以鼠标浮停看详情解释和举例，也可以浮停查看提供的全部字段和方法

package main
import "github.com/gin-gonic/gin"
func main() {
	r := gin.Default()
	//获取浏览器url传来的数据

	//使用Query获取浏览器url传来的数据

	//当在浏览器输入http://localhost:8888/index1?query=haha时，运行下面页面上会显示"haha"
	r.GET("/index1", func(c *gin.Context) {
		message := c.Query("query")
		c.JSON(200, message)
	})
	//当在浏览器输入http://localhost:8888/index11?key=wowo时，运行后页面上会显示"wowo"
	r.GET("/index11", func(c *gin.Context) {
		value := c.Query("key")
		c.JSON(200, value)
	})

	//使用DefaultQuery获取浏览器url传来的数据

	//当前端不是以c.DefaultQuery中"name"为key传参，str就会是defaultValue的值
	r.GET("/index2", func(c *gin.Context) {
		str := c.DefaultQuery("name", "客户端不是以name为key传值的")
		c.JSON(200, str)
	})

	//使用GetQuery获取浏览器url传来的数据  具体可以鼠标浮停看解释
	r.GET("/index3", func(c *gin.Context) {
		strs, ok := c.GetQuery("text")
		if !ok {
			c.JSON(200, "客户端不是以text为key传值的")
		}
		c.JSON(200, strs)
	})

	//一次获取浏览器url传来的多个数据
	r.GET("/index4", func(c *gin.Context) {
		str1 := c.Query("query")
		str2 := c.Query("name") //todo:url上没有key为name的值传来时，返回值是空字符串，即零值
		str3 := c.Query("text")
		c.JSON(200, []string{str1, str2, str3})
	})

	//在浏览器输入http://localhost:8888/index5?qiepian=大话西游&qiepian=娶你的
	//会返回多个key为qiepian的值
	r.GET("/index5", func(c *gin.Context) {
		slice := c.QueryArray("qiepian")
		c.JSON(200, slice)
	})
	r.Run(":8888")
}
