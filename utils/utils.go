package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ScanNumber(propmt string) int {
	fmt.Print(propmt)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	option, _ := strconv.Atoi(input)
	return option
}

func ScanString(propmt string) string {
	fmt.Print(propmt)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}