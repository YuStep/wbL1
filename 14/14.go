package main

import (
	"fmt"
	"reflect"
)

func getType(val interface{}) {
	// Используем reflect.TypeOf для получения типа переменной
	t := reflect.TypeOf(val)

	// Выводим имя типа
	fmt.Printf("Тип переменной: %v\n", t)

	// Проверяем конкретный тип переменной
	switch t.Kind() {
	case reflect.Int:
		fmt.Println("Это int")
	case reflect.String:
		fmt.Println("Это string")
	case reflect.Bool:
		fmt.Println("Это bool")
	case reflect.Chan:
		fmt.Println("Это channel")
	default:
		fmt.Println("Неопределенный тип")
	}
}

func main() {
	// Примеры переменных разных типов
	var intValue int = 42
	var stringValue string = "Hello, World!"
	var boolValue bool = true
	var channelValue chan int

	// Определение типов переменных
	getType(intValue)
	getType(stringValue)
	getType(boolValue)
	getType(channelValue)
}
