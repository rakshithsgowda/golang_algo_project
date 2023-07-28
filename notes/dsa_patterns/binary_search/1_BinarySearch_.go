package main

import "fmt"

// Recursive Binary Search
func binary_Search(element int,A []int)bool{
	low:=0
	high  := len(A)-1
	if low<=high{
		// mid point
		mid:=(high+low)/2
		if A[mid]>element{
			return binary_Search(element,A[:mid])
		}else if A[mid]<element{
			return binary_Search(element,A[mid+1:])
		}else{
			// if found retrun true
			return true
		}
	}
	// if the element is not found return false
	return false
}

func main(){
	items:=[]int{1,2,9,10,11,45,73,80,200}
	fmt.Println(binary_Search(45,items))
	fmt.Println(binary_Search(100,items))
}

