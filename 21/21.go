package main

import "fmt"

type Adapter interface {
	Request()
}

type Container struct {
}

func (c Container) ReqContainer() string {
	return "Item"
}

type Result struct {
	*Container
}

func NewResult(c *Container) Adapter {
	return &Result{c}
}

func (r Result) Request() {
	fmt.Println(r.ReqContainer())
}

func main() {
	var c Container
	r := NewResult(&c)
	r.Request()
}
