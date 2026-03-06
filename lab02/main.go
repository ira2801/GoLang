package main

import (
	"fmt"
	"math"
)

// Частина 1: Слайси та Масиви

func task1() {
	fmt.Println("Завдання 1: Масиви та Слайси")

	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	result := make([]int, 10)

	for i := 0; i < 10; i++ {
		result[i] = a[i] + b[i]
	}

	fmt.Println("Результат:", result)

}

// Частина 2: Інтерфейси та Структури

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct{ Radius float64 }

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

type Rectangle struct{ Width, Height float64 }

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

type Triangle struct{ A, B, C float64 }

func (t Triangle) Perimeter() float64 { return t.A + t.B + t.C }
func (t Triangle) Area() float64 {
	p := t.Perimeter() / 2
	return math.Sqrt(p * (p - t.A) * (p - t.B) * (p - t.C))
}

func main() {
	task1()

	fmt.Println("Завдання 2: Фігури")
	shapes := []Shape{
		Circle{Radius: 8},
		Rectangle{Width: 10, Height: 5},
		Triangle{A: 2, B: 3, C: 4},
	}

	for _, s := range shapes {
		fmt.Printf("Фігура: %T | Площа: %f | Периметр: %f\n", s, s.Area(), s.Perimeter())
	}
}
