package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	arr, err := RemoveItem(arr, 1)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(arr)
}

func RemoveItem(arr []int, index int) ([]int, error) {
	if index >= len(arr) {
		return nil, errors.New(fmt.Sprintf("index out of range [%d] with length %d", index, len(arr)))
	}
	arr[index] = arr[len(arr)-1]
	return arr[:len(arr)-1], nil
}
