package main

import "fmt"

func main() {
	var arr = []int{1, 2}
	arr[0], arr[1] = arr[1], arr[0]
	fmt.Println(arr)
}
