package main

import "fmt"

func main() {
	str1 := []rune("главрыба")
	logRune(str1)
	fmt.Println()
	newStr := make([]rune, len(str1))

	for i := len(str1) - 1; i >= 0; i -= 1 {
		newStr[len(str1)-1-i] = str1[i]
	}

	logRune(newStr)
}

func logRune(run []rune) {
	for _, s := range run {
		fmt.Print(string(s))
	}
}
