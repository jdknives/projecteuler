package main

import "fmt"

func main () {
	slice := (make([]int, 10001))
	for c := 1; c ; c++ {
			for d := 1; d < c; d++ {
				if (c%d == 0 && !(d != 1)) == true {
					slice = append(slice, c)
					if len(slice) == 10001 {
						break
					}
				}
				fmt.Println(slice[10001])
			}
		}
	}