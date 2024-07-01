package main

import "fmt"

func gotoDemo0() {
	fmt.Println("goto(跳转到指定标签)")

	var a int = 10

LOOP:
	for a < 20 {
		if a == 15 {
			/* 跳过迭代 */
			a = a + 1
			goto LOOP
		}
		fmt.Printf("a的值为 : %d\n", a)
		a++
	}
}

func gotoDemo2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				//设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}

	return

breakTag:
	fmt.Println("结束for循环")
}

func main() {
	gotoDemo0()
	gotoDemo2()

}
