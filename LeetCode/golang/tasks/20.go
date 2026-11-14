package tasks

func IsValid20(s string) bool {
	stack := []byte{}
	data := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for i := range s {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
			continue
		}

		if len(stack) == 0 {
			return false
		}

		last := data[s[i]]
		if stack[len(stack)-1] != last {
			return false
		}

		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}
