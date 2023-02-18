package main

import (
	"fmt"
	"strings"
)

type Container struct {
	Map map[string]int
}

func NewContainer() *Container {
	return &Container{Map: make(map[string]int)}
}

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
