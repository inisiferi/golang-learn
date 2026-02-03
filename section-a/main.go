package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

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

	var names = [4]string{"agus", "feri", "xx", "yyy"}
	for k, v := range names {
		fmt.Println(k, "=>", v)
	}

	var fruits = [...]string{"banana", "apple", "grape"}
	for _, v := range fruits {
		fmt.Println(v)
	}

	var animals = [2]string{"ikan", "embe"}
	fmt.Println(animals[1])

	var tools = []string{"knife", "spoon", "fork"}
	tools = append(tools, "cup")
	toolsA := tools[0:2]
	toolsB := tools[0:3:3]
	fmt.Println(toolsA)
	fmt.Println(toolsB)

	fruitTotals := map[string]int{"banana": 2, "orange": 10}
	fruitTotals["tomato"] = 1
	for k, v := range fruitTotals {
		fmt.Println(k, v)
	}

	m := map[string]User{}
	u := m["feri"] // copy value
	u.Age = 31
	m["feri"] = u // assign ulang

	fmt.Println(m["feri"])

	listUser := make(map[string]*User)
	listUser["feri"] = &User{Age: 1}
	listUser["danis"] = &User{Age: 10}
	fmt.Println(listUser)

	var chickens = []map[string]string{
		map[string]string{"name": "chicken blue", "gender": "male", "age": "young"},
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"], chicken["age"])
	}

	dogs := [][][]string{
		{
			{"hi", "helo"},
			{"welcome", "danke"},
		},
		{
			{"do", "ogs"},
			{"a", "B"},
		},
	}

	dogss := [][][][]string{
		dogs,
		dogs,
	}

	fmt.Printf("%v\n", dogss)

	fmt.Println(printMessage("Hello", []string{"Feri", "Danish"}))

	fmt.Println(randomValue(2, 10))

	return
}

type User struct {
	Age int
}

func printMessage(message string, names []string) string {
	return message + strings.Join(names, " ")
}

var randomizer = rand.New(rand.NewSource(time.Now().UnixNano()))

func randomValue(min int, max int) int {
	return randomizer.Int()%(max-min+1) + min
}
