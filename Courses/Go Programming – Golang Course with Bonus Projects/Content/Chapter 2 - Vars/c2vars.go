// unsigned (uint) have no negative int only positive
// 8,16,32,64 - how much data can be stored, if you want to be performant based, use the smallest bit needed
//best number vars to stick with = int, unint, float64 (all 3 will do most the heavy lifting )
// const is to declare a constant (sub price, plan etc) it is immutable, but can be computed (math)

package main

import (
	"fmt"
)

func main() {
	// //init vars (hard way)
	// var smsSendingLimt int // 1
	// var costPerSMS float64 //1.0
	// var hasPermission bool
	// var username string // ""

	// // //init vars (infer)
	// model := "Nexus 9 -"
	// serial := " KD6-3.7"
	// congrats := "Happy Birthday"
	// fmt.Println(congrats, model, serial)

	// // //multi-assign
	// fixer, jobNum, sc := "Wakako", 48, 666
	// fmt.Println("Your current fixer is: ", fixer, ". This is job number: ", jobNum, ". Your street cred will increase by: ", sc)

	// // //truncation
	// accountAge := 2.6
	// accountAgeInt := int(accountAge)
	// fmt.Println("Your account is ", accountAgeInt, " years old. And has exactly ", accountAge, "years of account history")

	//constants (cont)
	const BladeRunnerID = "369"
	const NexusModel = "NM9"
	const IntelStats = "9.99"
	const successrate = 98.76
	const Serial = BladeRunnerID + NexusModel + IntelStats
	// fmt.Println("Your serial number is: ", Serial)

	//prints
	msg := fmt.Sprintf( // we are assigning a formatted string to msg and then printing it with fmt.Println
		"Hi %s, your Success Rate is %.1f percent",
		Serial,
		successrate,
	)
	fmt.Println(msg)

	//conditionals (if > (else if) > else) the curly brackets are almost like a pipeline
	contractValue := 2000
	minContractValue := 200
	if contractValue >= minContractValue {
		fmt.Println("Contract Accepted by: ", Serial)
	} else {
		fmt.Println("Contract Declined by: ", Serial)
	}
	// go booger sugar conditional shortening
	if contractValue, minContractValue := 2000, 200; contractValue > minContractValue {
		fmt.Println("Contract Accepted By: ", Serial)
	}

}
