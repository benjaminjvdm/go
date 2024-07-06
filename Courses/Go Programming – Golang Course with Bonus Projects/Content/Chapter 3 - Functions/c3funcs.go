// vars are passed by value not address
// vars being passed to funcs are copies!!! you have to assign the var and pass it to change the var (x := func(x))
// if there are multiple return values in a func you need to wrap them in brackets
// to ignore the other return values, you can just give it a "_" when invoking the function. Compiler will look the other way
// early returns / guard clauses - its when we return early within the function , mostly used for when error handlling and for conditionals as it is much more logical and linear. Instead of a nested tree it is a straight line and either it is yes or no...simple

package main

import "fmt"

func sub(x int, y int) int {
	//x,y are the ints, and the final int is the result
	return x - y
	// return (what we wanna do) - func signature
}

func concat(x, y string) string {
	//always tell Go what the types should be (err otherwise), you can actually declare the final type which will apply to all
	return x + y
}

func incrementContracts(contractsSoFar, contractsToAdd int) int {
	contractsSoFar = contractsSoFar + contractsToAdd
	return contractsSoFar
}

func getNames() (string, string) {
	return "Joe", "Mama"
}

func yearsUntilEvents(age int) (
	yearsUntilAdult int, // ALWAYS ALWAYS ALWAYS, MAKE SURE TO TO NAME THE RETURNS AND INPUTS, YOU WILL GET A UNDECLARED ERROR
	yearsUntilDrinking int,
	yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental // explicit return - giving all the names
	//return                                                          // implicit/naked return - letting the func decide
}

func main() {
	///using a concat func I wrote
	cantavoid1 := "Death +"
	cantavoid2 := " Taxes"
	fmt.Println(concat(cantavoid1, cantavoid2))

	//using an incrementing function I wrote to update the var with using the var when it was passed into the func to get an updated var
	contractsSoFar := 555
	const contractsToAdd = 111
	contractsSoFar = incrementContracts(contractsSoFar, contractsToAdd)
	fmt.Println("You've completed ", contractsSoFar, " contracts")

	//using a underscore to not have a error with regards to unused vars. Compiler looks the other way
	firstName, _ /*lastName*/ := getNames()
	fmt.Println("Welcome to Tyrone Industries, ", firstName)

}
