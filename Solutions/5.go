package main

import (
	"fmt"
)

var x int
var cn int = 20

func main() {
	i := 1
	for {
		if isdivisible(i) {
			fmt.Println(i)
			break
		} else {
			i++
		}
	}
}

func isdivisible(n int) bool {
	cnt := 0

	for i := 1; i < cn+1 ; i++ {
		if n%i == 0 {
			cnt += 1
		}
	}

	if cnt == cn {
		return true
	} else {
		return false
	}
}
