//todo:一下是gin.Context下通过中间件来使用的Set和Get系列方法
//Set(key string, value interface{})
//Get(key string) (value interface{}, exists bool)
//MustGet(key string) interface{}
//GetString(key string) (s string)
//GetBool(key string) (b bool)
//GetInt(key string) (i int)
//GetInt64(key string) (i64 int64)
//GetFloat64(key string) (f64 float64)
//GetTime(key string) (t time.Time)
//GetDuration(key string) (d time.Duration)
//GetStringSlice(key string) (ss []string)
//GetStringMap(key string) (sm map[string]interface{})
//GetStringMapString(key string) (sms map[string]string)
//GetStringMapStringSlice(key string) (smss map[string][]string)
package main
import "github.com/gin-gonic/gin"
func middleware(c *gin.Context) {
	c.Set("message", "客户端传来的数据是")
	c.Next() //todo:此处调用处理器，执行完处理器后继续运行下一行的c.JSON
	c.JSON(200, "这里是中间件运行结束")
}
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./a.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(200, "a.html", nil)
	})
	r.POST("/login", middleware, func(c *gin.Context) {
		message := c.MustGet("message")//todo:获取由中间件通过c.Set()设置的值
		username := c.PostForm("xingming")
		c.JSON(200, []interface{}{message, username})
	})

	r.Run(":8888")
}
//当在前端a.html页面填写Clarence并点击提交后，运行结果如下：
/*
["客户端传来的数据是","Clarence"]"这里是中间件运行结束"
*/