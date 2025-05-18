package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ScanNumber() int {
	fmt.Print("Pilih Opsi: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	option, _ := strconv.Atoi(input)
	return option
}

func ScanString() string {
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}