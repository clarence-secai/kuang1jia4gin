package main

import (
	"html/template"
	"net/http"
)

//将模板文件中的{{.}}进行自定义
type student struct {
	Name string
	Age  int64
}

func main() {
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		//下面定义替换默认的{{.}}为{[.]}
		t := template.New("a.html").Delims("{[", "]}")
		//自定义结束后才可以解析模板文件
		t.ParseFiles("./a.html")
		pupil := student{Name: "jack", Age: 27}
		t.Execute(w, pupil)
	})
	http.ListenAndServe(":8888", nil)
}
