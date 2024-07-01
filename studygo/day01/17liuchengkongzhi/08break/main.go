package main

import "fmt"

//单循环跳出
func demo0() {
	for i := 0; i < 10; i++ {
		if i == 2 {
			break
		}
		fmt.Printf("%v\n", i)
	}
	fmt.Println("。。。。。。")
}

//多循环跳出
func demo1() {
BREAKDEMO1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break BREAKDEMO1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	fmt.Println("...")
}

func continueDemo() {
forloop1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
}

//九九乘法口诀表
func cfb() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%2d ", j, i, i*j) // %2d是为了对齐输出
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println("break(跳出循环)")
	// demo0()
	// demo1()
	// continueDemo()
	cfb()

}
