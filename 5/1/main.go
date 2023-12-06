package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var rg *regexp.Regexp = regexp.MustCompile(`^.+:[\n|\s]*`)

func main() {
	inp := open_file("input.txt")
	inp1 := inp[0]
	inp2 := inp[1:]

	inp1 = rg.ReplaceAllString(inp1, "")
	for i, e := range inp2 {
		inp2[i] = rg.ReplaceAllString(e, "")
	}

	inp3 := proc1(inp2)

	inp4 := proc2(inp3, " ")
	inp5 := str_2_int_arr(inp1, " ")

	// for _, e := range inp4 {
	// 	fmt.Println(e)
	// }
	// fmt.Println(inp5)

	// for _, e := range inp5 {
	// 	fmt.Println(e, proc5(inp4[0], e))
	// }

	res := math.MaxInt
	for _, e := range inp5 {
		// fmt.Println(proc6(e, inp4))
		res = int(
			math.Min(
				float64(res),
				float64(
					proc6(e, inp4),
				),
			),
		)
	}
	fmt.Println(res)

}

func open_file(filename string) []string {
	d, err := os.ReadFile(filename)
	if err != nil {
		os.Exit(1)
	}

	return strings.Split(string(d), "\n\n")
}

func proc1(inp []string) [][]string {
	res := [][]string{}

	for _, e := range inp {
		d1 := strings.Split(e, "\n")
		res = append(res, d1)
	}

	return res
}

func str_2_int_arr(inp string, inp1 string) []int {
	d := strings.Split(inp, inp1)
	res := []int{}

	for _, e := range d {
		d1, err := strconv.Atoi(e)
		if err != nil {
			fmt.Println(fmt.Errorf(e))
			os.Exit(1)
		}
		res = append(res, d1)
	}

	return res
}

func proc2(inp [][]string, inp1 string) [][][]int {
	res := [][][]int{}

	for _, e := range inp {
		// d1 := str_2_int_arr(e[0], inp1)
		// d2 := str_2_int_arr(e[1], inp1)
		d3 := [][]int{}
		for _, e1 := range e {
			d1 := str_2_int_arr(e1, inp1)
			d3 = append(d3, d1)
		}
		res = append(res, d3)
	}

	return res
}

func proc3(inp int, inp1 int, inp2 int) (bool, int, int) { //check if within range
	floor := inp
	ceil := inp + inp1 - 1

	// fmt.Println("FLOOR", floor, "\n", "CEIL", ceil)

	return (floor <= inp2) && (inp2 <= ceil), floor, ceil
}

func proc4(inp []int, inp1 int) (bool, int) {
	d1, floor, _ := proc3(inp[1], inp[2], inp1)
	if d1 {
		return true, inp1 - floor + inp[0]
	}
	return false, inp1
}

func proc5(inp [][]int, inp1 int) int {
	d1, d2 := proc4(inp[0], inp1)
	if d1 {
		return d2
	}
	for i := 1; i < len(inp); i++ {
		d1, d3 := proc4(inp[i], inp1)
		if d1 {
			return d3
		}
	}
	// for _, e := range inp {
	// 	d1, d3 := proc4(e, inp1)
	// 	if d1 {
	// 		return d3
	// 	}
	// }
	return d2
}

func proc6(seed int, m [][][]int) int {
	res := seed

	for _, e := range m {
		res = proc5(e, res)
	}

	return res
}
