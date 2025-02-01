package array

import (
	"fmt"
)	

func PrintSlice (s []int){
	fmt.Printf("len: %d  cap: %d value: %v \n", len(s),cap(s), s)
}