package utils

import "fmt"

func DisplayOption() int{
	var option int
	fmt.Print("Select Option : ")
	fmt.Scan(&option)
	return option
}