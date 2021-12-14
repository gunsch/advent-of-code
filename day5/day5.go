package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {

	size := 1000

	scanner := bufio.NewScanner(os.Stdin)
	regex_read := regexp.MustCompile(`(?P<x1>\d+),(?P<y1>\d+) -> (?P<x2>\d+),(?P<y2>\d+)`)

	number_lines := make([][]int, size)
	for i := range number_lines {
		number_lines[i] = make([]int, size)
	}

	all_sets_of_lines := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		coordinates := strings_to_nums(regex_read.FindStringSubmatch(line)[1:])
		fmt.Println(coordinates)
		all_sets_of_lines = append(all_sets_of_lines, coordinates)

		record_points(number_lines, coordinates)
	}

	// Count them all
	count_most := 0
	for i := 0; i < len(number_lines); i++ {
		for j := 0; j < len(number_lines[i]); j++ {
			if number_lines[i][j] >= 2 {
				count_most++
			}
		}
	}
	// for _, row := range number_lines {
	// 	for _, val := range row {
	// 		if val >= 2 {
	// 			count_most++
	// 		}
	// 	}
	// }
	fmt.Printf("total >= 2 = %d\n", count_most)

	// fmt.Println(all_sets_of_lines)
	// fmt.Println(number_lines)

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

}

func record_points(number_lines [][]int, coordinates []int) {
	// X-version
	if coordinates[1] == coordinates[3] {
		y_coordinate := coordinates[1]
		x1, x2 := sort(coordinates[0], coordinates[2])
		fmt.Printf("Filling row %d from %d to %d\n", y_coordinate, x1, x2)
		for i := x1; i <=  x2; i++ {
			number_lines[i][y_coordinate]++
		}
		return
	}

	// Y-version
	if coordinates[0] == coordinates[2] {
		x_coordinate := coordinates[0]
		y1, y2 := sort(coordinates[1], coordinates[3])
		fmt.Printf("Filling col %d from %d to %d\n", x_coordinate, y1, y2)
		for i := y1; i <= y2; i++ {
			number_lines[x_coordinate][i]++
		}
		return
	}

	// Diagonal version
	x1, x2 := coordinates[0], coordinates[2]
	y1, y2 := coordinates[1], coordinates[3]
	x_order := 1
	y_order := 1
	if x2 < x1 {
		x_order = -1
	}
	if y2 < y1 {
		y_order = -1
	}
	i := x1
	j := y1
	last_loop := false
	fmt.Printf("Filling diagonally %d,%d to %d,%d. order %d,%d\n", x1, y1, x2, y2, x_order, y_order)
	for {
		if j == y2 { // we're done
			last_loop = true
			if i != x2 {
				log.Fatal("noooooo")
			}
		}

		number_lines[i][j]++

		i += x_order
		j += y_order
		if last_loop {
			break
		}
	}
}

func strings_to_nums(strs []string) []int {
	nums_array := []int{}
	for _, v := range strs {
		val, _ := strconv.Atoi(v)
		nums_array = append(nums_array, val)
	}
	return nums_array
}

func sort(a int, b int) (int, int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}
