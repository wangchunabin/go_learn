package main

import "fmt"

func main() {
	//十进制
	var a int = 10
	fmt.Printf("%d \n", a)
	fmt.Printf("%b \n", a)

	//八进制 以0开头
	var b int = 077
	fmt.Printf("%o \n", b)

	//十六进制 以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c)
	fmt.Printf("%X \n", c)

}
