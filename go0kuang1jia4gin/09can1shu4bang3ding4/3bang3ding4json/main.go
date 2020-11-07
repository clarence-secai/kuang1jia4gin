package main
import "github.com/gin-gonic/gin"

//todo:json的tag标签是针对postman发送的json数据借助shouldbind进行反射绑定，以及在c.JSON向前端返回数据时tag替
//代结构体字段的数据。因此总的来说，
//1、前端发数据来借助shouldbind用键值对的键去比对相应类型（如json、form、uri、binding)的tag后，把值赋给
//相应的后端字段
//2、后端使用c.JSON向前端发送数据，会将各个结构体字段变成tag是json的json冒号后面的内容作为键值对的键连同
//值一起给前端


//todo:当postman向后端发json数据，会发现无需json的tag标签，只要键值对的键跟这里的字段能对应上，不区分大小写
//都会成功绑定到结构体字段，前端多出的对不上的字段，就不予绑定，前端缺少的字段，返回时就取零值.例如前端
//postman发送json数据{"hobby":"mary","haha":99}，绑定后返回的是{"Hobby": "mary","Age": 0}
//todo:如果加上json的tag标签，前端的键值对的键就需和后端这里的json的tag标签内容一致，而不再是跟后端
//这里的与之对应的结构体字段一样或大小写不区分的一样,比如取消注释下面的Like字段
//postman发送{"hobby":"mary","haha":99,"Like":"paly"} ，返回的结果是
//{"Hobby": "mary","Age": 0,"wahaha": ""}

type kid struct {
	Hobby string
	Age   int64
	//Like string `json:"wahaha"`
}

func main() {
	r := gin.Default()
	//下面用postman传json格式数据
	r.POST("/into", func(c *gin.Context) {
		var k kid
		err := c.ShouldBind(&k)//todo:具体绑定情况见上面的解释
		if err != nil {
			c.JSON(200, "c.ShouldBind获取客户端数据赋值进入结构体失败")
			return
		} else {
			c.JSON(200, k)

		}
	})

	r.Run(":8888")
}
