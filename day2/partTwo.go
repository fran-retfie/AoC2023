package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	sum := 0

	scanner := bufio.NewScanner(os.Stdin)
	gameRe := regexp.MustCompile("[;:][^;]*")
	redRe := regexp.MustCompile("\\d* red")
	greenRe := regexp.MustCompile("\\d* green")
	blueRe := regexp.MustCompile("\\d* blue")
	for scanner.Scan() {

		gameTexts := gameRe.FindAllString(scanner.Text(), -1)
		var maxred int = 0
		var maxgreen int = 0
		var maxblue int = 0

		for _, gameText := range gameTexts {

			red := redRe.FindAllString(gameText, -1)
			for _, r := range red {
				val, _ := strconv.Atoi(r[:len(r)-4])
				if val > maxred {
					maxred = val
				}
			}

			green := greenRe.FindAllString(gameText, -1)
			for _, r := range green {
				val, _ := strconv.Atoi(r[:len(r)-6])
				if val > maxgreen {
					maxgreen = val
				}
			}

			blue := blueRe.FindAllString(gameText, -1)
			for _, r := range blue {
				val, _ := strconv.Atoi(r[:len(r)-5])
				if val > maxblue {
					maxblue = val
				}
			}
		}
		sum += maxred * maxgreen * maxblue
	}
	fmt.Println(sum)
}
