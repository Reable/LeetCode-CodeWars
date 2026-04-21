package tasks

import (
	"strings"
)

func Convert1(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	rows := make([][]rune, numRows)

	row := 0
	direction := 1

	for _, char := range s {
		rows[row] = append(rows[row], char)

		row += direction

		if row == 0 {
			direction = 1
		} else if row == numRows-1 {
			direction = -1
		}
	}

	var result strings.Builder
	for _, part := range rows {
		result.WriteString(string(part))
	}

	return result.String()
}

func convert2(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	var result strings.Builder

	temp := make(map[int][]rune, numRows)

	row := 1
	direction := true // true->down, false->up

	for _, v := range s {
		if row == numRows {
			direction = false
		} else if row == 1 {
			direction = true
		}

		if direction {
			temp[row] = append(temp[row], v)

			row++
		} else {
			temp[row] = append(temp[row], v)

			row--
		}
	}

	for _, part := range temp {
		result.WriteString(string(part))
	}
	return result.String()
}
