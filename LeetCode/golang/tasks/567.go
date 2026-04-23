package tasks

func CheckInclusion567(s1 string, s2 string) bool {

	if len(s2) < len(s1) {
		return false
	}

	if len(s1) == 0 {
		return true
	}

	var s1Count [26]int
	var windowCount [26]int

	for i := 0; i < len(s1); i++ {
		s1Count[s1[i]-'a'] += 1
		windowCount[s2[i]-'a'] += 1
	}

	if s1Count == windowCount {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		windowCount[s2[i]-'a']++
		windowCount[s2[i-len(s1)]-'a']--
		if s1Count == windowCount {
			return true
		}
	}

	return false
}

// func CheckInclusion(s1 string, s2 string) bool {
// 	if len(s1) > len(s2) {
// 		return false
// 	}

// 	l := 0

// 	for i := 0; i < len(s2); i++ {
// 		if s1[0] == s2[i] {
// 			l = 0

// 			temp := i
// 			length := len(s1) - 1

// 			// steps to left
// 			if temp-length >= 0 {
// 				for l < len(s1) {
// 					if s1[l] != s2[temp] {
// 						break
// 					}
// 					temp--
// 					l++
// 				}
// 				if l >= len(s1) {
// 					return true
// 				}
// 			}

// 			temp = i
// 			l = 0

// 			// steps to right
// 			if i+len(s1) <= len(s2) {
// 				fmt.Println("Steps to right")
// 				for l < len(s1) {
// 					if s1[l] != s2[temp] {
// 						break
// 					}
// 					temp++
// 					l++
// 				}
// 				if l >= len(s1) {
// 					return true
// 				}
// 			}
// 		}
// 	}
// 	fmt.Println("\n exit function")
// 	return false
// }
