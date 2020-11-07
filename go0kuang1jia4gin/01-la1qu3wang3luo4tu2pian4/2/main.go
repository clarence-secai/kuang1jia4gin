package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func myhander(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("a.html")
	t.Execute(rw, "golang你好")
}
func main() {
	fmt.Println("hello,world")
	http.HandleFunc("/index", myhander)

	http.ListenAndServe(":6666", nil)
}
