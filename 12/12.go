package main

import "fmt"

func main() {
	arr := [5]string{"cat", "cat", "dog", "cat", "tree"}
	mapa := map[string]bool{}

	for _, i := range arr {
		if !mapa[i] {
			mapa[i] = true
		}
	}

	array := make([]string, 0) // make([]string,len(mapa))

	for key := range mapa {
		array = append(array, key)
	}

	fmt.Println(array)
}
