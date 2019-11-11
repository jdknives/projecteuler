package main

import (
	"fmt"
	"sort"
	"strconv"
)

var palindromelist []int

func main() {

	for i := 100; i < 1000; i++ {
		for j := i; j < 1000; j++ {
			x := j * i
			if ispalindrome(x) {
				palindromelist = append(palindromelist, x)
			}
		}
	}

	sortslice(palindromelist)
}

func ispalindrome(num int) bool {
	str := strconv.Itoa(num)

	e := reverse(str)

	g := isequal(e, str)

	return g
}

func isequal(str1 string, str2 string) bool {
	if str1 == str2 {
		return true
	} else {
		return false
	}
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func sortslice(list []int) {
	sort.Ints(list)
	lastvalue := list[len(list)-1]
	fmt.Println(lastvalue)
}
