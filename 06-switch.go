package main

import (
	"fmt"
	"time"
)

func main() {
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
			fmt.Println("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
