package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var rg *regexp.Regexp = regexp.MustCompile(`Card[\s]+[0-9]+: `)
var rg1 *regexp.Regexp = regexp.MustCompile(`[\s]+`)

// var rg2 *regexp.Regexp = regexp.MustCompile(`^[\s]+|[\s]$`)

func main() {
	inp := open_file("input.txt")

	inp1 := proc0(inp)

	inp2 := proc1(inp1)

	inp3 := proc4(inp2)

	var sum float64 = 0
	for i, e := range inp3 {
		fmt.Println(i, proc5(e))
		d := float64(proc5(e))
		if d > 0 {
			sum += math.Pow(2, d-1)
		}
		// sum += math.Pow(2, float64(proc5(e)))
	}

	fmt.Println(sum)
}

func open_file(filename string) []string {
	d, err := os.ReadFile(filename)
	if err != nil {
		os.Exit(1)
	}

	return strings.Split(string(d), "\n")
}

func proc0(inp []string) []string {
	res := []string{}
	for _, e := range inp {
		d := rg.ReplaceAllString(e, "")
		res = append(res, d)
	}

	return res
}

func proc1(inp []string) [][]string {
	res := [][]string{}
	for _, e := range inp {
		d := rg1.ReplaceAllString(e, " ")
		d1 := strings.Split(d, " | ")
		res = append(res, d1)
	}

	return res
}

func proc2(inp []string) [][]int { //per line of input
	d := [][]string{
		strings.Split(inp[0], " "),
		strings.Split(inp[1], " "),
	}

	d1 := [][]int{}
	d1 = append(d1, proc3(d[0]), proc3(d[1]))

	return d1
}

func proc3(inp []string) []int { //converts to arr of int
	res := []int{}
	for _, e := range inp {
		d, err := strconv.Atoi(e)
		if err != nil {
			continue
		}
		res = append(res, d)
	}

	return res
}

func proc4(inp [][]string) [][][]int {
	res := [][][]int{}

	for _, e := range inp {
		res = append(res, proc2(e))
	}
	return res
}

func proc5(inp [][]int) int {
	res := 0
	for _, e := range inp[0] {
		if slices.Contains(inp[1], e) {
			res += 1
		}
	}

	return res
}
