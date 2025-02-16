package array

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 1024}

func RangeExample(){

	for i, value := range pow {
        fmt.Printf("2**%d = %d\n", i, value)
    }
}