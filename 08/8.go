package main

import (
	"fmt"
)

func setBit(num int64, pos uint, value bool) int64 {
	if value {
		// Установить i-й бит в 1
		return num | (1 << pos)
	} else {
		// Установить i-й бит в 0
		return num &^ (1 << pos)
	}
}

func main() {
	// Пример переменной int64
	var num int64 = 42
	fmt.Printf("Исходное значение: %d\n", num)

	// Установить 3-й бит в 1
	num = setBit(num, 3, true)
	fmt.Printf("Установлен 3-й бит в 1: %d\n", num)

	// Установить 5-й бит в 0
	num = setBit(num, 5, false)
	fmt.Printf("Установлен 5-й бит в 0: %d\n", num)
}
