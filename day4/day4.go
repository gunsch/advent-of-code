package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const ( // iota is reset to 0
	state_no_board      = iota // c0 == 0
	state_reading_board = iota // c1 == 1
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Read the bingo order
	scanner.Scan()
	first_input := scanner.Text()

	bingo_order_strs := strings.Split(first_input, ",")
	bingo_order := strings_to_nums(bingo_order_strs)
	fmt.Println(bingo_order)

	// Read the boards
	bingo_boards := [][][]int{}
	current_board_id := 0
	current_board_row := 0
	current_board := [][]int{}

	re := regexp.MustCompile("[0-9]+")

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			if current_board_row > 0 {
				current_board_id++
			}
			if len(current_board) == 5 {
				bingo_boards = append(bingo_boards, current_board)
			}
			current_board_row = 0
			current_board = [][]int{}
			continue
		}

		all_numbers := re.FindAllString(line, -1)
		current_board = append(current_board, strings_to_nums(all_numbers))
		fmt.Printf("%d/%d: %s\n", current_board_id, current_board_row, all_numbers)
		current_board_row++
	}
	bingo_boards = append(bingo_boards, current_board)

	fmt.Printf("we have %d boards\n", len(bingo_boards))
	fmt.Println(bingo_boards[1])

	// Play bingo!
	played_bingo_nums := make(map[int]bool)

	// Part 1
	var winning_board [][]int
	winning_number := 0

	// Part 2
	// Map of board ID to finished
	finished_boards := make(map[int]bool)
	var final_board [][]int
	final_number := 0

bingo:
	for _, v := range bingo_order {
		fmt.Printf("Evaluating %d\n", v)
		played_bingo_nums[v] = true
		for board_id, board := range bingo_boards {
			if find_bingo(board, played_bingo_nums) {
				if !finished_boards[board_id] {
					finished_boards[board_id] = true
					fmt.Printf("Finished board %d\n", board_id)

					// Winner for part 1
					if len(finished_boards) == 1 {
						winning_board = board
						winning_number = v

						// Part 1
						fmt.Println("Part 1: found a winner!")
						fmt.Println(winning_board)
						score := calculate_score(winning_board, played_bingo_nums)
						fmt.Printf("score = %d\n", score)
						total := winning_number * score
						fmt.Printf("total = %d\n", total)
					}

					// Last board
					if len(finished_boards) == len(bingo_boards) {
						final_board = board
						final_number = v

						// Part 2
						fmt.Println("Part 2: found a winner!")
						fmt.Println(final_board)
						score2 := calculate_score(final_board, played_bingo_nums)
						fmt.Printf("score = %d\n", score2)
						total2 := final_number * score2
						fmt.Printf("total = %d\n", total2)

						break bingo
					}
				}
			}
		}
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}

func calculate_score(winning_board [][]int, played_bingo_nums map[int]bool) int {
	total := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			number := winning_board[i][j]
			if !played_bingo_nums[number] {
				total += number
			}
		}
	}
	return total
}

func find_bingo(board [][]int, pbm map[int]bool) bool {
	// Check each row and column
	for i := 0; i < 5; i++ {
		// Check row [i]
		if pbm[board[i][0]] && pbm[board[i][1]] && pbm[board[i][2]] && pbm[board[i][3]] && pbm[board[i][4]] {
			return true
		}
		// Check col [i]
		if pbm[board[0][i]] && pbm[board[1][i]] && pbm[board[2][i]] && pbm[board[3][i]] && pbm[board[4][i]] {
			return true
		}
	}
	return false
}

func strings_to_nums(strs []string) []int {
	nums_array := []int{}
	for _, v := range strs {
		val, _ := strconv.Atoi(v)
		nums_array = append(nums_array, val)
	}
	return nums_array
}
