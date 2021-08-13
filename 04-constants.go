package main

import (
	"fmt"
	"math"
	"reflect"
)

const s string = "constant"

func main() {
	// const 用来声明常量
	// Go supports constants of character, string, boolean, and numeric values.

	fmt.Println(s)

	const n = 5000000
	fmt.Println(n, reflect.TypeOf(n))

	const d = 3e20 / n
	fmt.Println(d, reflect.TypeOf(d))

	fmt.Println(int64(d), reflect.TypeOf(int64(d)))
	fmt.Println(math.Sin(n), reflect.TypeOf(n))
}
