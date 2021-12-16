package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

type point struct {
	x     int
	y     int
	score int
}

var NO_POINT = point{
	x:     -1,
	y:     -1,
	score: -1,
}

var SIZE = 100

func printO(octopi [][]int) {
	for _, y := range octopi {
		for _, x := range y {
			fmt.Printf(fmt.Sprintf("%d", x))
		}
		fmt.Println()
	}
}

func printO2(octopi [][]int) {
	for _, y := range octopi {
		for _, x := range y {
			fmt.Printf(fmt.Sprintf("%d ", x))
		}
		fmt.Println()
	}
}

func main() {

	// regex_insertion := regexp.MustCompile(`(?P<x1>[A-Z][A-Z]) -> (?P<x3>[A-Z])`)
	scanner := bufio.NewScanner(os.Stdin)

	risk_map := make([][]int, SIZE)
	for i := range risk_map {
		risk_map[i] = make([]int, SIZE)
	}

	row := 0
	for scanner.Scan() {
		input := scanner.Text()

		for i, v := range input {
			risk_map[row][i], _ = strconv.Atoi(string(v))
		}
		row++
	}
	printO(risk_map)

	// runDjikstras(risk_map)

	fmt.Println("Part 2")
	big_risk_map := buildBigRiskMap(risk_map)
	printO(big_risk_map)
	runDjikstras(big_risk_map)
}

func buildBigRiskMap(risk_map [][]int) [][]int {
	new_risk_map := make([][]int, SIZE * 5)
	for i := range new_risk_map {	
		new_risk_map[i] = make([]int, SIZE * 5)
	}

	fmt.Println("START")
	fmt.Println(len(new_risk_map))
	fmt.Println(len(new_risk_map[10]))

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			stampRiskMap(new_risk_map, risk_map, i, j)
		}
	}

	return new_risk_map
}

func stampRiskMap(new_risk_map [][]int, risk_map [][]int, i, j int) {
	startX := SIZE * i
	startY := SIZE * j
	offset := j + i
	// fmt.Printf("10/0=%d", new_risk_map[10][0])
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			new_value := (risk_map[i][j] + offset)
			if new_value > 9 {
				new_value -= 9
			}
			new_risk_map[startX + i][startY + j] = new_value
		}
	}
}

func runDjikstras(risk_map [][]int) {

	size := len(risk_map)
	distance_map := make([][]int, size)
	for i := range distance_map {
		distance_map[i] = make([]int, size)
		for j := range distance_map[i] {
			distance_map[i][j] = math.MaxInt32
		}
	}
	// Starting point
	distance_map[0][0] = 0 // risk_map[0][0]
	to_visit := []point{{0, 0, 0}}

	steps := 0
	for len(to_visit) > 0 && steps < 10000000 {
		current_point := to_visit[0]
		to_visit = to_visit[1:]

		// add all neighbors to the "next" list, starting from lowest cost
		// go to the lowest-cost node

		// fmt.Printf("current point: %d/%d (%d)\n", current_point.x, current_point.y, current_point.score)

		neighbor_points := generate_neighbor_points(current_point, risk_map, distance_map)

		// fmt.Println("neighbor points")
		// fmt.Println((neighbor_points))
		to_visit = append(to_visit, neighbor_points...)

		// fmt.Printf("to visit: %v\n", to_visit)

		// Re-sort to_visit by lowest points
		sort.Slice(to_visit, func(i, j int) bool {
			return to_visit[i].score < to_visit[j].score
		})

		steps++
	}

	fmt.Println("-------")
	printO2(distance_map)
}

func generate_neighbor_points(current_point point, risk_map, distance_map [][]int) []point {
	neighbor_points := []point{
		generate_neighbor_point(current_point, risk_map, distance_map, current_point.x+1, current_point.y),
		generate_neighbor_point(current_point, risk_map, distance_map, current_point.x-1, current_point.y),
		generate_neighbor_point(current_point, risk_map, distance_map, current_point.x, current_point.y-1),
		generate_neighbor_point(current_point, risk_map, distance_map, current_point.x, current_point.y+1),
	}

	filtered_neighbor_points := make([]point, 0, len(neighbor_points))
	for _, item := range neighbor_points {
		if item != NO_POINT {
			filtered_neighbor_points = append(filtered_neighbor_points, item)
		}
	}

	return filtered_neighbor_points
}

func generate_neighbor_point(current_point point, risk_map, distance_map [][]int, x, y int) point {
	// Out of bounds = not a neighbor
	if x < 0 || y < 0 || x >= len(distance_map) || y >= len(distance_map[x]) {
		return NO_POINT
	}

	// If it's a lower cost, update and put it back on the queue
	current_best_score_for_neighbor := distance_map[x][y]
	new_score_from_current_point := current_point.score + risk_map[x][y]

	// fmt.Printf("current_point %d,%d (%d) new_point %d,%d (%d), new best %d\n",
	// 	current_point.x, current_point.y, current_point.score,
	// 	x, y, current_best_score_for_neighbor, new_score_from_current_point)

	if new_score_from_current_point < current_best_score_for_neighbor {
		// Update distance map, and return visit
		fmt.Printf("UPDATE: %d,%d=%d\n", x, y, new_score_from_current_point)
		distance_map[x][y] = new_score_from_current_point
		return point{
			x:     x,
			y:     y,
			score: new_score_from_current_point,
		}
	}
	return NO_POINT
}
