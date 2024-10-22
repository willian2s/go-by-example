package main

import "fmt"

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) *Person {
	p := Person{
		name: name,
		age:  age,
	}
	return &p
}

func main() {
	fmt.Println(Person{"Bob", 20})
	fmt.Println(Person{name: "Alice", age: 20})
	fmt.Println(Person{name: "Fred"})
	fmt.Println(&Person{name: "Ann", age: 40})
	fmt.Println(NewPerson("John", 18))

	// Instance a struct
	s := Person{name: "Carlos", age: 50}
	fmt.Println(s.name)

	// Pointer to struct
	sp := &s
	fmt.Println(sp.age)

	// Change value to 's'
	sp.age = 51
	fmt.Println(sp.age)
	fmt.Println(s.age)

	// Anonymous struct
	dog := struct {
		name   string
		isGood bool
	}{
		name:   "Rex",
		isGood: true,
	}
	fmt.Println(dog)
}
