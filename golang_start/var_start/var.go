package main

import (
	"fmt"
)

func main() {
	var (
		num int
		floatval float64
		str string
		boolval bool
		byteval byte
		runeval rune
	)
	// var(
	// 	inc int = 1 << (iota*10)
	// 	inc1
	// 	inc2
	// )
	// fmt.Println(inc, inc1, inc2)
	// 不允许以上行为，iota只能在const类型中使用
	const (
		b int = 1 << (iota*10)
		kb
		mb
	)// 编译期算定
	fmt.Println(floatval,str, boolval, byteval, runeval)
	fmt.Println("num:", num)
	fmt.Println("b:", b)
	fmt.Println("kb:", kb)
	fmt.Println("mb:", mb)
	fmt.Printf("%q 带引号 %T 类型 %t 布尔\n", "hi", 3.14, true)
}
