package main

import "fmt"

//闭包-----搞不懂----------------------------------
func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	var f = adder()
	fmt.Println(f(10))

}
