package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)


type fold_instruction struct {
	is_x bool
	value int
}

type point struct {
	x int
	y int
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	regex_fold := regexp.MustCompile(`fold along (?P<dir>[xy])=(?P<value>\d+)`)

	points := []point {}
	fold_instructions := []fold_instruction {}

	for scanner.Scan() {
		input := scanner.Text()

		// fmt.Printf("parsing: %s\n", input)

		if regex_fold.MatchString(input) {
			fmt.Println("is a fold")
			matches := regex_fold.FindStringSubmatch(input)
			value, _ := strconv.Atoi(matches[2])
			fold_instructions = append(fold_instructions, fold_instruction{is_x: matches[1] == "x", value: value})
			continue
		}

		if len(input) == 0 {
			continue
		}
		
		xy_text := strings.Split(input, ",")
		x, _ := strconv.Atoi(xy_text[0])
		y, _ := strconv.Atoi(xy_text[1])
		points = append(points, point{x: x, y: y})
	}

	fmt.Println(points)
	fmt.Println(fold_instructions)

	for _, instr := range fold_instructions {
		points = execute_fold(points, instr)
		fmt.Printf("there are %d points left\n", len(points))
		print_paper(points)
	}

}

func print_paper(points []point) {
	size := largest_value(points) + 1
	blah := make([][]bool, size)

	for _, point := range points {
		if blah[point.y] == nil {
			blah[point.y] = make([]bool, size)
		}
		blah[point.y][point.x] = true
	}
	for _, row := range blah {
		for _, truthy := range row {
			if truthy {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func largest_value(points []point) int {
	largest := 0
	for _, point := range points {
		if point.x > largest {
			largest = point.x
		}
		if point.y > largest {
			largest = point.y
		}
	}
	return largest
}



func execute_fold(points []point, instr fold_instruction) []point {
	updated_points := make([]point, len(points))
	copy(updated_points, points)

	if instr.is_x {
		fmt.Printf("folding horizontally at x=%d\n", instr.value)
		for index, point := range updated_points {
			if point.x > instr.value {
				new_x := point.x - 2 * (point.x - instr.value)
				point.x = new_x
				updated_points[index] = point
			}
		}
	} else {
		fmt.Printf("folding vertically at y=%d\n", instr.value)
		for index, point := range updated_points {
			if point.y > instr.value {
				new_y := point.y - 2 * (point.y - instr.value)
				// fmt.Printf("moving %d,%d to %d,%d\n", point.x, point.y, point.x, new_y)
				point.y = new_y
				updated_points[index] = point
			}
		}
	}

	// fmt.Println("updated all values")
	// fmt.Println(updated_points)
	return removeDuplicateValues(updated_points)
}

func removeDuplicateValues(points []point) []point {
    keys := make(map[point]bool)
    list := []point{}
     for _, entry := range points {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }
    return list
}

func strings_to_nums(strs []string) []int {
	nums_array := []int{}
	for _, v := range strs {
		val, _ := strconv.Atoi(v)
		nums_array = append(nums_array, val)
	}
	return nums_array
}
