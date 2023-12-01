package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var numWords = map[string]rune{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

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
		if a == 0 {
			if unicode.IsDigit(rune(line[ct])) {
				a = rune(line[ct])

			} else if num := numberWord(line[ct:]); num != 0 {
				a = num
			}
		}
		if b == 0 {
			if unicode.IsDigit(rune(line[length-ct-1])) {
				b = rune(line[length-ct-1])
			} else if num := numberWord(line[length-ct-1:]); num != 0 {
				b = num
			}
		}

		if a != 0 && b != 0 {
			break
		}
	}
	value, _ := strconv.ParseUint(string(a)+string(b), 10, 32)
	return value
}

func numberWord(line []byte) rune {
	var found rune

	s := string(line)

	for k, v := range numWords {
		if strings.HasPrefix(s, k) {
			return v
		}
	}

	return found
}
