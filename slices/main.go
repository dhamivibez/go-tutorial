package main

import "fmt"

func main() {
	fmt.Println("Slices")
	var firstSlice = []int{15, 2, 30, 20, 10}

	secondSlice := make([]int, len(firstSlice))
	copy(secondSlice, firstSlice)
	secondSlice[0] = 300

	fmt.Println("This is the first slice: ", firstSlice)
	fmt.Println("This is the second slice: ", secondSlice)

	// // doubleInPlace(slices)
	// fmt.Printf("Slices before appending squares: %v\n", slices)
	// slices = appendSquares(slices)
	// fmt.Printf("Slices after appending squares: %v \n", slices)

}

// func doubleInPlace(nums []int) {
// 	for i := range nums {
// 		nums[i] *= 2
// 	}

// 	fmt.Println(nums)
// }

// func appendSquares(nums []int) []int {
// 	for _, n := range nums {
// 		nums = append(nums, n*n)
// 	}
// 	return nums

// }
