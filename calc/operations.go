package calc

import (
	"fmt"
)

var operationsCount int

func init() {
	operationsCount = 0
}

func GetOperationsCount() int {
	return operationsCount
}

type Calculator interface {
	Sum(nums ...float64) float64
	Max(nums ...float64) float64
	Min(nums ...float64) float64
	Divide(a, b float64) (float64, error)
}

type Calc struct{}

func Sum(nums ...float64) float64 {
	var total float64
	for _, num := range nums {
		total += num
	}
	operationsCount++
	return total
}

func (c Calc) Sum(nums ...float64) float64 {
	return Sum(nums...)
}

func Max(nums ...float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	for _, num := range nums {
		if num > res {
			res = num
		}
	}
	operationsCount++
	return res
}

func (c Calc) Max(nums ...float64) float64 {
	return Max(nums...)
}

func Min(nums ...float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	for _, num := range nums {
		if num < res {
			res = num
		}
	}
	operationsCount++
	return res
}

func (c Calc) Min(nums ...float64) float64 {
	return Min(nums...)
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Помилка: ділення на нуль неможливе")
	}
	operationsCount++
	return a / b, nil
}

func (c Calc) Divide(a, b float64) (float64, error) {
	return Divide(a, b)
}
