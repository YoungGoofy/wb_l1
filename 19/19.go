package main

import "fmt"

func main() {
	str := "hello world"
	var result string
	r := []rune(str)
	for i := len(str) - 1; i >= 0; i-- {
		result += string(r[i])
	}
	fmt.Println(result)
}
