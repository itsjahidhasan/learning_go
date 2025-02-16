package problemsolve

import (
	"fmt"
	"math"
)

//finding array index using binary search


func BinarySearch(arr []int, x int) int {
	fmt.Println("binary search")
	start :=0
	end := len(arr) - 1
	// mid := int(math.Ceil(float64((start+end)/2)))
	
	for start <= end {
		mid := int(math.Ceil(float64((start+end)/2)))
		if arr[mid] == x {
			if arr[mid-1] ==x{
				return mid-1
			}
			 return mid
		}else if arr[mid] < x{
			
            start = mid + 1

        
		}else{
                end = mid - 1
            }
		 
	}
	return 0;
}