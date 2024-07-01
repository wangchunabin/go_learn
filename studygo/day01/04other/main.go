package main

import "fmt"

func variable() {
	fmt.Println("============声明变量==========")

	//声明变量
	var age int
	age = 18
	fmt.Println(age)

	//声明变量并初始化
	var name string = "小明"
	fmt.Println(name)

	//批量声明变量
	var (
		a int
		b string
		c bool
	)
	a = 1
	b = "test"
	c = false
	fmt.Println(a, b, c)

	//类型推导
	var d = 10
	var e = "test"
	var f = true
	fmt.Println(d, e, f)

	//短变量声明，局部变量
	n := 10
	m := 200
	fmt.Println(n, m)

	//匿名变量，在使用多重赋值时，如果想要忽略某个值，可以使用。 匿名变量用一个下划线_表示
	var x, _ = foo()
	var _, y = foo()
	fmt.Println("x=", x, "y=", y)
}

func constVariable() {
	//常量
	fmt.Println("===========常量==========")
	const pi = 3.1415926
	fmt.Println(pi)
	//批量设置
	const (
		ca = 100
		ce = 200
	)
	fmt.Println(ca, ce)

	fmt.Println("============常量计数器==========")

	const (
		n1 = iota //0
		n2        //1
		n3        //2
		n4        //3
	)
	fmt.Println(n1, n2, n3, n4)

	//使用_跳过某些值
	const (
		n11 = iota //0
		n22        //1
		_
		n44 //3
	)
	fmt.Println(n11, n22, n44)

	//iota声明中间插队,iota在const关键字出现时将被重置为0
	const (
		n111 = iota //0
		n222 = 100  //100
		n333 = iota //2
		n444        //3
	)
	const n5 = iota //0

	fmt.Println(n111, n222, n333, n444, n5)

	//定义数量级 （这里的<<表示左移操作，1<<10表示将1的二进制表示向左移10位，也就是由1变成了10000000000，也就是十进制的1024。同理2<<2表示将2的二进制表示向左移2位，也就是由10变成了1000，也就是十进制的8。）
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)
	fmt.Println(KB, MB, GB, TB, PB)

	//多个iota定义在一行
	const (
		a1, b1 = iota + 1, iota + 2 //1,2
		c1, d1                      //2,3
		e1, f1                      //3,4
	)
	fmt.Println(a1, b1, c1, d1, e1, f1)

}

func foo() (int, string) {
	return 10, "hello"
}

func main() {
	variable()
	constVariable()

}
