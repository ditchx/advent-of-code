package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	inputFile := os.Args[1]

	fp, _ := os.Open(inputFile)

	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanLines)

	var total uint64
	for scanner.Scan() {
		total += calibrationValue([]byte(scanner.Text()))
	}

	fmt.Println(total)

}

func calibrationValue(line []byte) uint64 {
	var a, b rune
	length := len(line)

	for ct := 0; ct < length; ct++ {
		if unicode.IsDigit(rune(line[ct])) && a == 0 {
			a = rune(line[ct])
		}
		if unicode.IsDigit(rune(line[length-ct-1])) && b == 0 {
			b = rune(line[length-ct-1])
		}

		if a != 0 && b != 0 {
			break
		}
	}
	value, _ := strconv.ParseUint(string(a)+string(b), 10, 32)
	fmt.Println(value)
	return value
}
