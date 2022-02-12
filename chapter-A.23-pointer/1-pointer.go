package main

import (
	"fmt"
	"reflect"
)

func main() {
	var numberA int = 4
	var numberB *int = &numberA
	var numberC *int = numberB
	var numberD **int = &numberB

	var reflectValue = reflect.ValueOf(numberD)
	fmt.Println("tipe  variabel :", reflectValue.Type())

	fmt.Println("numberA (value)   :", numberA)  // 4
	fmt.Println("numberA (address) :", &numberA) // 0xc0000140a8

	fmt.Println("numberB (value)   :", *numberB) // 4
	fmt.Println("numberB (address) :", numberB)  // 0xc0000140a8

	fmt.Println("numberC (value)   :", *numberC) // 4
	fmt.Println("numberC (address) :", numberC)  // 0xc0000140a8

	fmt.Println("numberD (value)   :", *numberD) // 0xc0000140a8
	fmt.Println("numberD (address) :", numberD)  // 0xc00000e028
}
