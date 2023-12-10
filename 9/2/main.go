package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inp := read_input("input.txt")
	inp1 := proc1(inp)

	inp2 := [][][]int{}
	for _, e := range inp1 {
		inp2 = append(inp2, proc2(e))
	}
	// fmt.Println(inp2[0])
	// fmt.Println(inp2[1])
	// fmt.Println(inp2[3])

	for i := range inp2 {
		proc3(inp2[i])
		// fmt.Println("=============")
		// fmt.Println(inp2[i])
		// fmt.Println("=============")
	}

	sum := 0
	for _, e := range inp2 {
		sum += proc4(e)
		// break
	}

	// proc4(inp2[2])
	fmt.Println(sum)
}

func read_input(filename string) []string {
	d, err := os.ReadFile(filename)
	if err != nil {
		os.Exit(1)
	}

	return strings.Split(string(d), "\n")
}

func str_conv(inp string) []int {
	d := strings.Split(inp, " ")
	res := []int{}
	for _, e := range d {
		d1, err := strconv.Atoi(e)
		if err != nil {
			os.Exit(1)
		}
		res = append(res, d1)
	}
	return res
}

func proc1(inp []string) [][]int {
	res := [][]int{}
	for _, e := range inp {
		res = append(res, str_conv(e))
	}

	return res
}

func proc2(root []int) [][]int { //build pyramid minimally
	pyr := [][]int{}
	pyr = append(pyr, root)
	for i := 0; i < len(root)-1; i++ {
		pyr = append(pyr, []int{})
	}
	max_idx := 0

	for i := 0; i < len(root)-1; i++ { //iterate over root

		for j := 1; j < len(pyr); j++ { //pyramid level
			if len(pyr[j-1]) > 1 { // if there are more than 2 elements in the previous level
				idx := len(pyr[j])
				pyr[j] = append(pyr[j], pyr[j-1][idx+1]-pyr[j-1][idx])
			}
		}
		if len(pyr[i]) >= 2 && pyr[i][1] == pyr[i][0] && pyr[i][0] == 0 {
			max_idx = i
		}
	}

	pyr = pyr[:max_idx]

	return pyr
}

func proc3(inp [][]int) { //fill pyramid
	max_idx := len(inp) - 1

	for len(inp[1]) != len(inp[0])-1 {
		inp[max_idx] = append(inp[max_idx], inp[max_idx][0])
		// fmt.Println(inp[max_idx])

		for i := max_idx - 1; i >= 1; i-- {
			idx_top := len(inp[i+1]) - 1
			idx := len(inp[i]) - 1
			d := inp[i+1][idx_top] + inp[i][idx]
			inp[i] = append(inp[i], d)
		}
	}
}

func proc4(inp [][]int) int {
	sum := 0

	// fmt.Println(inp)
	for i := len(inp) - 1; i > -1; i-- {
		sum = inp[i][0] - sum
	}
	// fmt.Println(sum)

	return sum
}
