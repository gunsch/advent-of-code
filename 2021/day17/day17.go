package main

import "fmt"

type point struct {
	x int
	y int
}

func main() {

	// target area: x=244..303, y=-91..-54
	// input puzzle:
	// target area: x=20..30, y=-10..-5

	doPuzzle(20, 30, -10, -5)
	doPuzzle(244, 303, -91, -54)
}

func doPuzzle(x1, x2, y1, y2 int) {
	var topLeft = point{x: x1, y: y2}
	var bottomRight = point{x: x2, y: y1}
	
	highPoint := 0
	successCount := 0
	for i := -305; i < 305; i++ {
		for j := -305; j < 305; j++ {
			probeSuccess, probeHighPoint := launchProbe(point{i, j}, topLeft, bottomRight)
			if probeSuccess {
				successCount++
			}
			if probeSuccess && probeHighPoint > highPoint {
				fmt.Printf("NEW HIGH: %d for %v\n", probeHighPoint, point{i,j})
				highPoint = probeHighPoint
			}
		}
	}
	fmt.Printf("Found %d total probe successes\n", successCount)
}

func launchProbe(originalVelocity, topLeft, bottomRight point) (bool, int) {
	currentPosition := point{x: 0, y: 0}
	steps := []point{currentPosition}
	currentVelocity := originalVelocity
	highPoint := 0
	success := false
	for {
		// fmt.Printf("position: %v, velocity: %v\n", currentPosition, currentVelocity)

		if currentPosition.y > highPoint {
			highPoint = currentPosition.y
		}

		if currentPosition.x >= topLeft.x &&
			currentPosition.x <= bottomRight.x &&
			currentPosition.y <= topLeft.y &&
			currentPosition.y >= bottomRight.y {
			// fmt.Printf("SCORE: velocity (%v) hit spot (%v) in the square. high point %d\n", originalVelocity, currentPosition, highPoint)
			success = true
			break
			}

		if currentPosition.x > bottomRight.x ||
			currentPosition.y < bottomRight.y {
			// fmt.Printf("FAIL: velocity %v, exiting because (%v) past bottomRight (%v)\n", originalVelocity, currentPosition, bottomRight)
			success = false
			break
		}

		currentPosition, currentVelocity = moveStep(currentPosition, currentVelocity)
		steps = append(steps, currentPosition)
	}

	return success, highPoint
}

func moveStep(currentPosition, velocity point) (point, point) {
	var newPosition = point{
		x: currentPosition.x + velocity.x,
		y: currentPosition.y + velocity.y,
	}
	var newVelocity = point{
		x: 0,
		y: velocity.y - 1,
	}
	if velocity.x < 0 {
		newVelocity.x = velocity.x + 1
	} else if velocity.x > 0 {
		newVelocity.x = velocity.x - 1
	}
	return newPosition, newVelocity
}
