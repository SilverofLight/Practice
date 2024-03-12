package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type User struct {
	id          int
	name        string
	habits      string
	createdTime string
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("login.tql")
		if err != nil {
			log.Fatal("Parse err: ", err)
		}
		log.Println(t.Execute(w, nil))
	} else {
		_ = r.ParseForm()
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])
		if pwd := r.Form.Get("password"); pwd == "123456" {
			fmt.Fprintf(w, "Hello, %s", r.Form.Get("username"))
		} else {
			fmt.Fprintf(w, "密码错误")
		}
	}
}

func main() {
	var err error
	db, err = sql.Open("mysql",
		"root:woshi1gg@tcp(127.0.0.1:3306)/Go-database?charset=utf8")
	checkErr(err)
	http.HandleFunc("/login", login)
	http.HandleFunc("/info", userInfo)
	err1 := http.ListenAndServe(":8080", nil)
	if err1 != nil {
		log.Fatal("ListenAndServe err: ", err1)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal("err: ", err)
	}
}

func queryByName(name string) User { //按照用户名查询
	user := User{}
	stmt, err := db.Prepare("select * from user where name=?")
	if err != nil {
		log.Println("prepare error: ", err)
		return user
	}
	defer stmt.Close()

	rows, err := stmt.Query(name)
	if err != nil {
		log.Println("Query error: ", err)
		return user
	}
	defer stmt.Close()

	fmt.Println("\nafter deleting records: ")
	for rows.Next() {
		var id int
		var name string
		var habits string
		var createdTime string
		err := rows.Scan(&id, &name, &habits, &createdTime)
		if err != nil {
			log.Println("Scan error: ", err)
			return user
		}
		fmt.Printf("{%d, %s, %s, %s}\n", id, name, habits, createdTime)
		user = User{id, name, habits, createdTime}

	}
	return user
}

func store(user User) {
	stmt, err := db.Prepare("INSERT INTO user SET name=?,habits=?,created_time=?")
	if err != nil {
		log.Println("Prepare error: ", err)
		return
	}
	defer stmt.Close()

	t := time.Now().UTC().Format("2006-01-02")
	res, err := stmt.Exec(user.name, user.habits, t)
	if err != nil {
		log.Println("Execution error: ", err)
		return
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("RowsAffected error: ", err)
		return
	}
	fmt.Println("Rows affected: ", rowsAffected)
}

func userInfo(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	if r.Method == "POST" {
		user1 := User{
			name:   r.Form.Get("username"),
			habits: r.Form.Get("habits"),
		}
		store(user1)
		fmt.Fprintf(w, "%v", queryByName(user1.name))
	}
}
