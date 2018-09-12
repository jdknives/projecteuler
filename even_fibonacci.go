package main

import "fmt"

var si2 []int
var d int = 4000000

func main () {
	fmt.Println(add())
}

func add () int {

	slice := make([]int,d)
	slice [0] = 0
	slice [1] = 1
	for i := 2; i < d; i++ {
		slice [i] = slice [i-1] + slice [i-2]
		if slice [i] >= d{
			break
		}
	}

	//fmt.Println(slice)

	var b int
	for _, k := range slice{
		if k % 2 == 0 {
			si2 = append(si2, k)
		}
	}

	//fmt.Println(si2)

	for _, j := range si2 {
		b += j
	}
	return b
}