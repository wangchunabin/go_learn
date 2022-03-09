package main

import (
	"fmt"
)

func main() {
	//要求 可以从控制台接收用户信息 【姓名，年龄，薪水，是否通过考试】
	//方式1 fmt.Scanln
	var name string
	var age byte
	var sal float32
	var isPass bool
	fmt.Println("请输入姓名：")
	fmt.Scanln(&name)

	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)

	fmt.Println("请输入薪水：")
	fmt.Scanln(&sal)

	fmt.Println("请输入是否通过考试：")
	fmt.Scanln(&isPass)

	fmt.Printf("姓名：%v \n 年龄：%v \n 薪水：%v \n 是否通过考试：%v", name, age, sal, isPass)
}
