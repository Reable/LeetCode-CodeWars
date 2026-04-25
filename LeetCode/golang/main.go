package main

import (
	"fmt"
	"golang/tasks"
)

func main() {
	fmt.Println(tasks.RemoveElement27([]int{3, 2, 2, 3}, 3))
	fmt.Println(tasks.RemoveElement27([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
