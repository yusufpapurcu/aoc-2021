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
	log.Println(solution(strings.NewReader("00100\n11110\n10110\n10111\n10101\n01111\n00111\n11100\n10000\n11001\n00010\n01010")))

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
func solution(s io.Reader) int64 {
	// conver reader to bufio type
	// not necessary I added for process big inputs easily
	testcase := bufio.NewReader(s)
	var list []string
	for {
		// read until \n
		line, _, err := testcase.ReadLine()
		if err != nil {

			// if file ends we shut down for loop
			log.Println(err)
			break
		}
		list = append(list, string(line))
	}

	oxygen_rating := bitCriteria(list, true)
	co2_rating := bitCriteria(list, false)

	oxygen_rating_int, _ := strconv.ParseInt(string(oxygen_rating), 2, 64)
	co2_rating_int, _ := strconv.ParseInt(string(co2_rating), 2, 64)

	// finaly we have answer !!!
	return (oxygen_rating_int) * (co2_rating_int)
}

func bitCriteria(bitlist []string, isHigh bool) string {
	templist := bitlist
	for i := 0; i < len(bitlist[0]); i++ {
		counts := func() map[int]int {
			var counts = map[int]int{}
			for _, item := range templist {
				for index, bit := range item {
					switch bit {
					case '0':
						if isHigh {
							counts[index] = counts[index] + 0
						} else {
							counts[index]++
						}
					case '1':
						if isHigh {
							counts[index]++
						} else {
							counts[index] = counts[index] + 0
						}
					default:
						panic("What the heck is " + string(bit))
					}
				}
			}
			return counts
		}()

		wanted := func() []byte {
			var wanted = make([]byte, len(counts))
			for key, value := range counts {
				if float64(value) > float64(len(templist))/2 {
					wanted[key] = '1'
				} else if float64(value) == float64(len(templist))/2 {
					wanted[key] = '2'
				} else {
					wanted[key] = '0'
				}
			}
			return wanted
		}()

		wanted = func() []byte {
			res := make([]byte, len(wanted))
			for index, bit := range wanted {
				if bit == '2' {
					if isHigh {
						res[index] = '1'
					} else {
						res[index] = '0'
					}
				} else {
					res[index] = bit
				}
			}
			return res
		}()

		templist = func(list []string) []string {
			var res []string
			for _, item := range list {
				if rune(item[i]) == rune(wanted[i]) {
					res = append(res, item)
				}
			}
			return res
		}(templist)

		if len(templist) == 1 {
			break
		}
	}
	return templist[0]
}
