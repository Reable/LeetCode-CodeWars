package tasks

func RomanToInt13(s string) int {
	data := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	count := 0

	for i:=0; i<len(s); i++ {
		curr := data[s[i]]

		if i+1 < len(s) && curr < data[s[i+1]] {
			count-=curr
		} else {
			count+=curr
		}
	}

	return count
}
