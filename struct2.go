package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

func main() {
	// var pers1 Person
	// pers1.firstName = "Chris"
	// pers1.lastName = "woodward"

	// pers1 := Person{"Chris", "woodward"}

	pers1 := Person{lastName: "woodward", firstName: "Chris"}

	upPerson(&pers1)
	fmt.Printf("The name of the person is %s %s\n", pers1.firstName, pers1.lastName)

	pers2 := new(Person)
	pers2.firstName = "Chris"
	pers2.lastName = "woodward"
	// 解指针用法是合法的
	(*pers2).lastName = "woodward"
	upPerson(pers2)
	fmt.Printf("The name of the person is %s %s\n", pers2.firstName, pers2.lastName)

	pers3 := &Person{"Chris", "woodward"}
	upPerson(pers3)
	fmt.Printf("The name of the person is %s %s\n", pers3.firstName, pers3.lastName)
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}
