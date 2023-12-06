package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var dict map[string]string = map[string]string{
	"zero": "0",
	"one": "1",
	"two": "2",
	"three": "3",
	"four": "4",
	"five": "5",
	"six": "6",
	"seven": "7",
	"eight": "8",
	"nine": "9",
}

type Pair struct {
	a int
	b int
}

func main() {
	inp, err := read_n_split_input()
	if err != nil { os.Exit(1) }

	replace_str_slice(inp)

	sum := 0
	for i := range inp {
		a, b := solution(inp[i])
		//pair_res = append(pair_res, Pair{a, b})
		_, err := strconv.Atoi(string(inp[i][a]))
		if err != nil { os.Exit(1) }
		_, err = strconv.Atoi(string(inp[i][b]))
		if err != nil { os.Exit(1) }
		
		e, _ := strconv.Atoi(string(inp[i][a]) + string(inp[i][b]))
		sum += e
	}

	fmt.Println(sum)
}

func read_n_split_input() ([]string, error) {

	inp, err := os.ReadFile("input.txt")
	if err != nil { return []string{}, err }

	str_inp := string(inp)
	return strings.Split(str_inp, "\n"), nil
}

func replace_str_slice(inp []string) {
	for i, e := range inp {

		//first 
		ind := len(e) - 1
		closest_key := ""
		for k := range dict {
			res_ind := strings.Index(e, k)
			if res_ind < 0 { continue }
			if res_ind < ind {
				ind = res_ind
				closest_key = k
			}
		}
		if closest_key != "" {
			inp[i] = strings.Replace(inp[i], closest_key, string(closest_key[0]) + dict[closest_key] + string(closest_key[len(closest_key) - 1]), 1)
		}

		//last
		inversed_str := StringReverse(inp[i])
		ind = len(e) - 1
		closest_key = ""

		for k := range dict {
			inversed_key := StringReverse(k)
			res_ind := strings.Index(inversed_str, inversed_key)

			if res_ind < 0 { continue }
			if res_ind < ind {
				ind = res_ind
				closest_key = k
			}
		}
		if closest_key != "" {
			inp[i] = StringReverse(strings.Replace(inversed_str, StringReverse(closest_key), dict[closest_key], 1))
		}
	}
}


func StringReverse(InputString string) (ResultString string) {
	// iterating and prepending
	for _, c := range InputString {
		ResultString = string(c) + ResultString
	}
	return
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