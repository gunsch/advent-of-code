package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type fish struct {
	value int
	count int
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	input := scanner.Text()
	numbers := strings_to_nums(strings.Split(input, ","))

	fishes := make([]fish, len(numbers))
	for i, v := range numbers {
		fishes[i] = fish{value: v, count: 1}
	}

	for i := 0; i < 257; i++ {
		fishes_sum := 0
		for _, v := range fishes {
			fishes_sum += v.count
		}
		fmt.Printf("after %d days: %d\n", i, fishes_sum)

		n_news := 0
		for i, v := range fishes {
			if v.value == 0 {
				fishes[i] = fish{value: 6, count: v.count}
				n_news += v.count
			} else {
				fishes[i] = fish{value: v.value - 1, count: v.count}
			}
		}
		if n_news > 0 {
			fishes = append(fishes, fish{value: 8, count: n_news})
		}
	}
	
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
