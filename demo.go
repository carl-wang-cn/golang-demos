package main

import (
	"fmt"
	"math"
	"reflect"
	"time"
)

func main() {
	// testValues()
	testVar()
	testConst()
	testFor()
	testSwitch()
	// testArray()
}

func testValues() {
	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1) // 通过逗号间隔的元素，print的时候，会在中间加一个空格
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println("true && false:", true && false)
	fmt.Println("true || false:", true || false)
	fmt.Println("!true:", !true)
}

func testVar() {
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

const s string = "constant"

func testConst() {
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

func testFor() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for x := 0; ; x += 1 {
		fmt.Println("loop")
		if x > 5 {
			break
		}
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

func testSwitch() {
	i := 4
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3: // 这里啥也print不出来哦，或者print个乱七八糟的东西，跟cpp的switch不太一样哦
	default:
		fmt.Println("default")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: // 用逗号分隔不同的值
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch { //一种替代if else的用法
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

func testArray() {
	var a [5]int
	fmt.Println("emp:", a) // 默认初始化为 [0 0 0 0 0]

	a[4] = 100 // 赋值
	fmt.Println("set:", a)
	fmt.Println("get: ", a[4])  // 取值
	fmt.Println("len:", len(a)) // array的长度

	b := [5]int{1, 2, 3, 4, 6} // 显式初始化
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}

// get the time of the specific timezone
func testTime() {
	loc, _ := time.LoadLocation("America/New_York")
	t := time.Now().In(loc)
	d := time.Date(t.Year(), t.Month(), t.Day(), 16, 0, 0, 0, loc)
	fmt.Println("t", t)
	fmt.Println("d", d)
}
