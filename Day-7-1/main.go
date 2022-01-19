package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// unnecessary but it's a habit lul
	log.SetFlags(log.Lshortfile)

	// reading file
	k, err := os.Open("examplecase.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	// basic example for testing
	log.Println(solution(k))

	// reading file
	f, err := os.Open("testcase.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// solution
	log.Println(solution(f))
}

// solution function for solution implement
func solution(s io.Reader) int {
	// conver reader to bufio type
	// not necessary I added for process big inputs easily
	testcase := bufio.NewReader(s)

	var crabs []int
	for {
		// read until \n
		line, _, err := testcase.ReadLine()
		if err != nil {
			// if file ends we shut down for loop
			log.Println(err)
			break
		}
		splitted := strings.Split(string(line), ",")
		crabs = make([]int, len(splitted))
		for i, val := range splitted {
			tmp, _ := strconv.Atoi(val)
			crabs[i] = tmp
		}
	}

	possible := possibles(findMin(crabs...), findMax(crabs...))

	var res []int
	for _, num := range possible {
		res = append(res, calc(crabs, num))
	}

	var min = res[0]
	var min_index int
	for index, num := range res {
		if min > num {
			min = num
			min_index = index
		}
	}
	fmt.Println(possible[min_index], min)
	// finaly we have answer !!!
	return min
}

func calc(crabs []int, point int) int {
	fuel := 0
	for _, crab := range crabs {
		fuel += abs(crab - point)
	}
	return fuel
}

func findMax(nums ...int) int {
	var res int
	for index, val := range nums {
		if index == 0 || val > res {
			res = val
		}
	}
	return res
}
func findMin(nums ...int) int {
	var res int
	for index, val := range nums {
		if index == 0 || val < res {
			res = val
		}
	}
	return res
}
func possibles(min int, max int) []int {
	nums := make([]int, 0, max-min+1)
	for v := min; v <= max; v++ {
		nums = append(nums, v)
	}
	return nums
}
func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
