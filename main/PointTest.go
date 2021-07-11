package main

import "fmt"

func main() {

	a := 123        // 变量a的值为123
	var b *int = &a // 变量b的值为a的内存地址
	fmt.Println("a:", a)
	// a: 123
	fmt.Println("b:", b)
	// b: 0xc00001e0a8

}
