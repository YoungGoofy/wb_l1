package main

import "fmt"

// AddNumber Добавление чисел в канал
func AddNumber(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

// Multiplication Умножение чисел из 1 канала на 2 и отправка во 2 канал
func Multiplication(ch1, ch2 chan int) {
	for item := range ch1 {
		ch2 <- item * 2
	}
	close(ch2)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go AddNumber(ch1)
	go Multiplication(ch1, ch2)
	for item := range ch2 {
		fmt.Println(item)
	}
}
