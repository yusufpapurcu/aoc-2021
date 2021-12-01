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
	log.Println(solution(strings.NewReader("199\n200\n208\n210\n200\n207\n240\n269\n260\n263")))

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

	// declare needed variables
	var latest, increments int

	// this is a so bad solution for first time run
	// but I dont have time to find better solution
	var isFirst bool = true
	for {
		// read until \n
		line, _, err := testcase.ReadLine()
		if err != nil {

			// if file ends we shut down for loop
			log.Println(err)
			break
		}

		// convert it to integer
		num, err := strconv.Atoi(string(line))
		if err != nil {
			log.Println(err)
			continue
		}

		// shitty implementation
		if isFirst {
			isFirst = !isFirst
			latest = num
		}

		// check values
		if num > latest {
			increments++
		}

		// set for next run
		latest = num
	}

	// finaly we have answer !!!
	return increments
}
