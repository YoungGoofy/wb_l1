package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	test1()
	test2()
	test3()
	test4()
}

// При добавлении значения в канал, горутина завершает работу
func test1() {
	timeout := make(chan bool)
	go func() {
		fmt.Println("Текст внутри горутины 1.")
		for {
			select {
			case <-timeout:
				return
			default:
				// Other code
			}
		}
	}()
	time.Sleep(time.Second)
	fmt.Println("Текст перед выходом горутины 1.")

	timeout <- true
}

// При истичении n-ого времени, горутина завершает свою работу
func test2() {
	timeout := time.After(time.Second)
	go func() {
		fmt.Println("Текст внутри горутины 2.")

	}()
	for {
		select {
		case <-timeout:
			fmt.Println("Текст перед выходом горутины 2.")
			return
		default:
			// Other code
		}
	}
}

// Завершение работы горутины с помощью сигнала с контекстом (не очень корректно)
func test3() {
	ch := make(chan string)
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	go func() {
		for {
			fmt.Println(<-ch)
			time.Sleep(time.Second)
		}
	}()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Текст перед выходом из горутины 3")
			cancel()
			return
		default:
			ch <- "Текст в горутине 3"
		}
	}
}

// завершение горутины с помощью сигналов (корректно)
func test4() {
	ch := make(chan string)
	timeout := make(chan os.Signal)
	signal.Notify(timeout, syscall.SIGTERM)
	go func() {
		for {
			fmt.Println(<-ch)
			time.Sleep(time.Second)
		}
	}()
	for {
		select {
		case <-timeout:
			fmt.Println("Текст перед выходом из горутины 4")
			return
		default:
			ch <- "Текст в горутине 4"
		}
	}
}
