package main

import "fmt"

//匿名函数

func main() {
	//将匿名函数保存到变量中
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(10, 20) //通过变量名调用匿名函数

	//自执行函数，匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)
}
