package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode/utf8"
)

func main() {

	loc := flag.String("filelocation", "", "Enter the location of the file here.")
	slice := slicefromfile(loc)

	var xyz int

	for i := 0; i < 989; i++ {
		yxz := multiplyval(slice[i : i+13])
		if yxz > xyz {
			xyz = yxz
		} else {
			continue
		}
	}
	fmt.Println(xyz)
	fmt.Println(slice)
}

func multiplyval(nums []int) int {
	bb := 1
	fmt.Println(nums)
	for i := 0; i < 13; i++ {
		bb = bb * nums[i]
	}
	fmt.Println(bb)
	return bb
}

func slicefromfile(loc *string) []int {

	file, err := os.Open(*loc)
	if err != nil {
		log.Fatal("Could not open file.")
	}

	defer file.Close()

	fileinfo, _ := file.Stat()

	filesize := fileinfo.Size()

	buffer := make([]byte, filesize)

	file.Read(buffer)

	str := string(buffer)

	slice := explode(str, -1)

	var slice1 []int

	for i := 0; i < 1000; i++ {
		z, _ := strconv.Atoi(slice[i])
		slice1 = append(slice1, z)
	}

	return slice1
}

func explode(s string, n int) []string {
	l := utf8.RuneCountInString(s)
	if n < 0 || n > l {
		n = l
	}
	a := make([]string, n)
	for i := 0; i < n-1; i++ {
		ch, size := utf8.DecodeRuneInString(s)
		a[i] = s[:size]
		s = s[size:]
		if ch == utf8.RuneError {
			a[i] = string(utf8.RuneError)
		}
	}
	if n > 0 {
		a[n-1] = s
	}
	return a
}
