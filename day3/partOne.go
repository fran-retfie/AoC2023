package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	sum := 0

	scanner := bufio.NewScanner(os.Stdin)
	partNumber_re := regexp.MustCompile("\\d*")

	for scanner.Scan() {
		text := scanner.Text()
		for _, slice := range partNumber_re.FindAllStringIndex(text, -1) {
			isPartNumber := false
			if slice[0] != slice[1] {
				fmt.Println(slice)
			}
			if slice[0] > 0 {
				if text[slice[0]-1] != '.' {
					isPartNumber = true
				}
			}
			if slice[1] < len(text) {
				if text[slice[1]] != '.' {
					isPartNumber = true
				}
			}
		}
	}
	fmt.Println(sum)
}
