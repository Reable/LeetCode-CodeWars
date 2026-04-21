package main

import (
	"fmt"
	"strings"
)

func test6() {
	strings := []string{
		"PAYPALISHIRING",
		// "PAYPALISHIRING",
		"AB",
	}

	numRows := []int{
		3,
		// 4,
		1,
	}

	results := []string{
		"PAHNAPLSIIGYIR",
		// "PINALSIGYAHRPI",
		"AB",
	}

	for i := 0; i < len(strings); i++ {
		result := convert1(strings[i], numRows[i])
		fmt.Println("Correct answer: ", result == results[i], result)
	}
}

func convert1(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	rows := make([][]rune, numRows)

	row := 0
	direction := 1

	for _, char := range s {
		rows[row] = append(rows[row], char)

		if row == 0 {
			direction = 1
		} else if row == numRows-1 {
			direction = -1
		}

		row += direction
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
