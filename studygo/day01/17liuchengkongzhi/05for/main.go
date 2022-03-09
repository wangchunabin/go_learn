package main

import "fmt"

func main() {
	s := "abc"

	for i, n := 0, len(s); i < n; i++ {
		fmt.Println(s[i])
	}

	n := len(s)
	for n > 0 {
		n--
		println(s[n])
	}

	//九九乘法表
	//i表示乘数
	var num int = 9
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v \t", j, i, j*i)
		}
		fmt.Println()
	}

}
