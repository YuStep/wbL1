package main

import (
	"fmt"
	"strings"
)

func main() {

	//создаем новую строку и логаем ее
	base := "snow dog sun"
	fmt.Println(base)
	//создаем массив слов из исходной строки, разделенной пробелами и объявляем новую строку
	splited := strings.Split(base, " ")
	newString := ""

	//проходимся по нашему списку слов в обратном порядке и кидаем каждое новое слово в новую строку, добавляя пробел
	for i := len(splited) - 1; i >= 0; i -= 1 {
		newString += splited[i] + " "
	}

	//делаем трим чтобы убрать крайние пробелы, присваиваем исходной строке новое значение
	base = strings.TrimRight(newString, " ")
	fmt.Println(base)
}
