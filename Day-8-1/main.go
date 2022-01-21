package main

import (
	"bufio"
	"io"
	"log"
	"os"
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
	var total int
	for {
		// read until \n
		line, _, err := testcase.ReadLine()
		if err != nil {
			// if file ends we shut down for loop
			log.Println(err)
			break
		}
		output := strings.Split(string(line), "|")[1]
		var count int
		for _, digits := range strings.Fields(output) {
			switch len(digits) {
			case 2, 3, 4, 7:
				count++
			}
		}
		total += count
	}

	// finaly we have answer !!!
	return total
}
