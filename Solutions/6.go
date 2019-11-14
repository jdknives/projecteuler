package main

import "fmt"

func main () {
	i := 100

	x := sumsquare(i)
	y := squaresum(i)

	z := y - x
	fmt.Println(z)
}

func squaresum(num int) int {
	sumformula := ((num * num) + num)/2

	return sumformula * sumformula
}

func sumsquare(num int) int {
	squareformula := 0
	for i:=1; i <  num+1; i++{
		squareformula = squareformula + (i*i)
	}
	return squareformula
}