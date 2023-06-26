package main
import "fmt"
/*	
	--Author: Vetri--
	-- Jun 20 2023 --
	Binar Search needs to have sorted array
	The array can be dynamic of static it depends based on the situation
	
	Worst Case analysis::::
		N (for data insertion) + log N (for data search in last position) => 2N => O(N)
*/

func search(arr []int,searchElement int,start int,end int)int{
	if arr[(start+end)/2]==searchElement {
		return (start+end)/2
	} else if arr[(start+end)/2]<searchElement {
		return search(arr,searchElement,start+end/2,end)

	} else {
		return search(arr,searchElement,start,start+end/2)
	}
	return -1
}

func main(){
	fmt.Println("enter the number of elements")
	var no_elements int
	fmt.Scanln(&no_elements)
	
	fmt.Println("enter the elements in the array")
	var array=make([]int,no_elements)
	for i:=0 ; i<no_elements ; i++{
		fmt.Println("enter the element",i+1,)
		fmt.Scanln(&array[i])
	}

	fmt.Println("enter the element to search")
	var element int
	fmt.Scanln(&element)
	fmt.Println("found at position ",search(array,element,0,len(array)-1)+1)
	
}