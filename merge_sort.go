package main

/*
	--Author: Vetri--
	-- Jun 28 2023 --

	Worst Case analysis::::
	Time Complexity N*log(N) log N for divide and conquer approach and N for sorting the array
*/
import (
	"fmt"
	"sort"
)

var no_elements int

func input() []int {
	fmt.Println("enter the number of elements")
	fmt.Scanln(&no_elements)

	fmt.Println("enter the elements in the array")
	var array = make([]int, no_elements)
	for i := 0; i < no_elements; i++ {
		fmt.Println("enter the element", i+1)
		fmt.Scanln(&array[i])
	}
	return array
}
func merge(arr []int) []int {
	if len(arr) > 1 {
		merge(arr[0 : len(arr)/2])
		merge(arr[(len(arr) / 2):len(arr)])
	}
	fmt.Println(arr)
	sort.Ints(arr)
	return arr

}
func main() {
	// fmt.Println("hello", input())
	fmt.Println(merge([]int{4, 3, 2, 1, 5, 7, 8}))
}
