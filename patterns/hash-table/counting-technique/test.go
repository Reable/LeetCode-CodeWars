package main

import "fmt"

func main() {
	str := "ABAABCC"
	counst := make(map[byte]int)

	for i := range str {
		counst[str[i]]++
	}

	odd_count := 0
	for i := range counst {
		if counst[i]%2 != 0 {
			odd_count++
		}
	}

	fmt.Println(odd_count <= 1)
}
