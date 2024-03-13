package main

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" //import your used driver
)

// Model Struct
type User struct {
	UserId int    `orm:"pk"` //主键
	Name   string `orm:"size(100)"`
}

func init() { // init函数自动执行一次
	//数据库设置
	orm.RegisterDataBase("default", "mysql",
		"root:woshi1gg@tcp(127.0.0.1:3306)/Go-database?charset=utf8", 30)

	//注册model，可多个
	orm.RegisterModel(new(User))

	//创建表
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()
	user := User{Name: "aoho"}

	//插入
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	//更新
	user.Name = "boho"
	num, err := o.Update(&user)
	fmt.Printf("Num: %d, ERR: %v\n", num, err)

	//读取
	u := User{UserId: user.UserId}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	var maps []orm.Params
	res, err := o.Raw("SELECT * FROM user").Values(&maps)
	fmt.Printf("NUM: %d, ERR: %v\n", res, err)
	for _, term := range maps {
		fmt.Println(term["user_id"], ":", term["name"])
	}

	//删除
	num, err = o.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
