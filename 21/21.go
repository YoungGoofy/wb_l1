package main

import "fmt"

// Adapter
// Создается интерфейс адаптер
type Adapter interface {
	Request()
}

// Container
// Существующий класс, не удовлетворяющий интерфейсу
type Container struct {
}

// ReqContainer
// Метод существующего класса
func (c Container) ReqContainer() string {
	return "Item"
}

// Result
// Создается новая структура
type Result struct {
	*Container
}

// NewResult
// Создается конструктор, который возврвщает интерфейс в теле структуры
func NewResult(c *Container) Adapter {
	return &Result{c}
}

// Request
// Создается метод, удовлетворяющий интерфейсу
func (r Result) Request() {
	fmt.Println(r.ReqContainer())
}

// Данный паттерн работает, как обертка для существующих классов
func main() {
	var c Container
	r := NewResult(&c)
	r.Request()
}
