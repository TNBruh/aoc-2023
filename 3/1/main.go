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

var num_rg *regexp.Regexp = regexp.MustCompile("[0-9]+")
var sym_rg *regexp.Regexp = regexp.MustCompile("[\x21-\x2d\x2F\x3a-\x40\x5b-\x60\x7b-\x7e]")

func main() {
	inp, err := read_n_split_input("input.txt");
	if err != nil { os.Exit(1) }

	inp1 := proc0(inp)

	inp2 := proc2(inp)
	inp3 := proc3(inp2, len(inp[len(inp)-1]))
	fmt.Println(inp1)
	fmt.Println(inp3)
	fmt.Println("===========")
	// for _, e := range inp1 {
	// 	fmt.Println(proc4(e, len(inp[len(inp)-1]), len(inp)))
	// }

	inp4 := proc5(inp1, inp3, len(inp[len(inp)-1]), len(inp))

	sum := 0
	for i, e := range inp4 {
		if e {
			line := inp[inp1[i][2]]
			num, err := strconv.Atoi(
				string(line[inp1[i][0]:inp1[i][1]]),
			)
			if err != nil {
				fmt.Println(e)
				os.Exit(1)
			}
			sum += num

		}
	}
	fmt.Println(sum)
}

func read_n_split_input(filename string) ([]string, error) {

	inp, err := os.ReadFile(filename)
	if err != nil {
		return []string{}, err
	}

	str_inp := string(inp)
	return strings.Split(str_inp, "\n"), nil
}

func proc0(inp []string) [][]int { //finds the start and end index of all numbers
	res := [][]int{}
	for i, e := range inp {
		a := append([][]int{}, num_rg.FindAllStringIndex(e, 9999)...)
		for j := range a {
			a[j] = append(a[j], i)
		}
		res = append(res, a...)
	}

	return res
}

func proc1(inp []int, sym_locs []int, line_count int, inp_count int) bool { //input is start, end index, then the line it is in; then the entire splitted input. finds if there's a symbol around

	for _, e := range proc4(inp, line_count, inp_count) {
		if slices.Contains(sym_locs, e) {
			return true
		}
	}

	return false
}

func proc5(inp [][]int, sym_locs []int, line_count int, inp_count int) []bool {
	res := []bool{}
	for _, e := range inp {
		res = append(res, proc1(e, sym_locs, line_count, inp_count))
	}

	return res
}

func proc2(inp []string) [][]int { //maps all symbols
	res := [][]int{}
	for _, e := range inp {
		entry := []int{}
		for _, k := range sym_rg.FindAllStringIndex(e, 9999) {
			entry = append(entry, k[0])
		}

		res = append(res, entry)
	}

	return res
}

func proc3(inp [][]int, l int) []int { //converts coord-like position into a linear position. 2nd input is length per line
	res := []int{}

	for i := range inp {
		for _, e:= range inp[i] {
			res = append(res, 
				e + i * l,
			)
		}
	}

	return res
}

func proc4(pos []int, line_count int, inp_count int) []int { //converts symbol-finding positions into linear positions
	res := []int{}
	low_x := pos[0] - 1
	hi_x := pos[1] + 1
	low_y := pos[2] - 1
	hi_y := pos[2] + 2
	
	low_x = int(math.Max(float64(low_x), 0))
	hi_x = int(math.Min(float64(line_count), float64(hi_x)))

	low_y = int(math.Max(float64(low_y), 0))
	hi_y = int(math.Min(float64(inp_count), float64(hi_y)))
	for i := low_y; i < hi_y; i++ {
		for j := low_x; j < hi_x; j++ {
			res = append(res, i * line_count + j)
		}
	}

	return res
}