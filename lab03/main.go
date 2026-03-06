package main

import (
	"fmt"
	"lab03/calc"
)

func main() {
	fmt.Println("Завдання 1")

	s := calc.Sum(5.5, 10, 4.5)
	fmt.Printf("Sum: %.2f\n", s)

	mx := calc.Max(12, 54, 2, 88)
	fmt.Printf("Max: %.2f\n", mx)

	mn := calc.Min(-5, 0, 10, -20)
	fmt.Printf("Min: %.2f\n", mn)

	d, err := calc.Divide(10, 4)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Divide: %.2f\n", d)
	}

	fmt.Println("Всього операцій:", calc.GetOperationsCount())

	fmt.Println()

	fmt.Println("Завдання 2")

	var myCalc calc.Calculator = calc.Calc{}

	fmt.Printf("Interface Sum: %.2f\n", myCalc.Sum(1, 2, 3, 4))
	fmt.Printf("Interface Max: %.2f\n", myCalc.Max(100, 500, 200))
	fmt.Printf("Interface Min: %.2f\n", myCalc.Min(0.5, 0.1, 0.9))

	res, err := myCalc.Divide(20, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Interface Divide: %.2f\n", res)
	}

	fmt.Println("Всього операцій:", calc.GetOperationsCount())
}
