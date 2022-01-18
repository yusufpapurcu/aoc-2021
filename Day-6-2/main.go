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

	var lanternFishes = make([]int, 9)
	for {
		// read until \n
		line, _, err := testcase.ReadLine()
		if err != nil {
			// if file ends we shut down for loop
			log.Println(err)
			break
		}
		fishes := strings.Split(string(line), ",")
		for _, val := range fishes {
			tmp, _ := strconv.Atoi(val)
			lanternFishes[tmp]++
		}
	}

	for i := 0; i < 256; i++ {
		newFish := lanternFishes[0]
		copy(lanternFishes, lanternFishes[1:])
		lanternFishes[6] += newFish
		lanternFishes[8] = newFish
	}
	var total int
	for _, count := range lanternFishes {
		total += count
	}

	// finaly we have answer !!!
	return total
}
