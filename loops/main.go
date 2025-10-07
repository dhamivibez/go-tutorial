package main

import "fmt"

func main() {
	fmt.Println("Loops")
	var size int
	var nums = []int{}

	fmt.Print("How many numbers? ")
	fmt.Scanln(&size)

	for i := 0; i < size; i++ {
		var num int
		fmt.Printf("Enter number %d: ", i+1)
		fmt.Scanln(&num)
		nums = append(nums, num)
	}

	sum := 0

	for _, n := range nums {
		sum += n
	}

	average := float64(sum) / float64(len(nums))

	fmt.Printf("Numbers: %v", nums)
	fmt.Println()
	fmt.Println("Sum:", sum)
	fmt.Println("Average: ", average)

}
