package main

import (
	"fmt"
	"time"
)

func main() {
	// n := time.Now()
	// fmt.Println("The time right now is: ", n)

	t := time.Date(2009, time.November, 10, 23, 0,0,0, time.UTC)
	fmt.Println("The 't' is: ", t)
	fmt.Println(t.Format(time.ANSIC))

}
