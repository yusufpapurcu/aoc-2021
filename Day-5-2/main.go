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

	var points = []point{}
	var mapsize int
	for {
		mock_point := point{}
		// read until \n
		line, _, err := testcase.ReadLine()
		if err != nil {
			// if file ends we shut down for loop
			log.Println(err)
			break
		}
		sides := strings.Split(string(line), "->")
		if len(sides) == 2 {
			from := strings.Split(sides[0], ",")
			to := strings.Split(sides[1], ",")
			if len(from) == 2 && len(to) == 2 {
				mock_point.from_x, _ = strconv.Atoi(strings.TrimSpace(from[0]))
				mock_point.from_y, _ = strconv.Atoi(strings.TrimSpace(from[1]))
				mock_point.to_x, _ = strconv.Atoi(strings.TrimSpace(to[0]))
				mock_point.to_y, _ = strconv.Atoi(strings.TrimSpace(to[1]))

				max := findMax(mock_point.from_x,
					mock_point.from_y,
					mock_point.to_x,
					mock_point.to_y)
				if max > mapsize {
					mapsize = max
				}
				points = append(points, mock_point)
			}
		}
	}

	game_map := processToMap(points, mapsize+1)
	printMap(game_map)

	// finaly we have answer !!!
	return calculateDanger(game_map)
}

func printMap(game_map [][]int) {
	for _, arr := range game_map {
		for _, item := range arr {
			if item == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(item)
			}
		}
		fmt.Println()
	}
}

func calculateDanger(game_map [][]int) int {
	var danger int
	for _, arr := range game_map {
		for _, num := range arr {
			if num > 1 {
				danger++
			}
		}
	}
	return danger
}

func processToMap(points []point, mapsize int) [][]int {
	var game_map = make([][]int, mapsize)
	for index := range game_map {
		game_map[index] = make([]int, mapsize)
	}
	for _, pnt := range points {
		if pnt.from_x == pnt.to_x || pnt.from_y == pnt.to_y {
			markHVPoints(&game_map, pnt)
		} else if pnt.from_x == pnt.from_y && pnt.to_x == pnt.to_y {
			mark45DegreeType1(&game_map, pnt)
		} else if abs(pnt.from_x-pnt.to_x) == abs(pnt.from_y-pnt.to_y) {
			mark45DegreeType2(&game_map, pnt)
		}
	}
	return game_map
}

func mark45DegreeType1(game_map *[][]int, pnt point) {
	for i := 0; i <= abs(pnt.from_x-pnt.to_x); i++ {
		mark := findMin(pnt.from_x, pnt.to_x) + i
		(*game_map)[mark][mark]++
	}
}

func mark45DegreeType2(game_map *[][]int, pnt point) {
	if pnt.from_y < pnt.to_y {
		if pnt.from_x < pnt.to_x {
			for i := 0; i <= abs(pnt.from_x-pnt.to_x); i++ {
				(*game_map)[pnt.from_y+i][pnt.from_x+i]++
			}
		} else {
			for i := 0; i <= abs(pnt.from_x-pnt.to_x); i++ {
				(*game_map)[pnt.from_y+i][pnt.from_x-i]++
			}
		}
	} else {
		if pnt.from_x < pnt.to_x {
			for i := 0; i <= abs(pnt.from_x-pnt.to_x); i++ {
				(*game_map)[pnt.from_y-i][pnt.from_x+i]++
			}
		} else {
			for i := 0; i <= abs(pnt.from_x-pnt.to_x); i++ {
				(*game_map)[pnt.from_y-i][pnt.from_x-i]++
			}
		}
	}
}

func markHVPoints(game_map *[][]int, pnt point) {
	if pnt.from_y == pnt.to_y {
		for i := 0; i <= abs(pnt.from_x-pnt.to_x); i++ {
			(*game_map)[pnt.from_y][findMin(pnt.from_x, pnt.to_x)+i]++
		}
	}
	if pnt.from_x == pnt.to_x {
		for i := 0; i <= abs(pnt.from_y-pnt.to_y); i++ {
			(*game_map)[findMin(pnt.from_y, pnt.to_y)+i][pnt.from_x]++
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

type point struct {
	from_x, from_y int
	to_x, to_y     int
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
