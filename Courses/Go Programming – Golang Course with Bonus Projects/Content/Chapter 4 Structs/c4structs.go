// structs are just collections (that's its type) of key:value pairs like py dicts
//  anon structs are only used once, stay clear of them and stick to named structs [type name struct]

package main

import (
	"fmt"
)

// created my first struct and I made it the profile of a Bladerunner. For the sake of the example serials are ints
type bladeRunner struct {
	serial        int
	contractsDone int
	fixer         string
	successRate   float64
}

// this function uses the above struct to call Fixer and Serial to fmt.Printf
func test(m bladeRunner) {
	fmt.Printf("Sending: '%s' to %v\n", m.fixer, m.serial)
	fmt.Println("=========================================>")
}

// this is a nested struct, we pass a new var type (user) which itself is a struct
type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

// we are writing a func that sees if we can send a message because 2 keys * 2 pairs equals 4 different parts of a user to validate
func canSendMesssage(mToSend messageToSend) bool {
	if mToSend.sender.name == "" {
		return false
	}
	if mToSend.recipient.name == "" {
		return false
	}
	if mToSend.sender.number == 0 {
		return false
	}
	if mToSend.recipient.number == 0 {
		return false
	}
	return true

}

// embeded structs are different to nested as these structs can have the entire struct in its struct
type car struct {
	make  string
	model string
}

type truck struct {
	car     // car is its own struct but we are putting it in the truck struct as some car manufacturers are truck manufacturers and "make" and "model" still applies here
	bedSize int
}

// embedded structs are great for public and private data as you can embed them

type authenticationInfo struct {
	username string
	password string
}

func (authI authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf(
		"Authorization: Basic %s:%s",
		authI.username,
		authI.password,
	)
}

func test(authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicAuth())
	fmt.Println("=====================>")
}

func main() {
	//using our struct and func we created for the bladeRunner
	test(bladeRunner{
		fixer:  "Wakako",
		serial: 1235453115,
	})
}
