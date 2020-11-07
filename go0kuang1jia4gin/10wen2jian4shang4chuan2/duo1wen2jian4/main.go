//todo:gin.Context下处理前端传来的文件的系列方法如下
//FormFile(name string) (*multipart.FileHeader, error)
//MultipartForm() (*multipart.Form, error)
//SaveUploadedFile(file *multipart.FileHeader, dst string) error
package main
import (
	"github.com/gin-gonic/gin"
	"path"
)

func handler1(c *gin.Context) {
	mulform, _ := c.MultipartForm()
	fileHeaders := mulform.File["duowenjian"] //todo:此处就是填HTML文件input里的name="duowenjian"
	for _, v := range fileHeaders {
		dst := path.Join("./", v.Filename)
		c.SaveUploadedFile(v, dst)
		c.JSON(200, "上传成功")
	}
}
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./b.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(200, "b.html", nil)
	})
	r.POST("/uploads", handler1)

	r.Run(":8888")
}
