package main

import "fmt"

func main() {
	s := "abc"
	// 忽略 2nd value，支持 string/array/slice/map。
	for i := range s {
		fmt.Println(s[i])
	}

	//忽略index
	for _, c := range s {
		println(c)
	}

	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		println(k, v)
	}
	//range 会复制对象。
	a := [3]int{0, 1, 2}
	for i, v := range a {
		if i == 0 {
			a[1], a[2] = 999, 999
			fmt.Println(a)
		}
		a[i] = v + 100
	}
	fmt.Println(a)
}
