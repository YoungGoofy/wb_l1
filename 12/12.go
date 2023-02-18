package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	var newArr []string
	for _, item := range arr {
		newArr = CheckUnique(newArr, item)
	}
	for _, item := range newArr {
		fmt.Printf("{%s} ", item)
	}
}

func CheckUnique(slice []string, value string) []string {
	// Проверяет, есть ли элемент в списке,
	// если есть, то возвращает список.
	for _, item := range slice {
		if item == value {
			return slice
		}
	}
	// если условие не сработало, возвращает обновленный список
	return append(slice, value)
}
