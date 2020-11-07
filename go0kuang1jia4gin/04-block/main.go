 package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//仔细分析下面代码在浏览器中显示的情形，可见含有block的主文件a.html和用define来
//填充block的辅文件b.html，在t.ExecuteTemplate执行主文件还是辅文件时，会有差异：
//1、当辅文件中{{template "a.html" .}}无论有点还是无点，执行主文件，正常传参给主板和辅板，
//但辅板define外的内容丢失；
//2、当辅文件中{{template "a.html" }}无点，执行辅文件，主板、辅板、辅板define外的内容都执行，
//但不传参给主板、辅板，只传参给辅板define外的内容
//3、当辅文件中{{template "a.html" .}}有点，执行辅文件，主板、辅板、辅板define外的内容都执行，
//且都正常传参
//总结：上述情况一般执行辅板文件（执行主板文件缺乏意义，辅文件中{{template "a.html" }}无点，
//则不给主板block辅板define对应的部分传参（但执行相应文本），只给辅板define之外的部分传参；
//有点，则均传参执行。因此，实际操作时必须按照第3点要求来操作，
//即主文件的block、辅文件的template，都写点并且执行辅文件

func myhandler1(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./a.html", "./b.html")
	//todo:注意，要按照模板的包含与被包含的关系，a.html必须写在b.html文件前面
	// t.Execute(w,"主板") 这个数据默认是传给了a.html文件里的{{.}}
	t.ExecuteTemplate(w, "b.html", "b主板")
}
func myhandler2(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./a.html", "./c.html")
//todo:注意，上面解析文件，需要写相对路径文件名，下面的执行某个模板，只能写文件本身名字，
// 或者define自定义的模板名字（不定义就默认把文件名当做模板名），不能再有相对路径

	t.ExecuteTemplate(w, "c.html", "c主板")
	fmt.Fprint(w, "分割线-----------------------------")
	t.ExecuteTemplate(w, "a.html", "主板")

	/*跑起来后浏览器显示如下
	aa这里是大框架，绝大多数模板文件都是这个框架里的某些部分填充内容:c主板
	这是c.html文件里的内容，接收后台来的数据:c主板
	aa这里是大框架，绝大多数模板文件都是这个框架里的某些部分填充内容:c主板
	ccc这是c.html文件block之外的内容，接收后台来的数据:c主板
	---分割线----------------------------- aa这里是大框架，绝大多数模板文件都是这个框架里的某些部分填充内容:主板
	这是c.html文件里的内容，接收后台来的数据:主板
	aa这里是大框架，绝大多数模板文件都是这个框架里的某些部分填充内容:主板
	*/
}

func main() {
	http.HandleFunc("/index1", myhandler1)
	http.HandleFunc("/index2", myhandler2)

	http.ListenAndServe(":8888", nil)
}
