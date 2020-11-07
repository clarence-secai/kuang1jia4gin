package main

//是html还是text，由引入的包名决定
import (
	"html/template" //安全模式，以纯字符串文本的形式显示，不会去识别，本例中a和s都以字符串形式完整显示
	"net/http"
	//"text/template"//不安全模式，会自动做识别处理，本例中s会显示123，a显示可以导航到b站的链接
)

func hanshu(str string) template.HTML {
	return template.HTML(str) //todo:即解除html/template自带的安全模式
	//这里将原本不被安全模式的html/template包自动识别的内容，改成像text/template一样可自动识别
}
func myhandler(w http.ResponseWriter, r *http.Request) {
	t := template.New("b.html").Funcs(template.FuncMap{"unsafe": hanshu})
	t.ParseFiles("./b.html")
	var b = "<a href='https://www.bilibili.com'>b站</a>"
	t.Execute(w, b)
}

func main() {

	//如果是接收客户端传来数据予以显示，如何方便地选择决定哪些予以识别显示哪些予以直接字符串显示呢，
	//这就用到template.Funcs函数，就可以让前端自主决定是否调用后台的自定义函数来选择是识别模式还是安全模式
	http.HandleFunc("/index3", myhandler)

	http.ListenAndServe(":8888", nil)
}
