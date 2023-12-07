package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	cards    map[string]int
	og_cards string
	bid      int
}

var cards map[string]int = map[string]int{
	"A": 12,
	"K": 11,
	"Q": 10,
	"T": 8,
	"9": 7,
	"8": 6,
	"7": 5,
	"6": 4,
	"5": 3,
	"4": 2,
	"3": 1,
	"2": 0,
	"J": -1,
}

func main() {
	inp, inp1 := read_input("input.txt")

	inp2 := []map[string]int{}

	for _, e := range inp {
		inp2 = append(inp2, proc1(e))
	}

	// for _, e := range inp2 {
	// 	fmt.Println(e, proc2(e), proc3(e))
	// }

	inp3 := [][]Hand{[]Hand{}, []Hand{}, []Hand{}, []Hand{}, []Hand{}, []Hand{}, []Hand{}}
	for i, e := range inp2 {
		ind := proc2(e)
		h := Hand{
			cards:    e,
			og_cards: inp[i],
			bid:      conv_str(inp1[i]),
		}
		inp3[ind] = append(inp3[ind], h)
	}

	// fmt.Println(inp3)
	// fmt.Println("==============")

	for i := range inp3 {
		slices.SortFunc(
			inp3[i],
			func(a, b Hand) int {
				for i := range a.og_cards {
					if cards[string(a.og_cards[i])] < cards[string(b.og_cards[i])] {
						return -1
					} else if cards[string(a.og_cards[i])] > cards[string(b.og_cards[i])] {
						return 1
					}
				}
				return 0
			},
		)
	}

	inp4 := []Hand{}
	for _, e := range inp3 {
		inp4 = append(inp4, e...)
	}

	sum := 0
	for i, e := range inp4 {
		sum += e.bid * (i + 1)
	}
	fmt.Println(sum)

	// fmt.Println(inp3)
	// fmt.Println("==============")
	// for i, e := range inp2 {
	// 	h := Hand{
	// 		cards:    e,
	// 		og_cards: inp[i],
	// 		bid:      conv_str(inp1[i]),
	// 	}
	// 	inp3 = append(inp3, h)
	// }

	// for _, e := range inp3 {
	// 	fmt.Println(e)
	// }
}

func read_input(filename string) ([]string, []string) {
	d, err := os.ReadFile(filename)
	if err != nil {
		os.Exit(1)
	}

	d1 := strings.Split(string(d), "\n")

	res := []string{}
	res1 := []string{}

	for _, e := range d1 {
		d2 := strings.Split(e, " ")
		res = append(res, string(d2[0]))
		res1 = append(res1, string(d2[1]))
	}

	return res, res1
}

func conv_str(inp string) int {
	d, err := strconv.Atoi(inp)
	if err != nil {
		os.Exit(1)
	}

	return d
}

func proc1(inp string) map[string]int {
	res := map[string]int{}

	for k, _ := range cards {
		d := strings.Count(inp, k)
		if d != 0 {
			res[k] = d
		}
	}

	return res
}

func proc2(inp map[string]int) int { //find type
	if len(inp) == 1 {
		return 6
	} else if len(inp) == 2 { // 4-of-a-kind, full house,
		for _, v := range inp {
			if v == 4 {
				return 5
			} else if v == 3 {
				return 4
			}
		}
	} else if len(inp) == 3 { // 3-of-a-kind, 2 pair
		pair_count := 0
		for _, v := range inp {
			if v == 3 {
				return 3
			} else if v == 2 {
				pair_count += 1
			}
		}
		if pair_count == 2 {
			return 2
		}
	} else if len(inp) == 4 { // 1 pair
		return 1
	}

	return 0
}

func proc3(inp map[string]int) int { // count total unit
	sum := 0
	for k, v := range inp {
		sum += cards[k] * v
	}

	return sum
}
