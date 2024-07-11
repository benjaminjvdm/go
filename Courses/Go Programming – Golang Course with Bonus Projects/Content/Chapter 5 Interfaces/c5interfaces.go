// Interfaces is a collection of method signatures
// This is all done implicitly as if you have a type interface and your structs fulfill the requirements of the interface it will be implementede to the interface  implicitly

package main

// the rect struct fulfills the interface as it has area and perimeter funcs on the width and height in the rect struct as area and perim funcs
type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

// area func
func (r rect) area() float64 {
	return r.width * r.height
}

// perim func
func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

//////// code assignment
