package main

import (
	"html/template"
	"net/http"
)

type student struct {
	Name string
	Age  int64
}
func myhandler(w http.ResponseWriter, r *http.Request) {
	kua := func(s string) string { //被写进模板的函数，返回值只能是一个或两个中的第二个是error
		return s + "是高富帅"
	}
	stu := student{Name: "jack", Age: 27}
	t := template.New("a.html")
	//todo:上一行是为了提供下一步注册自定义函数所需的t，这里new的名字必须是a.html，和后面解析的模板文件名对应上
	//如果代码语句是   t,_:=template.New("abc").Parse("这是内容：{{.}}")
	//然后t.Execute(w,stu)则名字是临时的随便取 ，例如下面的handler2
	t.Funcs(template.FuncMap{"zidingyi": kua})
	//这时候t里面就包含了自定义注册进来的这个函数，下面的t.Execute(w,stu)在向前端传stu的同时这个自定义函数也传了，即有链式调用的感觉
	t.ParseFiles("./a.html")//todo:只能先注册自定义函数到模板上，然后解析模板，前端才能实现调用自定义函数
	t.Execute(w, stu)
}
func pre(s string)string{
	return s + " 你好"
}
func handler1(rw http.ResponseWriter,r *http.Request){
	t := template.New("b.html")
	myMap := make(map[string]interface{})
	myMap["pre"]=pre
	////接下来方式一
	//t.Funcs(myMap)   //和21行的方式是等价的
	//t.ParseFiles("./b.html")
	//t.Execute(rw,"")
	////接下来效果和方式一等价的方式二    因此链式调用和非链式调用效果是一样的，非链式调用前一步运行会影响后一步的t
	tt, _ := t.Funcs(myMap).ParseFiles("./b.html")
	tt.Execute(rw,"")

}

func handler2(w http.ResponseWriter, r *http.Request) {
	stu2 := student{Name: "tom", Age: 27}
	t, _ := template.New("a").Parse("这是无单独文件的模板文件：{{.Name}}的年龄是{{.Age}}")
	t.Execute(w, stu2) //浏览器上显示：    这是无单独文件的模板文件：tom的年龄是27
}
func main() {
	http.HandleFunc("/index", myhandler)
	http.HandleFunc("/index1", handler1)
	http.HandleFunc("/student", handler2)
	http.ListenAndServe(":8888", nil)
}
