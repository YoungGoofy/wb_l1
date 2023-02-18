package main

import "fmt"

func main() {
	arr1 := []int{8, 12, 5, 7, 3, 6, 7, 15, 4}
	arr2 := []int{16, 8, 3, 4, 2, 12, 0, 4, 66}
	fullMap := make(map[int]int)
	// Добавление элементов в словарь
	for _, item := range arr1 {
		fullMap[item] = 1
	}
	for _, item := range arr2 {
		// Проверка есть ли элемент в словаре,
		// если есть, то +1,
		// иначе добавляется новый элемент в словарь
		if _, ok := fullMap[item]; ok {
			fullMap[item]++
		} else {
			fullMap[item] = 1
		}

	}
	var result []int
	// Добавляет в список все неповторяющиеся элементы
	for key, value := range fullMap {
		if value < 2 {
			result = append(result, key)
		}
	}
	fmt.Println(result)
}
