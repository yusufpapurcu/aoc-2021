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

	// // reading file
	// k, err := os.Open("examplecase.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer k.Close()

	// // basic example for testing
	// log.Println(solution(k))

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

	var bingo = func() []int {
		line, _, err := testcase.ReadLine()
		if err != nil {
			// if file ends we shut down for loop
			log.Println(err)
			return nil
		}
		bingo_str := strings.Split(string(line), ",")

		var res = make([]int, len(bingo_str))
		for index, value := range bingo_str {
			temp, _ := strconv.Atoi(value)
			res[index] = temp
		}
		return res
	}()
	_ = bingo
	var game_tables = []table{}
	for {
		// read until \n
		line, _, err := testcase.ReadLine()
		if err != nil {
			// if file ends we shut down for loop
			log.Println(err)
			break
		}
		if len(line) == 0 {
			game_tables = append(game_tables, func() table {
				var res = table{}
				res.row_points = make(map[int]int)
				res.column_points = make(map[int]int)
				res.rows = make(map[int][]point)

				for i := 0; i < 5; i++ {
					// read until \n
					line, _, err := testcase.ReadLine()
					if err != nil {
						// if file ends we shut down for loop
						log.Println(err)
						return table{}
					}
					fields := strings.Fields(string(line))

					res.rows[i] = make([]point, len(fields))
					for index, value := range fields {
						temp, _ := strconv.Atoi(value)
						res.rows[i][index] = point{
							isMarked: false,
							num:      temp,
						}
					}
				}

				return res
			}())
		}
	}

	res, luckyNum := start_game(game_tables, bingo)
	if res.rows == nil {
		log.Println("No winner")
		return 0
	}

	sumUnmarked := sumTheUnmarkeds(res)

	// finaly we have answer !!!
	return (sumUnmarked) * (luckyNum)
}

func sumTheUnmarkeds(winner table) int {
	var res int
	for _, row := range winner.rows {
		for _, pnt := range row {
			if !pnt.isMarked {
				res += pnt.num
			}
		}
	}
	return res
}

func start_game(tables []table, bingo []int) (table, int) {
	table_count := len(tables)
	for _, called := range bingo {
		isWinner, winnerTable, luckyNum := game_turn(&tables, called, table_count)
		if isWinner {
			table_count -= luckyNum

			if table_count < 4 {
				fmt.Print("")
			}
			if table_count == 0 {
				return *winnerTable, called
			}
		}
	}
	return table{}, 0
}

func game_turn(tables *[]table, called int, table_count int) (bool, *table, int) {
	var winner = false
	var deleted = 0
	for table_index, player := range *tables {
		if player.rows == nil {
			continue
		}
		for row_index, row := range player.rows {
			for point_index, num := range row {
				if num.num == called && !num.isMarked {
					(*tables)[table_index].rows[row_index][point_index].isMarked = true

					player.row_points[row_index]++
					player.column_points[point_index]++
					if player.row_points[row_index] == 5 || player.column_points[point_index] == 5 {
						if table_count-deleted != 1 {
							(*tables)[table_index].rows = nil
							winner = true
							deleted++
						} else {
							return true, &player, 1
						}
					}
				}
			}
		}
	}
	return winner, &table{}, deleted
}

type point struct {
	num      int
	isMarked bool
}

type table struct {
	rows          map[int][]point
	row_points    map[int]int
	column_points map[int]int
}
