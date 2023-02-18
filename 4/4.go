package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//func Worker(ch chan int64, n int) {
//	for {
//		fmt.Printf("Worker №%d: %v\n", n, <-ch)
//	}
//}

func main() {
	// создается канал для чтения и записи
	ch := make(chan int64)
	// создается канал для остановки
	timeout := make(chan os.Signal)
	var count int
	_, err := fmt.Scan(&count)
	if err != nil {
		return
	}
	// инициализируется сигнал SIGTERM
	signal.Notify(timeout, syscall.SIGTERM)
	item := rand.New(rand.NewSource(time.Now().UnixNano()))
	// Запускается n-ое количество воркеров для чтения данных
	for i := 0; i < count; i++ {
		go func(i int) {
			for {
				fmt.Printf("Worker №%d: %v\n", i, <-ch)
			}
		}(i)
	}
	// запускается бесконечный цикл
	for {
		select {
		// после нажатия CTRL+C программа остановится
		case <-timeout:
			fmt.Println("\nSTOP")
			return
		default:
			// запись случайного числа в канал
			ch <- int64(item.Intn(15000))
			time.Sleep(time.Second)
		}
	}
}
