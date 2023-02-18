package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Container создается структура с мьютексом
type Container struct {
	sync.Mutex
	wg sync.WaitGroup
	m  map[int]int
}

// NewContainer Создание нового контейнера
func NewContainer() *Container {
	return &Container{
		m: make(map[int]int),
	}
}

// Add Добавление значения в словарь
func (c *Container) Add(key, value int) {
	// Блокировка потока для остальных горутин
	c.Lock()
	// Разблокировка потока при завершении программы
	defer c.Unlock()
	c.m[key] = value
	fmt.Printf("Goroutine №%d\n", key)
	// Сигнал о завершении выполнения
	c.wg.Done()
}

// Print Чтение данных
func (c *Container) Print() {
	fmt.Println(c.m)
}

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixMicro()))
	var arr [10]int
	container := NewContainer()
	for i := 0; i < len(arr); i++ {
		arr[i] = r.Intn(1000)
	}
	// Создание группы из 10 горутин
	container.wg.Add(10)
	for index, item := range arr {
		go container.Add(index, item)
	}
	// ожидает завершение всех горутин
	container.wg.Wait()
	container.Print()
}
