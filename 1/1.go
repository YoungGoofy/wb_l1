package main

import "fmt"

// Human Создается структура
type Human struct {
	Mouth string
	Hand  string
	Leg   string
}

// Action В структуру добавляется родительская структура
type Action struct {
	Human
}

func (h *Human) Speak() string {
	return "speaking"
}

func (h *Human) Move() string {
	h.Leg = "foot movement"
	h.Hand = "hand movement"
	return "Moving"
}

func main() {
	// Так как в структуре Action добавлена структура Human
	// то через структуру Action можно вызывать методы структуры Human
	var a Action
	fmt.Println(a.Speak())
	fmt.Println(a.Move())
	fmt.Println(a.Hand)
	fmt.Println(a.Leg)
}
