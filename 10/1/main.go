package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

var pipes map[string][][2]int = map[string][][2]int{
	"|": [][2]int{
		[2]int{0, -1},
		[2]int{0, 1},
	},
	"-": [][2]int{
		[2]int{1, 0},
		[2]int{-1, 0},
	},
	"L": [][2]int{
		[2]int{0, -1},
		[2]int{1, 0},
	},
	"J": [][2]int{
		[2]int{0, -1},
		[2]int{-1, 0},
	},
	"7": [][2]int{
		[2]int{0, 1},
		[2]int{-1, 0},
	},
	"F": [][2]int{
		[2]int{0, 1},
		[2]int{1, 0},
	},
	".": [][2]int{},
	"S": [][2]int{
		[2]int{0, 1},
		[2]int{1, 0},
		[2]int{-1, 0},
		[2]int{0, -1},
	},
}

func main() {
	inp := read_input("input.txt")
	inp1 := proc1(inp)

	fmt.Println(inp1)

	proc2(inp1, inp)

	for _, e := range inp {
		fmt.Println(e)
	}
}

func read_input(filename string) []string {
	d, err := os.ReadFile(filename)
	if err != nil {
		os.Exit(1)
	}

	return strings.Split(string(d), "\n")
}

func proc1(inp []string) [2]int { // find S coord
	x := -1
	for i := 0; i < len(inp); i++ {
		x = strings.Index(inp[i], "S")
		if x >= 0 {
			return [2]int{x, i}
		}
	}

	os.Exit(1)
	return [2]int{-1, -1}
}

func proc2(start [2]int, m []string) {
	d := map[[2]int][][2]int{
		start: pipes["S"],
	}
	fmt.Println(d)

	count := 0

	for len(d) > 0 {
		// c := strconv.Itoa(count)
		// fmt.Println("==============")
		// extract keys
		keys := [][2]int{}
		for k := range d {
			keys = append(keys, k)
		}

		// iterate over
		for _, k := range keys {
			// collapse the state to "."
			y := k[1]
			x := k[0]
			// temp_l := string(m[y][x])
			temp := []byte(m[y])
			temp[x] = '*'
			m[y] = string(temp)

			// list directions
			temp1 := [][2]int{}
			for _, e := range d[k] {
				// our coords is k
				// e is the direction we're heading
				d1 := [2]int{
					k[0] + e[0],
					k[1] + e[1],
				}
				temp1 = append(temp1, d1)
				// fmt.Println("FINDING NEXT PATH TO CONSIDER")
				// fmt.Println(temp1, k, e)
			}

			// check if we can go
			// listing all the holes
			approved := [][2]int{}
			entrance := [][2]int{}
			for _, e := range temp1 {
				// e is the position we're about to enter
				// fmt.Println(e)
				if e[0] < 0 || e[1] < 0 || e[0] >= len(m) || e[1] >= len(m[0]) {
					continue
				}
				d1 := pipes[string(m[e[1]][e[0]])] // the holes
				// fmt.Println(string(m[e[1]][e[0]]), d1)

				for _, e1 := range d1 {
					d2 := [2]int{
						e[0] + e1[0],
						e[1] + e1[1],
					}
					// fmt.Println(d2, e, e1, k, temp_l, d2 == k)

					if d2 == k {
						approved = append(approved, e)
						entrance = append(entrance, e1)
						break
					}
				}
			}

			// append new entries
			// fmt.Println("+++")
			for i, j := range approved {
				letter := string(m[j[1]][j[0]])
				// fmt.Println(j, letter)
				val := slices.Clone(pipes[letter])
				rm_idx := proc3(val, entrance[i])
				if rm_idx == -1 {
					fmt.Println("SKIPPED: ", i, j)
					continue
				}
				val = slices.Delete(val, rm_idx, rm_idx+1)
				// fmt.Println(letter, val)

				d[j] = val
			}
		}

		// delete keys
		for _, k := range keys {
			delete(d, k)
		}

		// for _, e := range m {
		// 	fmt.Println(e)
		// }
		// var w1 string
		// n, _ := fmt.Scanln(&w1)
		// println(n)

		count++
	}
	fmt.Println(count - 1)
}

func proc3(inp [][2]int, el [2]int) int {
	for i, e := range inp {
		if e[0] == el[0] && e[1] == el[1] {
			return i
		}
	}

	return -1
}
