package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	// unnecessary but it's a habit lul
	log.SetFlags(log.Lshortfile)

	// basic example for testing
	// log.Println(solution(strings.NewReader("199\n200\n208\n210\n200\n207\n240\n269\n260\n263")))

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

	// let's do some brute force
	testcase_arr := []int{}
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

		// create testcase_arr
		testcase_arr = append(testcase_arr, num)
	}

	var increments int
	for i := range testcase_arr {
		if i+3 > len(testcase_arr)-1 {
			break
		}

		if calculateChunk(testcase_arr[i:i+3]) < calculateChunk(testcase_arr[i+1:i+4]) {
			increments++
		}
	}

	// finaly we have answer !!!
	return increments
}

// calculateChunk function for cool looking if statement :D
func calculateChunk(arr []int) int {
	return arr[0] + arr[1] + arr[2]
}
