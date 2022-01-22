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

/*
FOUND: 0, 1, 3, 4, 5, 6, 7, 8, 9

1, 7, 4, 8 are identical
1-7 has "a" digit diff

6 - 9 - 0 has same amount of letter


2 - 3 - 5 has same amount of letter
3 - 2 has "c" 5 doesnt
3 - 5 has "f" 2 doesnt

3 contains all of 1's letters ***
3 and 9 has just 1 addition ***
3 - 1 remains all vertical lines, also 0 contains 2 of them ***
6 remained digit ***
comare 9 and 5, there is only 1 addition ***
*/
