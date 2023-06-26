package main;
import "fmt"

func binararr(arr [3]int, b []int){
	fmt.Println("function has been called",arr," + ",b,"=")
}

func main(){
	a := [3]int{10, 20, 30} 
	var b=[]int{1,2,3,4}
	fmt.Println(a)
	binararr(a,b)
}
