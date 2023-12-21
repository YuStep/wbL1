package main

import "fmt"

type Human struct {
	name string
	age  int
}

type Action struct {
	*Human
}

func main() {
	nprof := &Action{
		&Human{name: "Александр Шпак", age: 90},
	}
	fmt.Printf("New human | age: %d | name: %s\n", nprof.age, nprof.name)
	nprof.changeName("Александр Васильев (шпак)")
	nprof.changeAge(1900)
	fmt.Printf("New human | age: %d | name: %s", nprof.age, nprof.name)
}

func (a *Action) changeAge(change int) {
	a.age = change
}

func (a *Action) changeName(change string) {
	a.name = change
}
