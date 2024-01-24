package demo02
import "fmt"

type person struct{
	Name string//可以外部访问
	age int//私密
}

//工厂模式
func NewPerson (name string) *person{
	return &person{
		Name : name,
	}
}

//Set和Get函数，封装
func (p *person) SetAge(age int){
	if age > 0 && age < 150{
		p.age = age
	} else{
		fmt.Println("输入的年龄不合理")
	}
}

func (p *person) GetAge() int {
	return p.age
}