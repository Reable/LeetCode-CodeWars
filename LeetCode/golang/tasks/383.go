package tasks

import "fmt"

func CanConstruct383(ransomNote string, magazine string) bool {
	count1 := make(map[rune]int)
	count2 := make(map[rune]int)

	for _, v := range ransomNote {
		count1[v]++
	}
	for _, v := range magazine {
		count2[v]++
	}

	for i, v := range count1 {
		fmt.Println(i, v)
		if value, ok := count2[i]; !ok || value < v {
			return false
		}
	}

	return true
}
