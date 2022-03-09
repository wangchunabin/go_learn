package main

import (
	"fmt"
)

func main() {
	var name string
	var age byte
	var sal float32
	var isPass bool

	//方式2 按指定的方式输入
	fmt.Println("请输入姓名，年龄，薪水，是否通过考试，使用空格隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf("姓名：%v \n 年龄：%v \n 薪水：%v \n 是否通过考试：%v", name, age, sal, isPass)
}
