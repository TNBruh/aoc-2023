package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var rg *regexp.Regexp = regexp.MustCompile(`[,|(|)|=]+`)
var rg1 *regexp.Regexp = regexp.MustCompile(`[\x20]+`)

func main() {
	inp := read_input("input.txt")
	inp1 := proc1(inp[0])
	inp2, _ := proc2(inp[1])

	fmt.Println(inp1)

	fmt.Println(len(inp2))

	fmt.Println(proc3(inp1, inp2, "AAA"))

}

func read_input(filename string) []string {
	d, err := os.ReadFile(filename)
	if err != nil {
		os.Exit(1)
	}

	return strings.Split(string(d), "\n\n")
}

func proc1(inp string) []int {
	d := strings.Split(inp, "")
	res := []int{}

	for _, e := range d {
		if e == "L" {
			res = append(res, 0)
		} else {
			res = append(res, 1)
		}
	}

	return res
}

func proc2(inp string) (map[string][]string, string) {
	res := map[string][]string{}
	res1 := ""

	d := rg.ReplaceAllString(inp, "")
	d1 := rg1.ReplaceAllString(d, " ")
	d2 := strings.Split(d1, "\n")

	// fmt.Println(d1)
	for _, e := range d2 {
		d3 := strings.Split(e, " ")
		res[d3[0]] = []string{d3[1], d3[2]}
		if res1 == "" {
			res1 = d3[0]
		}
		// fmt.Println(d3[0], d3[1], d3[2])
	}

	return res, res1
}

func proc3(inp []int, inp1 map[string][]string, inp2 string) int {
	count := 0
	current := inp2
	for current != "ZZZ" {
		// fmt.Println(count)
		current = inp1[current][inp[count%len(inp)]]
		count += 1
	}

	return count
}
