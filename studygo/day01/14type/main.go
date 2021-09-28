package main

import "fmt"

//结构体
type person struct {
	name string
	city string
	age  int8
}

func main() {
	var p1 person
	p1.name = "王大锤"
	p1.city = "上海"
	p1.age = 18
	fmt.Printf("p1=%v\n", p1)  //p1={王大锤 上海 18}
	fmt.Printf("p1=%#v\n", p1) //p1=main.person{name:"王大锤", city:"上海", age:18}

	//匿名结构体
	var user struct {
		Name string
		Age  int
	}
	user.Name = "王大锤--"
	user.Age = 19
	fmt.Printf("%#v\n", user)

	//指针式结构体
	var p2 = new(person)
	fmt.Printf("%T\n", p2)     //*main.person
	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"", city:"", age:0}

	p2.name = "测试"
	p2.age = 18
	p2.city = "北京"
	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"测试", city:"北京", age:18}
}
