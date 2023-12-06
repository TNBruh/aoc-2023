package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	inp, err := read_n_split_input("input.txt")
	if err != nil { os.Exit(1) }
	// fmt.Println(inp[0])

	proc0(inp)

	inp1 := proc1(inp)

	inp2 := proc2(inp1)

	inp3 := proc6(inp2) 

	sum := 0
	for _, i := range inp3 {
		sum += (i[0] * i[1] * i[2])
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

func proc0(inp []string) { // remove "game X: "
	reg := regexp.MustCompile("^Game [0-9]+: ")
	for i, e := range inp {
		// inp[i] = strings.Replace(e, "")
		inp[i] = reg.ReplaceAllString(e, "")
	}
}

func proc1(inp []string) [][]string{ // split ";"
	res := [][]string{}
	for _, e := range inp {
		res = append(res, strings.Split(e, ";"))
	}
	return res
}

func proc2(inp [][]string) [][][3]int { // convert to [R, G, B]
	stage := [][][3]int{}
	for _, e := range inp {
		stage_entry := [][3]int{}
		for _, v := range e {
			stage := strings.Split(v, ",")
			c := [3]int{0, 0, 0}

			//R
			r_ind := proc3(stage, "red")
			if r_ind != -1 {
				c[0] = proc4(stage[r_ind])
			}

			//G
			g_ind := proc3(stage, "green")
			if g_ind != -1 {
				c[1] = proc4(stage[g_ind])
			}

			//B
			b_ind := proc3(stage, "blue")
			if b_ind != -1 {
				c[2] = proc4(stage[b_ind])
			}
			
			stage_entry = append(stage_entry, c)
		}
		
		stage = append(stage, stage_entry)
	}

	return stage
}

func proc3(inp []string, color string) int { //find index with said color
	for i, e := range inp {
		if strings.Contains(e, color) { return i }
	}
	return -1
}

func proc4(inp string) int { //extract number from string
	r := regexp.MustCompile("[0-9]+")
	res, _ := strconv.Atoi(r.FindString(inp))
	return res
}

func proc5(inp [][3]int) [3]int { // finds max value of RGB in each stage
	maxR, maxG, maxB := 0, 0, 0
	for _, i := range inp {
		if i[0] > maxR { maxR = i[0] }
		if i[1] > maxG { maxG = i[1] }
		if i[2] > maxB { maxB = i[2] }
	}

	return [3]int{maxR, maxG, maxB}
}

func proc6(inp [][][3]int) [][3]int { // maps input into only having max RGB value in each stage
	res := [][3]int{}
	for _, i := range inp {
		res = append(res, proc5(i))
	}

	return res
}
