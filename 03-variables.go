package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var 用来声明变量
	var a = "initial"
	fmt.Println(a, reflect.TypeOf(a))

	var b, c int = 1, 2
	fmt.Println(b, c, reflect.TypeOf(b))

	var d = true // golang会自动推断变量类型
	fmt.Println(d, reflect.TypeOf(d))

	var e int // 默认给该类型的零值，e == 0
	fmt.Println(e, reflect.TypeOf(e))

	f := "apple" // 等价于 var f string = "apple"
	fmt.Println(f, reflect.TypeOf(f))
}
