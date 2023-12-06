package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var rg *regexp.Regexp = regexp.MustCompile(`[a-z|A-Z]+:[\s]+`)
var rg1 *regexp.Regexp = regexp.MustCompile(`[\s]+`)

func main() {
	inp := read_input("input.txt")

	inp1 := proc1(inp)

	inp2 := [][]int{}

	for _, e := range inp1 {
		inp2 = append(inp2, str_2_int_arr(e, " "))
	}

	fmt.Println(inp2)
	d1, d2 := proc3(0, inp2[0][0], inp2[1][0])
	for d2-d1 > 1 {
		d1, d2 = proc3(d1, inp2[0][0], inp2[1][0])
	}
	d3 := inp2[0][0] - d2
	d4 := d3 - d2 + 1
	fmt.Println(d4)
	// sum := 1
	// for i := range inp2[0] {
	// 	sum *= proc2(inp2[0][i], inp2[1][i])
	// }

	// fmt.Println(sum)
}

func read_input(filename string) []string {
	d, err := os.ReadFile(filename)
	if err != nil {
		os.Exit(1)
	}

	return strings.Split(string(d), "\n")
}

func proc1(inp []string) []string {
	for i, e := range inp {
		inp[i] = rg.ReplaceAllString(e, "")
		inp[i] = rg1.ReplaceAllString(inp[i], "")
	}

	return inp
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

func proc2(inp int, inp1 int) int {
	floor := 0
	ceil := 0

	for i := 1; i < inp; i++ {
		if i*(inp-i) > inp1 {
			floor = i
			break
		}
	}

	for i := inp - 1; i >= floor; i-- {
		if i*(inp-i) > inp1 {
			ceil = i
			break
		}
	}

	fmt.Println(ceil, floor)
	return ceil - floor + 1
}

func proc3(floor int, ceil int, target int) (int, int) {
	for i := 0; (floor + int(math.Pow(2, float64(i)))) <= ceil; i++ {
		d := (floor + int(math.Pow(2, float64(i))))
		if (ceil-d)*d >= target {
			return (floor + int(math.Pow(2, float64(i-1)))), d
		}
	}
	os.Exit(1)
	return -1, -1
}
