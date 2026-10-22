package tasks

func IsLongPressedName925(name string, typed string) bool {

	if (len(name) > len(typed)) || len(name) == 0 || len(typed) == 0 {
		return false
	}

	p1 := 0
	p2 := 0

	for p2 < len(typed){
		if p1 < len(name) && name[p1] == typed[p2]{
			p1++
			p2++
		} else if p2 > 0 && typed[p2] == typed[p2-1]{
			p2++
		} else {
			return false
		}
	}

	return p1 == len(name)

}
