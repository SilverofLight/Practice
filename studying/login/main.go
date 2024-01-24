package main
import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)//获取请求方法
	if r.Method == "GET" {
		t, err := template.ParseFiles("login.tql")
        if err != nil {
            log.Fatal("Error parsing template:", err)
        }
		log.Println(t.Execute(w, nil))
	} else {
		//请求的是登录数据，那么执行登录的逻辑判断
		_ = r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		if pwd := r.Form.Get("password"); pwd == "123456" {//验证密码是否正确
			fmt.Fprintf(w, "欢迎登录，Hello %s!", r.Form.Get("username"))//这个写入到w的是输出到客户端的
		}else {
			fmt.Fprintf(w, "密码错误，请重新输入！")
		}
	}
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()//解析url传递的参数，对于POST则解析相应包的主体（request body）
	//注意：如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(r.Form)//这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("Val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello aoho!") //写入w，输出到客户端
}

func main(){
	http.HandleFunc("/", sayHelloName) //设置访问的路由
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":8080", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
/*这段代码是一个简单的Go语言程序，使用了net/http包创建一个基本的Web服务器。主要功能包括：

    路由设置：
        / 路径映射到 sayHelloName 函数。
        /login 路径映射到 login 函数。

    sayHelloName 函数：
        用于处理 / 路径的请求。
        解析请求参数，输出请求信息和表单数据。
        向客户端输出 "Hello aoho!"。

    login 函数：
        用于处理 /login 路径的请求。
        如果是 GET 请求，渲染登录页面模板并返回给客户端。
        如果是 POST 请求，解析表单数据并验证用户名和密码，返回相应的消息。

    模板文件 "login.tql"：
        包含一个简单的HTML表单，用于用户登录。

    程序入口 main 函数：
        设置路由处理函数。
        启动监听端口为 8080 的服务器。

回答的问题包括：

    POST请求：
        POST请求是HTTP协议中的一种请求方法，用于向服务器提交数据，通常用于创建新资源、更新现有资源或执行可能对服务器状态产生影响的操作。

    fmt.Fprintf(w, ...)：
        用于将格式化的字符串写入实现了 io.Writer 接口的对象，通常用于将响应信息写回给客户端。

    _ = r.ParseForm()：
        用于解析HTTP请求中的表单数据，将其存储在 r.Form 字段中。

    模板文件的作用：
        模板文件 "login.tql" 用于渲染登录页面的HTML内容，使用户能够在浏览器中看到并填写用户名和密码。

    日志输出：
        使用 fmt.Println 和 log.Println 在控制台输出一些请求和处理过程的日志信息，用于调试和观察程序运行情况。*/
//GPT问答：https://chat.openai.com/share/d1a655b0-d2f1-4640-8008-249df4530164