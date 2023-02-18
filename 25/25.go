package main

import (
	"fmt"
	"time"
)

func main() {
	sleep(5 * time.Second)
	fmt.Println("hello world")
}

// Функция ждет, когда сработает канал, для продолжения работы
func sleep(d time.Duration) {
	<-time.After(d)
}
