package main

import "github.com/gin-gonic/gin"

//路由组注册中间件的两种方式如下
//路由组注册的中间件，只在该路由组中发挥作用
//gin框架的中间件要求必须是gin.HandleFunc类型，也即func(*Context)类型
//滚轮查看可知类型HandleFunc的本质就是func(*Context)，和处理器的类型一样
func midhand(ok bool) gin.HandlerFunc {
	if ok {
		return func(c *gin.Context) {
			c.JSON(200, "这是中间件")
		}
	}
	return nil
}
func main() {
	r := gin.Default()
	r.Use(midhand(true)) //这里注册一个后面总中间件，不影响后面的路由组或单个路由再次调用该中间件，效果就是会实现两次

//在路由组中注册中间件的方式一
	indexgroup := r.Group("/index", midhand(true))
	    //可以对照之前的路由组在此处是否写处理器的不同，这里就等同于是在路由组内注册一个公共的中间件
	{
		indexgroup.GET("/path1", func(c *gin.Context) {
			c.JSON(200, "这里是/index/path1")
		})
		indexgroup.GET("/path2", func(c *gin.Context) {
			c.JSON(200, "这里是/index/path2")
		})
	}
	//下面展示了可以将一个中间件在另一个路由组中再次注册使用
	logingroup := r.Group("/login", midhand(true))
	{
		logingroup.GET("/road1", func(c *gin.Context) {
			c.JSON(200, "这里是/login/road1")
		})
		logingroup.GET("/road2", func(c *gin.Context) {
			c.JSON(200, "这里是/login/road2")
		})
	}

//注册中间件的方式二
	registgroup := r.Group("/regist")//todo:r.Group()的第二个参数是可变参数，故可以不填
	registgroup.Use(midhand(true))
	{
		registgroup.GET("/house1", func(c *gin.Context) {
			c.JSON(200, "这里是/regist/house1")
		})
	}

	r.GET("/ha", midhand(true), func(c *gin.Context) {
		c.JSON(200, "这里是/ha")
	})

	r.Run(":8888")
}
