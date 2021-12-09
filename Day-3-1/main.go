package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// unnecessary but it's a habit lul
	log.SetFlags(log.Lshortfile)

	// basic example for testing
	log.Println(solution(strings.NewReader("00100\n11110\n10110\n10111\n10101\n01111\n00111\n11100\n10000\n11001\n00010\n01010")))

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
func solution(s io.Reader) int64 {
	// conver reader to bufio type
	// not necessary I added for process big inputs easily
	testcase := bufio.NewReader(s)
	var binary_map = make(map[int]int)
	var lines int
	for {
		// read until \n
		line, _, err := testcase.ReadLine()
		if err != nil {

			// if file ends we shut down for loop
			log.Println(err)
			break
		}

		for index, value := range string(line) {
			if value == '1' {
				binary_map[(index)] += 1
			}
		}
		lines++
	}

	var gamma_rate = make([]rune, len(binary_map))
	var epsilon_rate = make([]rune, len(binary_map))
	for key, value := range binary_map {
		if value < lines/2 {
			gamma_rate[key] = '0'
			epsilon_rate[key] = '1'
		} else {
			gamma_rate[key] = '1'
			epsilon_rate[key] = '0'
		}
	}

	gamma_rate_int, _ := strconv.ParseInt(string(gamma_rate), 2, 64)
	epsilon_rate_int, _ := strconv.ParseInt(string(epsilon_rate), 2, 64)

	// finaly we have answer !!!
	return (gamma_rate_int) * (epsilon_rate_int)
}
