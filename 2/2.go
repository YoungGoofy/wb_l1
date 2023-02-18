package main

import "fmt"

//func main() {
//	ch := make(chan int)
//	var arr = []int{2, 4, 6, 8, 10}
//	for _, item := range arr {
//		go row(ch, item)
//	}
//	for i := 0; i < len(arr); i++ {
//		item := <-ch
//		fmt.Println(item)
//	}
//}
//
//func row(ch chan int, item int) {
//	ch <- item * item
//}

func main() {
	// Создается канал
	ch := make(chan int)
	// Запускается новая горутина
	go row(ch)
	//for {
	//	item, ok := <-ch
	//	if !ok {
	//		break
	//	}
	//	fmt.Println(item)
	//}
	// Через цикл возвращаем данные, которые поступают в канал
	for item := range ch {
		fmt.Println(item)
	}
}

func row(ch chan int) {
	// проходимся по слайсу
	for _, item := range []int{2, 4, 6, 8, 10} {
		// В канал добавляем квадраты чисел
		ch <- item * item
	}
	// после завершения цикла канал закрывается
	close(ch)
}
