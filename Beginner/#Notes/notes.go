package main

import (
	"fmt"
)

func main() {
	var six uint = 777
	fmt.Println("\n", six)
}

/*
Characteristics
. Go is compiled and statically typed.
. Go tool can run a file without precompiling
. Executables are OS specific
. No VM is needed
. Runtine is manually linked
. Doesn't support type inheritance
. No method or operator overloading
. No exception handling, you need to use conditional logic to handle errors
. You need to explicitly convert numbers

Syntax Rules
. Case sensitive
. vars are mixed case
. no need for semi colons, but good lang spec
. wrap your code blocks in braces, but the first brace needs to be on the preceding statement line
*/

/*
. Check this for writing GO with tests - https://github.com/quii/learn-go-with-tests
*/

/*
Vars and stuff
. You either let the compiler implicitly guess the var type or exclusively set it (var name type)
. Integers have unsigned and signed types + number of bits need to be called (eg. uint64 vs int64) affecting the range of the var
. Aliases can be used instead of full name like byte, uint, int and uintptr(unsigned int pointer)
. Unsigned integers depend on OS and can be 32/64
. Floats are only in 32/64 versions
. Complex numbers are a combination of a real and imaginary part available in 32/64
. Arrays + Slices for ordered data collections
. Maps and Structs for aggr. of values
. function is a type and be passed into another as an arg
. Pointers are the reference variables in Go and points to an address like normal
. You cant change types at runtime
. %T is the placeholder for var types in a fmt.println situation
. remember to add line feed via \n because fmt print dont do that
. := is the faster var assigning method
. const can be initialized outside of the main func
*/

