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
	inp2, inp3 := proc2(inp[1])

	fmt.Println(inp1)

	fmt.Println(len(inp2))

	fmt.Println(inp3)

	inp4 := []int{}
	for _, e := range inp3 {
		inp4 = append(inp4, proc3(inp1, inp2, e))
	}

	fmt.Println(inp4, "+++++++")
	// fmt.Println(proc5(inp4[0:2]))

	fmt.Println(proc6(inp4))
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

func proc2(inp string) (map[string][]string, []string) {
	res := map[string][]string{}
	res1 := []string{}

	d := rg.ReplaceAllString(inp, "")
	d1 := rg1.ReplaceAllString(d, " ")
	d2 := strings.Split(d1, "\n")

	// fmt.Println(d1)
	for _, e := range d2 {
		d3 := strings.Split(e, " ")
		res[d3[0]] = []string{d3[1], d3[2]}
		if d3[0][2] == 'A' {
			res1 = append(res1, d3[0])
		}
		// fmt.Println(d3[0], d3[1], d3[2])
	}

	return res, res1
}

func proc3(inp []int, inp1 map[string][]string, inp2 string) int {
	count := 0
	current := inp2
	for current[2] != 'Z' {
		// fmt.Println(count)
		current = inp1[current][inp[count%len(inp)]]
		count += 1
	}

	return count
}

func proc4(inp []int) int {
	for i := 0; i < len(inp)-1; i += 1 {
		inp[i+1] = proc5(inp[i : i+2])
		// fmt.Println(inp[i+1])
	}
	return inp[len(inp)-1]
}

func proc5(inp []int) int {
	d := 0
	// fmt.Println("=========")
	for inp[0]%inp[1] != 0 {
		d = inp[0] % inp[1]
		inp[0] = inp[1]
		inp[1] = d

		// fmt.Println(inp[0], inp[1], inp[0]/inp[1])
	}
	return inp[1]
}

func proc6(inp []int) int {
	fmt.Println(inp[0])
	for i := 0; i < len(inp)-1; i += 1 {
		fmt.Println(i, inp[i], inp[i+1], proc5([]int{inp[i], inp[i+1]}))
		inp[i+1] = inp[i] * inp[i+1] / proc5([]int{inp[i], inp[i+1]})
		// fmt.Println("+++++")
		// fmt.Println(proc5(inp[i:i+2]))
	}
	return inp[len(inp)-1]
}
