package main

import (
	"log"
	"math/rand"
)

func main() {
	someFunc()
	println(justString)
}

var justString string

func someFunc() {
	v := createHugeString(1 << 20) //1024
	// никак не контролируется правая граница слайса, на примере с размером строки < 100, словим панику по выходу
	// за границу
	rightIndex := 100
	switch {
	case rightIndex <= len(v)-1:
		justString = v[:100]
	default:
		log.Println("func SomeFunc(). Некорректная правая граница слайса, выход за пределы строки")
	}
}

// Реализовал данную функцию, т.к. её не было, что приводило к ошибке компиляции
func createHugeString(size int) string {
	res := make([]byte, size)
	for n := range res {
		res[n] = byte(rand.Intn(122-48) + 48)
	}
	return string(res)
}
