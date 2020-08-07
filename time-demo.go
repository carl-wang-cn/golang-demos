// get the time of the specific timezone

package main

import (
	"fmt"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("America/New_York")
	t := time.Now().In(loc)
	d := time.Date(t.Year(), t.Month(), t.Day(), 16, 0, 0, 0, loc)
	fmt.Println("t", t)
	fmt.Println("d", d)
}
