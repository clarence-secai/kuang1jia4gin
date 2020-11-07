package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	//r.Group可不写处理器，即使写，客户端网址也不能输入到/index就截止,否则依然404，
	//但写上，会在输入/index/path1时，连通Group的一起显示，即客户端显示为    访问的是/index访问的是/index/path1
	//indexgroup这个名字可自定义
	indexgroup := r.Group("/index", func(c *gin.Context) {
		c.JSON(200, "访问的是/index")
	},
	)

	//下面这个大括号也可不加，加上有助于区分
	{
		indexgroup.GET("/path1", func(c *gin.Context) {
			c.JSON(200, "访问的是/index/path1")
		})
		indexgroup.GET("/path2", func(c *gin.Context) {
			c.JSON(200, "访问的是/index/path2")
		})
		//可以夹杂其他请求类型如POST
		indexgroup.POST("/path3", func(c *gin.Context) {
			c.JSON(200, "访问的是/index/path3")
		})

		//还可以在上面的indexgroup内进行路由组的嵌套

		indexpathgroup := indexgroup.Group("/road", func(c *gin.Context) {
			c.JSON(200, "访问的是/index/road")
		}) //一般不建议写返回访问的是/index/road的上面的这个handler
		{
			//针对下面这个，在客户单需写localhost:8888/index/road/house1
			//返回客户端的是    "访问的是/index""访问的是/index/road""访问的是/index/road/house1"
			indexpathgroup.GET("/house1", func(c *gin.Context) {
				c.JSON(200, "访问的是/index/road/house1")
			})

			indexpathgroup.POST("/house2", func(c *gin.Context) {
				c.JSON(200, "访问的是/index/road/house2")
			})
		}

	}
	r.Run(":8888")

}
