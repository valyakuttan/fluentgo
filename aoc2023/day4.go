// cat sample.input | go run day4.go

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readLines() (lines []string) {
	// open file
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return
}

type Card struct {
	WNumbers map[int]bool
	CNumbers []int
}

func (c Card) Points() int {
	count := 0
	for _, n := range c.CNumbers {
		if c.WNumbers[n] {
			count++
		}
	}

	if count > 0 {
		return 1 << (count - 1)
	}
	return 0
}
func stringToNums(s string) (nums []int) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, n)
	}

	return
}

func main() {
	lines := readLines()

	var cards []Card
	for _, line := range lines {
		xs := strings.Split(line, ":")
		xs = strings.Split(xs[1], "|")

		wnumbers := make(map[int]bool)
		for _, n := range stringToNums(xs[0]) {
			wnumbers[n] = true
		}

		card := Card{wnumbers, stringToNums(xs[1])}
		cards = append(cards, card)
	}

	pts := make([]int, len(cards))
	for i, c := range cards {
		pts[i] = c.Points()
	}

	fmt.Println(pts)
}
