package main

/*
	--Author: Vetri--
	-- Jun 21 2023 --

	Worst Case analysis::::
		N for outer loop and N for inner loop => N*N
*/
import "fmt"

func insert(arr []int, size int) []int {
	var i, j, k int
	for i = 0; i < size; i++ {
		k = i
		for j = k - 1; j > -1; j-- {
			if arr[k] < arr[j] {
				temp := arr[k]
				arr[k] = arr[j]
				arr[j] = temp
				k--
			}
		}

	}
	return arr
}
func main() {
	fmt.Println("enter the number of elements")
	var no_elements int
	fmt.Scanln(&no_elements)

	fmt.Println("enter the elements in the array")
	var array = make([]int, no_elements)
	for i := 0; i < no_elements; i++ {
		fmt.Println("enter the element", i+1)
		fmt.Scanln(&array[i])
	}
	fmt.Println(insert(array, no_elements))
}
