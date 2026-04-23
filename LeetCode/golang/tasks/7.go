package tasks

import (
	"math"
	"strconv"
	"strings"
)

func Reverse(x int) int {
	var builder strings.Builder

	if x < 0 {
		builder.WriteByte('-')
		x = -x
	}

	strNumber := strconv.Itoa(x)

	pointer := len(strNumber) - 1

	for pointer >= 0 {
		builder.WriteByte(strNumber[pointer])
		pointer--
	}

	result, err := strconv.Atoi(builder.String())

	if err != nil {
		return 0
	}

	if result >= math.MinInt32 && result <= math.MaxInt32 {
		return result
	}

	return 0
}
