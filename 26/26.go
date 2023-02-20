package main

import (
	"fmt"
	"strings"
)

type Container struct {
	Map map[string]int
}

// NewContainer Конструктор
func NewContainer() *Container {
	return &Container{Map: make(map[string]int)}
}

// CheckUnique
// Проходимся циклом по данному слову
// Если данные повторяются, то в словаре счетчик будет больше 1, соответственно false
func (c *Container) CheckUnique(word []byte) bool {
	for _, b := range word {
		c.Map[string(b)]++
	}
	for _, value := range c.Map {
		if value > 1 {
			return false
		}
	}
	return true
}

func main() {
	container := NewContainer()
	word := "abCdefAaf"
	b := []byte(strings.ToLower(word))
	fmt.Println(container.CheckUnique(b))
}
