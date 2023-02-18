package main

import "fmt"

func main() {
	arr := []string{"hello", "world"}
	var result []string
	for i := len(arr) - 1; i >= 0; i-- {
		result = append(result, arr[i])
	}
	fmt.Println(result)
}
