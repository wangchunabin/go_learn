package main

import "fmt"

func intSum(x int, y int) int {
	return x + y
}

//类型简写
func intSub(x, y int) int {
	return x - y
}

func sayHello() {
	fmt.Println("Hello 沙河！")
}

//可变参数
func intSum2(x ...int) int {
	fmt.Println(x) //x是一个切片
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

//固定参数搭配可变参数使用时，可变参数要放在固定参数的后面
func intSum3(x int, y ...int) int {
	fmt.Println(x, y)
	sum := x
	for _, v := range y {
		sum = sum + v
	}
	return sum
}

//多返回值
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

func main() {
	//函数调用
	// sayHello()
	// ret := intSum(10, 20)
	// fmt.Println(ret)

	// sub := intSub(20, 10)
	// fmt.Println(sub)

	//调用
	ret1 := intSum2()
	ret2 := intSum2(10)
	ret3 := intSum2(10, 20)
	ret4 := intSum2(10, 20, 30)
	fmt.Println(ret1, ret2, ret3, ret4)

	//调用
	// ret5 := intSum3(100)
	// ret6 := intSum3(100, 10)
	// ret7 := intSum3(100, 10, 20)
	// ret8 := intSum3(100, 10, 20, 30)
	// fmt.Println(ret5, ret6, ret7, ret8) //100 110 130 160

	// ret9, ret10 := calc(20, 10)
	// fmt.Println(ret9, ret10)

}
