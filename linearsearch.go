package main;

import "fmt";
/*	
	--Author: Vetri--
	-- Jun 19 2023 --
	Linear Search does not need to be sorted
	The array can be dynamic of static it depends based on the situation
	
	Worst Case analysis::::
		N (for data insertion) + N (for data search in last position) => 2N => O(N)
*/
func main(){
	
	var arrayLength int;
	fmt.Println("enter the array length")
	fmt.Scanln(&arrayLength)

	var array=make([]int,arrayLength)
	for i:=0 ; i<arrayLength ; i++ {
		fmt.Println("enter the value for the element %d",i)
		fmt.Scanln(&array[i])
	}
	
	fmt.Println(array)

	var searchElement int
	fmt.Println("enter the search element")
	fmt.Scanln(&searchElement)
	
	for i:=0; i<arrayLength; i++ {
		if array[i]==searchElement{
			fmt.Println("the element is found at pos  %d",i+1)
			break;
		}	
	}
}
