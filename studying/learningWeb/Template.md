# Go语言的模板引擎

Go语言内置了文本模板引擎`text.template`和用于HTML文档的`html.template` 。它们的作用可以归纳为：

1. 模板文件通常定义为`.tmpl`和`.tpl`为后缀（也可以使用其他后缀），必须`utf8`编码
2. 模板文件中使用`{{`和`}}`包裹和标识需要传入的数据
3. 传给模板这样的数据就可以通过点号” `.` “ 来访问，如果数据是复杂类型的数据，可以通过`{{.FieldName}}`来访问它的字段。
4. 除`{{   }}`包裹的内容外，其他内容均不做修改原样输出

# 模板引擎的使用

分为定义模板文件、解析模板文件、模板渲染

## 定义模板文件：

其中，定义模板文件时需要我们按照相关语法规则去编写，后文会详细介绍

## 解析模板文件

上面定义好模板文件后，可以使用下面常用的方法去解析模板文件，得到模板对象：

```go
func (t *Template)Parse(src string)(*Template, error)
func ParseFiles(filenames ...string)(*Template, error)
func ParseGlob(paattern string)(*Template, error)
```

也可以使用`func New(name string) *Template` 函数创建一个名为name的模板，然后对其调用上面的方法去解析模板字符串或模板文件

## 模板渲染

简单来说就是使用数据去填充模板，当然实际上会复杂的多

```go
func (t *Template) Execute(wr io.Writer, data interface{})error
func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{})error
```

# 基本示例

## 模板定义

```html
<!DOCTYPE html>
<html lang="zh-CN">
<head>
		<title>Hello</title>
</head>
<body>
		<p>Hello {{ . }}</p>
</body>
</html>
```

## 解析和渲染模板

```go
package main
import (
		"net.http"
		"fmt"
		"html/template"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
		//解析：
		t,err := template.ParseFiles("./hello.tmpl")
		if err != nil {
				fmt.Println("Parse err:", err)
				return
		}
		//渲染：
		err := t.Execute(w, "Silver")
		if err != nil {
				fmt.Println("render template failed,err: ",err)
				return
		}
}

func main(){
		http.HandleFunc("/", sayHello)
		err := http.ListenAndServe(":9000", nil)
		if err != nil {
				fmt.Println("HTTP server start failed,err:",err)
				return
		}
}
```