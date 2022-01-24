package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
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
	// convert reader to bufio type
	// not necessary I added for process big inputs easily
	scanner := bufio.NewScanner(s)

	heightmap := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		newLevel := []int{}
		for _, n := range line {
			num, err := strconv.Atoi(string(n))
			if err != nil {
				log.Fatalf("Failed to convert %v to an integer: %v", n, err)
			}
			newLevel = append(newLevel, num)
		}
		heightmap = append(heightmap, newLevel)
	}

	riskLevel := 0
	for Y := 0; Y < len(heightmap); Y++ {
		for X := 0; X < len(heightmap[Y]); X++ {
			isRisk := true
			curLoc := heightmap[Y][X]
			if Y != 0 {
				if curLoc >= heightmap[Y-1][X] {
					isRisk = false
				}
			}
			if Y < len(heightmap)-1 {
				if curLoc >= heightmap[Y+1][X] {
					isRisk = false
				}
			}
			if X != 0 {
				if curLoc >= heightmap[Y][X-1] {
					isRisk = false
				}
			}
			if X < len(heightmap[Y])-1 {
				if curLoc >= heightmap[Y][X+1] {
					isRisk = false
				}
			}
			if isRisk {
				riskLevel += curLoc + 1
			}
		}
	}
	fmt.Printf("Risk level is %d\n", riskLevel)

	// finaly we have answer !!!
	return riskLevel
}
