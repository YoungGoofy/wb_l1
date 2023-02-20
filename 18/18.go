package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	workerPoolSize := 3
	ch := make(chan int, workerPoolSize)
	wgWrite := new(sync.WaitGroup)
	wgRead := new(sync.WaitGroup)
	wgRead.Add(1)
	counter := 0
	// Ждет когда в канал придет число, и после этого добавляет к счетчику + 1
	go CounterHandler(&counter, ch, wgRead)
	// Создает 3 горутины
	for i := 0; i < workerPoolSize; i++ {
		wgWrite.Add(1)
		go Proc("worker"+strconv.Itoa(i), ch, wgWrite)
	}
	wgWrite.Wait()
	close(ch)
	wgRead.Wait()
	fmt.Println("Общее количество итераций: ", counter)
}

func CounterHandler(counter *int, ch chan int, wg *sync.WaitGroup) {
	for range ch {
		*counter++
	}
	wg.Done()
}

// Proc
// Создается рандомное кол-во итераций
// Число отправляется в канал и вычитается - 1 из итерации
func Proc(id string, ch chan int, wg *sync.WaitGroup) {
	r := rand.New(rand.NewSource(time.Now().UnixMicro()))
	iterations := r.Intn(1000)
	for iterations > 0 {
		fmt.Printf("Работает %s. Осталось %d Итераций\n", id, iterations-1)
		ch <- 1
		iterations--
	}
	wg.Done()
}
