package tasks

func CheckIfExist1346(arr []int) bool {
	temp := make(map[int]bool, len(arr))

	for _, v := range arr {
		if temp[v*2] || (v%2 == 0 && temp[v/2]) {
			return true
		}
		temp[v] = true
	}
	
	return false
}
