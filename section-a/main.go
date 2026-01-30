package main

import "fmt"

func main() {
	fmt.Println("===Chapter-A.7-8===")
	fmt.Println("Hello World!")
	//commment hello inline
	/*
	 multi line comment
	*/

	fmt.Println("\n===Chapter-A.9===")
	var firstName string = "john"
	var lastName string
	lastName = "Wick"
	middleName := "Saputro"

	fmt.Printf("halo %s %s %s", firstName, middleName, lastName)

	var one, two, tree int = 1, 3, 4
	fmt.Println("\n", one, two, tree)

	name := new(string) 
	*name = "john wick"
	fmt.Println("name:", name)
	fmt.Printf("name: %s\n", *name)

	return
}
