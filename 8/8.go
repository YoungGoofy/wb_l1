package main

import (
	"fmt"
	"log"
	"strconv"
)

// NewBit Перевод из числа в бит
func NewBit(num int) string {
	return strconv.FormatInt(int64(num), 2)
}

// NewNum Перевод из бита в число
func NewNum(bit string) int64 {
	newNum, err := strconv.ParseInt(bit, 2, 64)
	if err != nil {
		log.Println(err)
	}
	return newNum
}

// ChangeBit Замена бита в позиции
func ChangeBit(num, pos, bit int) int64 {
	// проверка позиции
	if pos < 1 {
		log.Println("Позиция меньше 1")
		return int64(num)
	}
	// проверка бита
	if bit < 0 || bit > 1 {
		log.Println("Может быть только 1 или 0")
		return int64(num)
	}
	// перевод числа в биты
	newBit := NewBit(num)
	// проверка позиции
	if pos > len(newBit) {
		log.Printf("Есть всего %d позиций.", len(newBit))
		return int64(num)
	}
	// перевод числового бита в биты, а затем в байты для корректной работы
	s := []byte(NewBit(bit))
	// перевод битов в байты
	r := []byte(newBit)
	// замена бита в позиции
	r[pos-1] = s[0]
	fmt.Println(string(r))
	// перевод из байтов в строку и из строки в число
	return NewNum(string(r))
}

func main() {
	fmt.Println(ChangeBit(12, 4, 1))
}
