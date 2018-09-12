package main

import ("fmt"
"math")

const a float64 = 600851475143

func main() {
	slice := (make([]int, 0))
	slice = append(slice, 2)
	for c := 1 ; c < int(math.Sqrt(a)); c++ {
		if int(a) % c ==0 {
			for d:=1; d<c; d++ {
				if (c % d == 0 && !(d != 1)) == true {
					slice = append(slice, c)

		}
			}
		}
	}
	for e := 0; e <= len(slice); e ++{
		fmt.Println(int(a) / slice[e])
	}

}