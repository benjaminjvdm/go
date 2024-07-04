package main

import (
	"fmt"
	"math"
)

func main(){
	i1,i2,i3 := 12,45,68
	intSum := i1 + i2 + i3
	fmt.Println("Int Sum is: ", intSum)

	f1,f2,f3 := 12.545,45.4456,68.45646
	floatSum := f1 + f2 + f3
	fmt.Println("Float Sum is: ", floatSum)
	roundfloatSum := math.Round(floatSum*100)/100
	// the above is needed to have 2 dec after period
	fmt.Println("Float Sum is: ", roundfloatSum)


}