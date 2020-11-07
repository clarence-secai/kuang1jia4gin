//todo:gin.Context下处理前端传来的文件的系列方法如下
//FormFile(name string) (*multipart.FileHeader, error)
//MultipartForm() (*multipart.Form, error)
//SaveUploadedFile(file *multipart.FileHeader, dst string) error

package main
import  (
	"fmt"
	//"fmt"
	"github.com/gin-gonic/gin"
	"path"
)

//由于自己电脑系统的原因，路径是反斜杠，和一般的不一样，
//导致无法获得预期的效果。不过代码没错
func handler(c *gin.Context) {
	c.HTML(200, "a.html", nil)
}
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./a.html") //解析文件
	r.GET("/index", handler)
	r.POST("/upload", func(c *gin.Context) {
		fh, err := c.FormFile("wenjian") //获得上传的文件头信息*multipart.FileHeader
		if err != nil {
			c.JSON(200, "上传文件出错")
			return
		}
		//dst1 := fh.Filename//取到上传的文件头信息结构体的filename字段，即上传的文件的文件名
		//dst := fmt.Sprintf("I:/go/%s",dst1)

		//_,wenjianming := path.Split(f.Filename)//这一步不需要，Filename字段已经是不包含路径的纯文件名
		//c.JSON(200,f.Filename)
		//c.JSON(200,wenjianming)

		//也可以用下面的方式
		dst := path.Join("./", fh.Filename)
		fmt.Println(dst)
		err2 := c.SaveUploadedFile(fh, dst)
		if err2 != nil {
			c.JSON(200, err2)
			return
		}
		c.JSON(200, "上传成功")
	})

	r.Run(":8888")
}
