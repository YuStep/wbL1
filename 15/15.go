package main

import (
	"fmt"
	"strings"
)

var justString string

func createHugeString(s int) string {
	//создаем большую строку
	return strings.Repeat("s", s)
}

func someFunc() {
	//большая строка будет в переменной
	v := createHugeString(1 << 10)

	//слайс вместимостью 100 байт
	slice := make([]byte, 100)

	//копируем фрагмент огромной строки в заранее подготовленный слайс
	copy(slice, v[:100])
	justString = string(slice)
}

func main() {
	someFunc()
	fmt.Println(justString)
}
