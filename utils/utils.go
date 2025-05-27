package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
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

func WinratePercentage(wins, losses int)int {
	if wins == 0 && losses == 0 {
		return 0
	}
	return int(float64(wins)/float64(losses/wins)*100)
}