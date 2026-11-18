package tasks

func LongestCommonPrefix14(strs []string) string {

	prefix := make(map[string]int)

	for _, v := range strs {
		s := []byte(v)
		sum := []byte{}
		for i := range s {
			sum = append(sum, s[i])
			_, ok := prefix[string(sum)]
			if !ok {
				prefix[string(sum)] = 1
			} else {
				prefix[string(sum)] += 1
			}
		}
	}
	result := ""
	for k, v := range prefix {
		if v == len(strs) && len(k) > len(result) {
			result = k
		}
	}

	return result
}

func LongestCommonPrefix14_v2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		ch := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != ch {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}
