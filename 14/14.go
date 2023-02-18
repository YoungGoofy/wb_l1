package main

import (
	"fmt"
	"reflect"
)

func main() {
	var variable string
	n := 12
	var b bool
	ch := make(chan interface{})
	ch2 := make(chan int)
	fmt.Printf("The type of variable is: %s\n", printType(variable))
	fmt.Printf("The type of variable is: %s\n", printType(n))
	fmt.Printf("The type of variable is: %s\n", printType(b))
	fmt.Printf("The type of variable is: %s\n", printType(ch))
	fmt.Printf("The type of variable is: %s\n", printType(ch2))
}

func printType(item interface{}) reflect.Type {
	return reflect.TypeOf(item)
}
