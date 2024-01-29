package main

import (
	"fmt"
	"net/http"
	"log"
	"strings"
	"html/template"
)

func login(w http.ResponseWriter, r *http.Request){
	fmt.Println("method", r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("login.html")
		if err != nil {
			fmt.Println("err:",err)
		}
		log.Println(t.Execute(w, nil))
	}else {
		//请求的是登录数据
		_ = r.ParseForm()
		fmt.Println("username:",r.Form["username"])
		fmt.Println("password:",r.Form["password"])
		if pwd := r.Form.Get("password"); pwd == "123456" {
			fmt.Fprintf(w, "欢迎登录, Hello %s", r.Form.Get("username"))
		}else {
			fmt.Fprintf(w, "登录失败")
		}
	}
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm
	fmt.Println(r.Form)
	fmt.Println("path:", r.URL.Path)
	for k,v := range r.Form{
		fmt.Println("key:",k)
		fmt.Println("val:",strings.Join(v,""))
	}
	fmt.Fprintf(w, "Hello,aoho!")
}

func main(){
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("err:",err)
	}
}