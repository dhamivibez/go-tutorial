package main

import "fmt"

func main() {
	fmt.Println("Maps")
	// var size int
	users := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"John":  22,
	}

	for name, age := range users {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	// fmt.Print("How many users do you want to add? ")
	// fmt.Scanln(&size)

	// for i := 0; i < size; i++ {
	// 	var name string
	// 	var age int
	// 	fmt.Printf("What is the name user %d? ", i+1)
	// 	fmt.Scanln(&name)
	// 	fmt.Printf("What is the age of %s? ", name)
	// 	fmt.Scanln(&age)
	// 	users[name] = age
	// }

	// var query string

	// fmt.Print("Whose age do you want to get? ")
	// fmt.Scanln(&query)

	// age, exists := users[query]

	// if !exists {
	// 	fmt.Printf("%s is not in the list of users", query)
	// 	return
	// }

	// fmt.Printf("Age of %s is %d\n", query, age)

}
