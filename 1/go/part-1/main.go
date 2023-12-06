package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inp, err := os.ReadFile("input.txt");
	if err != nil {
		os.Exit(1)
	}
	str_inp := string(inp)
	
	lines := strings.Split(str_inp, "\n")

	ans0 := []int{}

	for _, e := range lines {
		f_d, l_d := solution(e)
		entry, _ := strconv.Atoi(string(e[f_d]) + string(e[l_d]))
		ans0 = append(ans0, entry)
	}

	res := 0
	for _, e := range ans0 {
		res += e
	}

	fmt.Println(res)
}

func solution(inp string) (int, int) {
	first_digit := 0
	last_digit := len(inp) - 1

	for ; first_digit < len(inp); first_digit++ {
		_, err := strconv.Atoi(string(inp[first_digit]))
		if err == nil {
			break
		}
	}

	for ; last_digit >= first_digit; last_digit-- {
		_, err := strconv.Atoi(string(inp[last_digit]))
		if err == nil {
			break
		}
	}

	return first_digit, last_digit
}