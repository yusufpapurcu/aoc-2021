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

	var lanternFishes = []fish{}
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
			lanternFishes = append(lanternFishes, fish{inc_duration: 6, timer: tmp})
		}
	}

	days := 0
	for {
		for index := range lanternFishes {
			lanternFishes[index].timer--
			if lanternFishes[index].timer < 0 {
				lanternFishes[index].timer = 6
				lanternFishes = append(lanternFishes, lanternFishes[index].newFish())
			}
		}
		days++
		if days == 80 {
			break
		}
	}

	// finaly we have answer !!!
	return len(lanternFishes)
}

type fish struct {
	inc_duration int
	timer        int
}

func (f fish) newFish() fish {
	return fish{
		inc_duration: 8,
		timer:        8,
	}
}
