package array

func AppendExample(){
	var s []int 
	PrintSlice(s)

	s = append(s, 0)
	PrintSlice(s)

	s = append(s, 1)
	PrintSlice(s)
}