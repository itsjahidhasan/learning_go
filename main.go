package main

import (
	// "learning_go/conditions"
	"fmt"
	problemSolve "learning_go/problem_solve"
)



func main(){
  	arr := []int{1,1,2,2,2,3,4,5,6,7,8}
	x:= 2
	
	found:=	problemSolve.BinarySearch(arr,x)
	fmt.Println("Index of",x,":",found)
}

 