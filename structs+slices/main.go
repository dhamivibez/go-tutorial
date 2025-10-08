package main

import "fmt"

type User struct {
	name string
	age  int
}

func main() {
	var size int
	users := []User{}

	fmt.Print("How many users do you want to add: ")
	fmt.Scanln(&size)

	for i := 0; i < size; i++ {
		var newName string
		var newAge int
		fmt.Printf("Enter name %d: ", i+1)
		fmt.Scanln(&newName)

		fmt.Printf("Enter %s's age: ", newName)
		fmt.Scanln(&newAge)

		newUser := User{name: newName, age: newAge}

		users = append(users, newUser)
	}

	var ageSum int

	for _, user := range users {
		ageSum += user.age
	}

	average := float64(ageSum) / float64(len(users))

	fmt.Println("The average age of the all users is", average)

}
