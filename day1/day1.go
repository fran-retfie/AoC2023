package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var digits = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}

	sum := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		minval := -1
		maxval := -1
		minindex := len(scanner.Text())
		maxindex := -1
		for key, value := range digits {
			indexl := strings.Index(scanner.Text(), key)
			if indexl != -1 && indexl < minindex {
				minindex = indexl
				minval = value
			}
			indexr := strings.LastIndex(scanner.Text(), key)
			if indexr != -1 && indexr > maxindex {
				maxindex = indexr
				maxval = value
			}
		}
		sum += minval*10 + maxval
	}
	fmt.Println(sum)
}
