package main

import "fmt"

var si []int
var ret []int

func main() {
	fmt.Println(add(1000))
}

func add(n int) int {
	list(n)
	var b int
	for _, i := range si {
		b += i
	}
	return b
}

func list(i int) {
	for j := 0; j < i; j++ {
		if j % 3 == 0 || j % 5 == 0 {
			si = append(si, j)
		}
	}
}
