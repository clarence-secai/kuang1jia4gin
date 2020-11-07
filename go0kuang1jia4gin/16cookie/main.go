package main
import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//下面的代码跑起来不知为什么用火狐浏览器可以正常运行符合预期结果
//但用explorer浏览器，cookie总是Goland-1f9bc8b0等一串字符
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./a.html")
	r.GET("/index", func(c *gin.Context) { //Goland-1f9bc8b0
		//获取前端发来请求中包含的cookie
		str, err := c.Cookie("usercookie")
		if err != nil {                    //表明没有找到对应的cookie
			fmt.Println("没有找到相应的cookie：", str)
			//todo:前端没有cookie时，这里获取的cookie的结果str是空字符串

			//todo:设置cookie
			c.SetCookie("usercookie", "choclate",
				600, "/", "localhost", false, true)

			//todo:设置cookie的方法中的各参数的意思解释如下：
			//name string, 设置cookie的键值对的键
			//value string, 设置cookie的键值对的值
			//maxAge int,   设置cookie在前端浏览器中保留时间，单位是秒
			//path string,  设置cookie在前端浏览器中保存的地址路径
			//domain string, 域名，应该是填浏览器所在计算机的域名
			//secure bool,   true表示必须以HTTPS访问，false表示可以以http就能访问
			//httpOnly bool   true表示只允许以http的方式获取cookie，不允许其他方式如js来获取cookie
			return
		}
		fmt.Println("找到的cookie为:", str)
		c.HTML(200, "a.html", str)
	})

	r.Run(":8888")
}
