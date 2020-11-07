package main

//是html还是text，由引入的包名决定
import (
	//"html/template" //安全模式，以纯字符串文本的形式显示，不会去识别，本例中a和s都以字符串形式完整显示
	"net/http"
	"text/template"//不安全模式，会自动做识别处理，本例中s会显示123，a显示可以导航到b站的链接
)


func main() {
	s := "<script>alert(123);</script>"
	a := "<a href='https://www.bilibili.com'>b站</a>"
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("./c.html")
		t.Execute(w,s)
	})
	http.HandleFunc("/index2", func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("./c.html")
		t.Execute(w,a)
	})

	http.ListenAndServe(":8888", nil)
}
