//todo:gin.Context下的PostForm系列的方法，是用于获取表单参数的，系列方法如下：
//PostForm(key string) string
//DefaultPostForm(key string, defaultValue string) string
//GetPostForm(key string) (string, bool)
//PostFormArray(key string) []string
//GetPostFormArray(key string) ([]string, bool)
//PostFormMap(key string) map[string]string
//GetPostFormMap(key string) (map[string]string, bool)
//更多可以鼠标浮停查看全部字段和方法，或浮停查看相应解释

package main
import "github.com/gin-gonic/gin"
type stu struct {
	UserName string
	PassWord string
	Message  string
}

func handler1(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
func handler2(c *gin.Context) {
	//获取客户端form表单提交的数据
	username := c.PostForm("username") //相当于GetPostForm忽略ok这个返回值
	password := c.PostForm("password")//没取到值就得到空字符串，即零值
	//也可以采用下面这种方式
	mima := c.DefaultPostForm("word", "没获取到name为word的表单数据")
	//还有下面这种方式
	//xinxi,ok := c.GetPostForm("word")
	//if !ok {
	//	return
	//}
	s := stu{UserName: username, PassWord: password, Message: mima}
	c.HTML(200, "login.html", s)
}
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html", "./login.html")
	r.GET("/index", handler1)
	r.POST("/login", handler2)

	r.Run(":8888")
}
