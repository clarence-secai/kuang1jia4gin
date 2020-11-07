package main

import (
	"html/template"
	"net/http"
	"strconv"
)

type Student struct {
	Name    string
	Age     int64
	Fun     func(string) string
	     //todo:Fun必须是函数类型，不能是自定义的本质是函数类型的类型，不然html文件中{{call .Fun "总是"}}无法有效
	Qiepian []string
}
func (s *Student) Game() int {  //前端可以调用结构体下无需传参的大写开头的方法
	a := 3
	return a
}
func myhandler(w http.ResponseWriter, r *http.Request) {
	study := func(s string) string {
		return s + "在努力学习"
	}
	stu := Student{Name: "jack", Age: 27, Fun: study, Qiepian: []string{"足球", "篮球", "乒乓"}}
	t, _ := template.ParseFiles("./b.html")
	t.Execute(w, &stu) //todo:这里千万记得是传地址，否则前端无法正常调用(s *Student)Game()int
}


//下面这样handler2操作是错误示范
//func Hanshu() string {
//	return "哈哈，我是纯函数"
//}
//func handler2(w http.ResponseWriter, r *http.Request) {
//	t, _ := template.ParseFiles("./c.html")
//	t.Execute(w, Hanshu) //会发现前端调用后端函数不是这么操作的
//}


type hobby struct {
	Name   string
	People int64
}
func (h *hobby) Str() string { //todo:记得方法名开头大写，否则客户端无{{.str}}是没法实现的
	return h.Name + " has " + strconv.FormatInt(h.People, 10) + " like it "
}
//下面这样前端也是错误示范，前端无法调用需要传参的方法
func (h *hobby)Two(s string)string{
	return s+" likes "+h.Name
}
func handler3(w http.ResponseWriter, r *http.Request) {
	aa := &hobby{Name: "soccer", People: 22} //千万记得&，否则客户端{{.Str}}无法实现
	t, _ := template.ParseFiles("./d.html")
	t.Execute(w, aa) //此时aa已经是指针类型
}


type number struct {
	No1 int64
	No2 int64
	No3 int64
}
func lasthandler(w http.ResponseWriter, r *http.Request) {
	num := number{0, 9, 5} //
	t, _ := template.ParseFiles("./e.html")
	t.Execute(w, num)
}

func main() {
	http.HandleFunc("/school", myhandler)
	//http.HandleFunc("/math", handler2)
	http.HandleFunc("/hobby", handler3)
	http.HandleFunc("/last", lasthandler)
	http.ListenAndServe(":8888", nil)
}
