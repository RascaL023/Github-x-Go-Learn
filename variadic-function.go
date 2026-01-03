package main

import "fmt"

func getTotal(nums ...int) (total int) {
	for _, num := range nums {
		total += num
	}

	return total
}

func main(){
	fmt.Println("Total: ", getTotal(10, 10, 5, 3, 2))

	// Index kosong = slice !!!
	nums := []int{10, 20, 5}
	fmt.Println("Total: ", getTotal(nums...))

}
