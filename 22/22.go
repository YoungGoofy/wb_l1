package main

import (
	"fmt"
	"math/big"
)

type Math struct {
	A *big.Int
	B *big.Int
}

func NewMath(A, B string) *Math {
	newA := new(big.Int)
	newB := new(big.Int)
	newA.SetString(A, 10)
	newB.SetString(B, 10)
	return &Math{A: newA, B: newB}
}

func (m *Math) Add() *big.Int {
	return m.A.Add(m.A, m.B)
}

func (m *Math) Div() *big.Int {
	return m.A.Div(m.A, m.B)
}

func (m *Math) Sub() *big.Int {
	return m.A.Sub(m.A, m.B)
}

func (m *Math) And() *big.Int {
	return m.A.And(m.A, m.B)
}

func main() {
	m := NewMath("24000000000000000000", "2")
	//fmt.Println(m.Add())
	//fmt.Println(m.Sub())
	fmt.Println(m.Div())
	fmt.Println(m.And())
}
