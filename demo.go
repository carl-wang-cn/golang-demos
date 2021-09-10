package main

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"time"
)

func main() {
	// testValues()
	// testVar()
	// testConst()
	// testFor()
	// testSwitch()
	// testArray()
	// testSlice()
	// testMap()
	// testRange()
	// testClosures()
	// testRecursiveClosures()
	// testPointer()
	// testStruct()
	// testMethods()
	// testInterfaces()
	// testError()
	// testCoro()
	// testChannel()
	// testNonBlockingChannel()
	// testClosingChannel()
	testRangeOverChannel()
	// testSelect()
	// testTimeout()
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
	var a [5]int           // 一定要写明len和type，slice只需要type
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

func testSlice() {
	// slice的定义跟array的不同，不需要指定len
	emp := []string{}
	fmt.Println("emp:", emp)
	fmt.Println("len(emp):", len(emp))

	s := make([]string, 3)
	fmt.Println("empty s:", s)

	t := []string{"g", "h", "i"}
	fmt.Println("t:", t)

	s[0] = "a"
	s[1] = "b"
	fmt.Println("s:", s)
	fmt.Println("get s[1]:", s[1])
	fmt.Println("len(s):", len(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("c copy from s:", c)
	fmt.Println("len(c):", len(c))

	s = append(s, "d")
	s = append(s, "e")
	s = append(s, "f")
	fmt.Println("apd d, e, f into s, s is now:", s)

	fmt.Println("s[2:5]:", s[2:5]) // 前闭后开
	fmt.Println("s[:5]:", s[:5])   // 前闭后开
	fmt.Println("s[2:]:", s[2:])   // 前闭后开

	twoD := make([][]int, 3) // 3个元素的slice，每个元素的类型为[]int类型的slice
	for i := 0; i < 3; i++ {
		innerLen := i + 1 // 跟array不同，内层的slice可以有不同的len
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

}

func testMap() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map m:", m)

	fmt.Println("m[\"k1\"]:", m["k1"])
	fmt.Println("len(m):", len(m))

	delete(m, "k2")
	fmt.Println("delete k2, then map m:", m)

	// 通过第2个返回值来判断该key是否在map中存在
	// 获取一个不存在的key，会返回一个空值，0或者""，无法判断是真的不存在还是值是0
	_, is_exist := m["k2"]
	fmt.Println("is_exist:", is_exist)

	n := map[string]int{"foo": 1, "name": 2}
	fmt.Println("map n:", n)

}

func testRange() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index that value == 3:", i)
		}
	}
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

// closures start
func intSeq() func() int { // 返回一个 func() int 类型的函数，无输入参数，返回值为int
	i := 0
	return func() int {
		i++
		return i
	}
}

func testClosures() {
	nextInt := intSeq()    // nextInt是一个闭包，里面有个i，i的生命周期没随着它的作用域结束而结束
	fmt.Println(nextInt()) // 连续调用3次同一个闭包
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	fmt.Println(intSeq()()) // 调用了3个不同的闭包，intSeq()每次都返回一个新的闭包
	fmt.Println(intSeq()())
	fmt.Println(intSeq()())

	newInt := intSeq()
	fmt.Println(newInt())
}

// closures end

func testRecursive(n int) int {
	if n == 0 {
		return 1
	}
	return n * testRecursive(n-1)
}

func testRecursiveClosures() { // 这跟闭包有毛线关系？不就是func定义在局部？
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}

// pointer start
// 用法跟c++基本一致的
func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
	fmt.Println("iprt addr:", iptr)
}

func testPointer() {
	i := 1
	fmt.Println("initial i:", i)
	fmt.Println("i addr:", &i)

	zeroval(i)
	fmt.Println("after call zeroval(i), i:", i)

	zeroptr(&i)
	fmt.Println("after call zeroptr(&i), i:", i)

}

// pointer end

// struct start
type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func testStruct() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})          // 未初始化的元素初始化未0
	fmt.Println(&person{name: "Ann", age: 40}) // 返回指针
	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age) // 指针也是用 . 来访问

	sp.age = 51 // 可以用指针来修改元素的值
	fmt.Println(sp.age)
}

// struct end

// methods start
// 有点像面向对象的class的类方法
type rect struct {
	width, height float64
}

func (r *rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func testMethods() {
	r := rect{width: 10, height: 3}

	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}

// methods end

// interfaces start
type geometry interface {
	area() float64
	perim() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println("g:", g)
	fmt.Println("g area:", g.area())
	fmt.Println("g perim:", g.perim())
}

func testInterfaces() {
	r := rect{3, 4}
	c := circle{5}

	// measure(r) // 这个是不行的
	measure(&r)
	measure(&c)
}

// interfaces end

// error start
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil // nil表示没错误
}

// 自定义error对象
// 通过给一个struct定义一个Error() method来实现
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		// return -1, &argError{arg, "can't work with it"}
		return -1, &argError{arg, "can't work with it"} // ? 这不是个struct？跟error怎么转换的？
	}
	return arg + 3, nil
}

func testError() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println("ae:", ae)
		fmt.Println("ok:", ok)
		fmt.Println("ae.arg:", ae.arg)
		fmt.Println("ae.prob:", ae.prob)
	}
}

// error end

// coroutines start
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func testCoro() {
	f("direct")

	go f("goroutine") // 启动一个coro去运行

	go func(msg string) {
		fmt.Println(msg)
	}("going") // 匿名函数，通过coro去运行

	time.Sleep(time.Second)
	fmt.Println("done")
}

// coroutines end

// channel start
func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

// 声明channel pings只能用来往channel内 sending string
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// 声明channel pings只能用来往channel外吐sting，channel pongs只能用来往channel内 sending string
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func testChannel() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
		messages <- "ping"
	}()

	msg := <-messages + <-messages
	fmt.Println(msg)

	done := make(chan bool)
	go worker(done)

	<-done // 没有这行，就不会等着worker执行完成

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "pass message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func testNonBlockingChannel() {
	// 默认在channel上收发消息，是阻塞的
	// 我们可以通过select来实现非阻塞的收发

	message := make(chan string)
	signal := make(chan bool)

	select {
	case msg := <-message:
		fmt.Println("recved message", msg)
	default:
		fmt.Println("no msg recved")
	}

	msg := "hi"
	select { // 因为message无receiver，所以msg发不出去
	case message <- msg:
		fmt.Println("sent msg", msg)
	default:
		fmt.Println("no msg sent")
	}

	select {
	case msg := <-message:
		fmt.Println("recved msg", msg)
	case sig := <-signal:
		fmt.Println("recved signal", sig)
	default:
		fmt.Println("no activity")
	}
}

func testClosingChannel() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("recved job", j)
			} else {
				fmt.Println("recved all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	close(jobs)
	fmt.Println("sent all jobs")

	<-done // 等着所有任务完成，收到true
}

func testRangeOverChannel() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}

// channel end

func testSelect() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(3 * time.Second)
		c1 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

func testTimeout() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second): // 超时时间1s
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second): // 超时时间3s
		fmt.Println("timeout 2")
	}
}

// get the time of the specific timezone
func testTime() {
	loc, _ := time.LoadLocation("America/New_York")
	t := time.Now().In(loc)
	d := time.Date(t.Year(), t.Month(), t.Day(), 16, 0, 0, 0, loc)
	fmt.Println("t", t)
	fmt.Println("d", d)
}
