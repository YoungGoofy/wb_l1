package main

import (
	"fmt"
)

func main() {
	// создается канал
	ch1 := make(chan int)
	ch2 := make(chan int)
	// запускается 1 горутина
	go row(ch1)
	// запускается 2 горутина
	go sum(ch1, ch2)
	// вывод данных
	Print(ch2)
}

func row(ch1 chan int) {
	// запускается цикл
	for _, item := range []int{2, 4, 6, 8, 10} {
		// в канал передается квадрат числа
		ch1 <- item * item
	}
	// канад закрывается после окончания цикла
	close(ch1)
}
func sum(ch1, ch2 chan int) {
	// создается переменная
	res := 0
	// по циклу забирает данные из канала
	for item := range ch1 {
		fmt.Println(item)
		// сохраняет данные в переменную
		res += item
	}
	// после окончания цикла переносит данные во 2 канал
	ch2 <- res
}

func Print(ch chan int) {
	// вывод данных
	fmt.Println(<-ch)
}
