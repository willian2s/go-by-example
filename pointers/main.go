package main

import "fmt"

func pointersTour() {
	i, j := 42, 2701
	var pointer *int
	fmt.Printf("pointer zero value = %d\n", pointer)

	fmt.Printf("i = %d\n", i)
	fmt.Printf("j = %d\n", j)

	p := &i
	fmt.Printf("address i = %d\n", &i)
	fmt.Printf("address p = %d\n", p)
	fmt.Printf("value p = %d\n", *p)

	*p = 21
	fmt.Printf("value p = %d\n", *p)
	fmt.Printf("i = %d\n", i)

	p = &j
	fmt.Printf("address j = %d\n", &j)
	fmt.Printf("address p = %d\n", p)
	fmt.Printf("value p = %d\n", *p)
	*p = *p / 37
	fmt.Printf("value p = %d\n", *p)
	fmt.Printf("j = %d\n", j)
}

func zeroval(ival int) {
	ival = 0
}

func zeroprt(iptr *int) {
	*iptr = 0
}

func pointersExample() {
	i := 1
	fmt.Printf("initial: %d\n", i)

	zeroval(i)
	fmt.Printf("zeroval: %d\n", i)

	zeroprt(&i)
	fmt.Printf("zeroprt: %d\n", i)

	fmt.Printf("pointer: %d\n", &i)
}

type Person struct {
	name string
	age  int
}

func Birthday(person *Person) {
	person.age += 1
}

func main() {
	person := Person{name: "Willian", age: 27}
	fmt.Printf("Person = %v\n", person)
	Birthday(&person)
	fmt.Printf("Person = %v\n", person)
}
