package main

import "fmt"

var slice []int
var o = 2

func main() {

	for  {
		if isprime(o) {
			slice = append(slice, o)
			o += 1
		} else if len(slice) == 10002 {
			 	break
		} else {
			o += 1
			continue
		}
	}
	fmt.Println(slice)
	fmt.Println(slice[10001])
}

func isprime(num int) bool {

	for i := 2; i < num; i++ {
		if num%i == 0 && i != 1 && i != 2 {
			return false
		}
	}

	return true
}
