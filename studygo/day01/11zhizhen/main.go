package main

import "fmt"

func main() {
	a := 10
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc00001a078
	fmt.Printf("b:%p type:%T\n", b, b) // b:0xc00001a078 type:*int
	fmt.Println(&b)                    // 0xc00000e018
	fmt.Println(*b)
	fmt.Println("-----------------")

	var c *int
	var d int = 4
	c = &d
	// c = d
	fmt.Println(c)
	fmt.Println(*c)
	fmt.Println(*&c)
	*c = 5
	fmt.Println(d)

	name := "zhangsan"
	modify1(name)
	fmt.Println(name)

	modify2(&name)
	fmt.Println(name)

}

func modify1(name string) {
	name = "lisi"
}

func modify2(name *string) {
	*name = "张三"
}
