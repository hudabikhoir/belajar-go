package main

import (
	"fmt"
)

func main() {
	var numberA int = 4
	var numberB *int = &numberA
	// variale ke pointer
	fmt.Println("numberA (value) \t:", numberA)
	fmt.Println("numberA (address) \t:", &numberA)
	// pointer ke variable
	fmt.Println("numberB (value) \t:", *numberB)
	fmt.Println("numberB (address) \t:", numberB)
	fmt.Println("--------------------------------------")
	// lakukan perubahan nilai variable numberA
	numberA = 5
	fmt.Println("numberA (value) \t:", numberA)
	fmt.Println("numberA (address) \t:", &numberA)
	// pointer ke variable
	fmt.Println("numberB (value) \t:", *numberB)
	fmt.Println("numberB (address) \t:", numberB)
	fmt.Println("--------------------------------------")

	var number = 4
	fmt.Println("before \t:", number)

	change(&number, 10)
	fmt.Println("after \t:", number)
	fmt.Println("address \t:", &number)
}

// original *int berisikan address variable number
func change(original *int, value int) {
	// *original berisikan value dari address tersebut
	*original = value
}
