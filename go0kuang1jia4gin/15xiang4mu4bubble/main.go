package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" //这个驱动不要忘记，不然连不上数据库
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	//"database/sql"
)

type Todo struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"` //默认零值是 false
}

func main() {
	mydb, err := gorm.Open("mysql", "root:413188ok@tcp(localhost:3306)/gormtable?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer mydb.Close()
	//下面在数据库中创建相应的表
	mydb.AutoMigrate(&Todo{})

	r := gin.Default()
	r.LoadHTMLFiles("./bubble/index.html")
	r.Static("/static", "./bubble/static")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	v1group := r.Group("/v1")
	{ //添加代办事项
		v1group.POST("/todo", func(c *gin.Context) {
			var todo1 Todo
			//c.ShouldBind(&todo1) //此时todo1各个字段内容已是客户端填好发送的信息
			//由于前后端总是以json数据沟通，故应该用下面这个绑定方式
			c.BindJSON(&todo1)               //此时todo1各个字段内容已是客户端填好发送的信息
			fmt.Println("新增任务todo1=", todo1) //运行结果为 新增任务todo1= {0 试一试shiyishi false}
			err := mydb.Create(&todo1).Error //这个也必须有&  id虽然是零值，但数据库表会自增给一个值
			if err != nil {
				c.JSON(http.StatusOK, err)
			} else {
				c.JSON(200, todo1) //成功向数据库表新增后，发给前端,其实没必要，少这一步没任何影响
			}

		})
		//查询所有事项
		v1group.GET("/todo", func(c *gin.Context) {
			//sql := "select * from todos"  还是不能混用sql语句，不能直接用gorm打开的mydb用在官方文档的db.Query方法里
			var todos []Todo
			mydb.Find(&todos)
			c.JSON(200, todos)
		})
		//修改某一个事项,实际上本不只需要更新前端改动后发来的数据，查询会自动到上一个处理器去
		v1group.PUT("/todo/:id", func(c *gin.Context) {
			idstr := c.Param("id")
			fmt.Println("前端发来的idstr=", idstr)
			var a string
			c.BindQuery(&a)
			fmt.Println("修改操作中浏览器传来的query参数绑定结果是a=", a) //结果是a=  没啥结果，可见BindQuery没取到url上的参数
			var todo1 Todo
			c.ShouldBind(&todo1)
			fmt.Println("todo1=", todo1) //运行结果为todo1={0 "" true}，
			// 可见c.ShouldBind是只能拿到请求体的内容，只获取到true，没获取到id和title
			todo1.ID, _ = strconv.ParseInt(idstr, 10, 64)

			mydb.Model(&todo1).Update("status", todo1.Status) //此时的todo1其实是{29,"",true}，
			// 因此gorm操作，只需要&todo1里的ID，其他字段无关系，gorm会按ID找到数据库表里的相应记录予以操作

			//c.JSON(200, todo1) //猜测前端要的其实就是它自己传到后端的那个true,但单纯只这一步
			//不更改数据库，会导致刷新网页后又变回原样，没真正更改
		})
		v1group.DELETE("/todo/:id", func(c *gin.Context) {
			var todo2 Todo
			todo2.ID, _ = strconv.ParseInt(c.Param("id"), 10, 64)
			mydb.Delete(&todo2) //todo2字段里只有ID有值，其他均为默认零值，可见gorm这种删除方式，
			// 只需要&todo2里的ID，根据ID找到数据库表对应的记录予以操作
		})
	}

	r.Run(":8888")
}
