package main

import "fmt"

type Container struct {
	Data map[float64][]float64
}

func NewContainer() *Container {
	return &Container{Data: make(map[float64][]float64)}
}

// поиск максимального значнеия в списке
func max(arr []float64) float64 {
	m := arr[0]
	for _, item := range arr {
		if item > m {
			m = item
		}
	}
	return m
}

// поиск минимального значения в списке
func min(arr []float64) float64 {
	m := arr[0]
	for _, item := range arr {
		if item < m {
			m = item
		}
	}
	return m
}

// MaxDozen округление максимального числа до десятков
func MaxDozen(arr []float64) float64 {
	m := max(arr)
	for m > 10 {
		m -= 10
	}
	return max(arr) - m
}

// MinDozen округление минимального числа до десятков
func MinDozen(arr []float64) float64 {
	m := min(arr)
	for m < -10 {
		m += 10
	}
	return min(arr) - (10 + m)
}

// Fix Объединение данных в подмножества
func (c *Container) Fix(arr []float64) {
	for i := MinDozen(arr); i <= MaxDozen(arr); i += 10 {
		for _, item := range arr {
			if item < 0 && i < 0 {
				if item < i && item > i-10 {
					c.Data[i] = append(c.Data[i], item)
				}
			} else if item > i && item < i+10 {
				c.Data[i] = append(c.Data[i], item)
			}
		}
	}
}

func main() {
	var arr = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	container := NewContainer()
	container.Fix(arr)
	fmt.Println(container.Data)
}
