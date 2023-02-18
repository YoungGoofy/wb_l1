package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// создается канал
	ch := make(chan int)
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}
	// создается канал для остановки горутины
	timeout := time.After(time.Duration(n) * time.Second)
	r := rand.New(rand.NewSource(time.Now().UnixMicro()))
	// запуск горутины
	go func() {
		for {
			fmt.Println(<-ch)
		}
	}()
	for {
		select {
		// спустя n-ое время канал закрывает горутину
		case <-timeout:
			return
		default:
			// добавляет в канал случайное число
			ch <- r.Intn(100)
			time.Sleep(time.Second)
		}
	}
}

//func readData(ch chan int) {
//	for {
//		fmt.Println(<-ch)
//	}
//}
