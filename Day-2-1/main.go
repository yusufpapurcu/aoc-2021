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
	log.Println(solution(strings.NewReader("forward 5\ndown 5\nforward 8\nup 3\ndown 8\nforward 2")))

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
	var horizontal, vertical int
	for {
		// read until \n
		line, _, err := testcase.ReadLine()
		if err != nil {

			// if file ends we shut down for loop
			log.Println(err)
			break
		}

		// get fields
		fields := strings.Fields(string(line))

		// convert it to integer
		num, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Println(err)
			continue
		}

		// match keywords for filtering
		switch fields[0] {
		case "forward":
			horizontal += num
		case "down":
			vertical += num
		case "up":
			vertical -= num
		}
	}

	// finaly we have answer !!!
	return horizontal * vertical
}
