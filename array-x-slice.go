package main

import "fmt"

func arr(){
	var str_arr [3]string
	str_arr[0] = "Hello"
	str_arr[1] = "World"
	str_arr[2] = "Go"

	fmt.Printf(
		"%s %s From %s\n",
		str_arr[0], str_arr[1], str_arr[2],
	)

	num_arr := [3]int{1, 2, 3,}
	print_detail("Array", num_arr[:])
}

func slices(){
	nums := [5]int{ 1, 2, 3, 4, 5, }
	slice1 := nums[2:4]
	slice2 := nums[:4]

	slice1[1] = 10
	slice2 = append(slice2, 50)
	slice2 = append(slice2, 60)
	slice2[1] = 69

	newSlice := make([]int, 3, 5)
	copy(newSlice, slice2[:2])


	print_detail("Array ", nums[:])
	print_detail("Slice1", slice1)
	print_detail("Slice2", slice2)
	print_detail("NSlice", newSlice)
}

func print_detail(keyword string, arr []int){
	fmt.Println(keyword, ":", arr)
	fmt.Println("Len:", len(arr))
	fmt.Println("Cap:", cap(arr), "\n")
}

func main(){
	arr()
	slices()
}
