package main

import "fmt"

type User struct {
	name  string
	email string
	age   int
}

func main() {
	user := User{name: "Jafar", email: "dhamivibez@gmail.com", age: 19}

	user.greet()
}

func (u User) greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", u.name, u.age)
}
