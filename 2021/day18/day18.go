package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

type node struct {
	value  int
	left   *node
	right  *node
	parent *node
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	totalNode, _ := buildNumber(input, 0)
	printNode(*totalNode, true)

	allInputs := []string{input}

	for scanner.Scan() {
		input := scanner.Text()
		newNode, _ := buildNumber(input, 0)
		printNode(*newNode, true)
		allInputs = append(allInputs, input)

		// Add the two together
		totalNode = addNumbers(totalNode, newNode)

		fmt.Printf("reducing %s\n", sprint(*totalNode))
		reduceNumber(totalNode)
		fmt.Printf("result: %s\n", sprint(*totalNode))
		fmt.Println()
	}

	magnitude := calculateMagnitude(totalNode)
	fmt.Printf("magnitude=%d\n", magnitude)

	largestMagnitude := 0
	for _, node1 := range allInputs {
		for _, node2 := range allInputs {
			if node1 == node2 {
				continue
			}
			leftNode, _ := buildNumber(node1, 0)
			rightNode, _ := buildNumber(node2, 0)
			// Otherwise... try adding
			sumNode := addNumbers(leftNode, rightNode)
			reduceNumber(sumNode)
			newMagnitude := calculateMagnitude(sumNode)
			if newMagnitude > largestMagnitude {
				largestMagnitude = newMagnitude
			}
		}
	}
	fmt.Printf("largest magnitude = %d\n", largestMagnitude)

	// fmt.Println()
	// printNode(*testInput, true)
	// explodeNestedPairs(testInput, 0)
	// printNode(*testInput, true)
	// fmt.Println()

	// testInput2, _ := buildNumber("[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]", 0)
	// reduceNumber(testInput2)
	// printNode(*testInput2, true)
	// testInput, _ := buildNumber("[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", 0)
	// magnitude := calculateMagnitude(testInput)
	// fmt.Printf("magnitude=%d\n", magnitude)
}

var COMMAND_NONE = -2
var COMMAND_EXPLODE = -3

func reduceNumber(number *node) (node, int) {
	for {
		// If any pair is nested inside four pairs, the leftmost such pair explodes.
		if explodeNestedPairs(number, 0) {
			continue
		}

		// If we didn't do an explosion, split the first large number we find
		if splitLargeNumbers(number) {
			continue
		}

		// If neither thing happened, then we're done.
		break
	}

	return *number, 0
}

func explodeNestedPairs(current *node, depth int) bool {
	if current.value != -1 {
		return false
	}

	// DFS but prioritize left-most
	if explodeNestedPairs(current.left, depth+1) {
		return true
	}
	if explodeNestedPairs(current.right, depth+1) {
		return true
	}

	// If we found something >4 levels deep... do an explode!
	if depth == 4 {
		// Implementing explode is gonna suck....
		// TODO
		fmt.Printf("Exploding %v, parent=%v\n", sprint(*current), sprint(*current.parent))
		addToLeft(current, current.left.value)
		addToRight(current, current.right.value)

		replacementNode := node{value: 0, parent: current.parent}
		replaceNode(current, &replacementNode)
		return true
	}

	return false
}

func addToLeft(current *node, value int) {
	// fmt.Printf("addToLeft, current=%s\n", sprint(*current))

	// Base case
	if current.parent == nil {
		// fmt.Printf("bailed out, base-cased\n")
		return
	}

	// If we're already the left branch, go up another branch
	if current.parent.left == current {
		addToLeft(current.parent, value)
		return
	}

	// Otherwise, if we find it... time to recurse down for the rightmost
	// current.parent.left != current
	nextNode := current.parent.left
	for {
		if nextNode.value > -1 {
			// fmt.Printf("addToLeft: found it! adding %d to %d\n", value, nextNode.value)
			nextNode.value += value
			return
		}
		// Descend down the right until found
		nextNode = nextNode.right
	}
}

func addToRight(current *node, value int) {
	// fmt.Printf("addToRight, current=%s\n", sprint(*current))

	// Base case
	if current.parent == nil {
		// fmt.Printf("bailed out, base-cased\n")
		return
	}

	// If we're already the right branch, go up another branch
	if current.parent.right == current {
		addToRight(current.parent, value)
		return
	}

	// Otherwise, if we find it... time to recurse down for the leftmost
	// current.parent.right != current
	nextNode := current.parent.right
	for {
		if nextNode.value > -1 {
			// fmt.Printf("addToRight: found it! adding %d to %d\n", value, nextNode.value)
			nextNode.value += value
			return
		}
		// Descend down the left until found
		nextNode = nextNode.left
	}
}

func replaceNode(current *node, newNode *node) {
	newNode.parent = current.parent
	if current.parent.left == current {
		current.parent.left = newNode
	} else if current.parent.right == current {
		current.parent.right = newNode
	} else {
		log.Fatalf("wtf!?")
	}
}

func splitLargeNumbers(current *node) bool {
	if current.value >= 10 {
		fmt.Printf("doing a split on %v\n", sprint(*current))
		leftValue := current.value / 2
		rightValue := current.value / 2 + current.value % 2
		// newCurrentNode, _ := buildNumber(fmt.Sprintf("[%d,%d]", leftValue, rightValue), 0)
		leftNode := node{value: leftValue}
		rightNode := node{value: rightValue}
		newCurrentNode := addNumbers(&leftNode, &rightNode)
		replaceNode(current, newCurrentNode)
		return true
	}

	if current.value != -1 {
		return false
	}

	if splitLargeNumbers(current.left) {
		return true
	}
	if splitLargeNumbers(current.right) {
		return true
	}

	return false
}

func addNumbers(left, right *node) *node {
	newNode := node{
		left:   left,
		right:  right,
		parent: nil,
		value:  -1,
	}
	left.parent = &newNode
	right.parent = &newNode
	return &newNode
}

func buildNumber(input string, startPos int) (*node, int) {
	inputRunes := []rune(input)
	nextChar := inputRunes[startPos]
	if nextChar == '[' {
		leftChild, nextPos := buildNumber(input, startPos+1)
		if inputRunes[nextPos] != ',' {
			log.Fatalf("Expected , at position %d, found %v\n", nextPos, inputRunes[nextPos])
		}
		rightChild, nextPos := buildNumber(input, nextPos+1)
		if inputRunes[nextPos] != ']' {
			log.Fatalf("Expected ] at position %d, found %v\n", nextPos, inputRunes[nextPos])
		}
		newNode := node{
			value:  -1,
			parent: nil,
			left:   leftChild,
			right:  rightChild,
		}
		leftChild.parent = &newNode
		rightChild.parent = &newNode
		return &newNode, nextPos + 1
	} else if unicode.IsNumber(nextChar) {
		nextPos := startPos + 1
		charAsInt, _ := strconv.Atoi(string(nextChar))
		return &node{value: charAsInt}, nextPos
	} else {
		log.Fatalf("got unexpected rune %v\n", nextChar)
	}
	return &node{}, -1
}

func calculateMagnitude(current* node) int {
	if current.value >= 0 {
		return current.value
	}
	return 3 * calculateMagnitude(current.left) + 2 * calculateMagnitude(current.right)
}

func printTree(toPrint *node, depth int) {
	spacing := ""
	for i := 0; i < depth; i++ {
		spacing += "  "
	}
	fmt.Printf("%s [%p] value=%d, left=%p, right=%p, parent=%p\n",
		spacing, toPrint, toPrint.value, toPrint.left, toPrint.right, toPrint.parent)
	if toPrint.value == -1 {
		printTree(toPrint.left, depth+1)
		printTree(toPrint.right, depth+1)
	}
}

func printNode(toPrint node, start bool) {
	if toPrint.value > -1 {
		fmt.Printf("%d", toPrint.value)
		return
	}

	if start != true && toPrint.parent == nil {
		log.Fatalf("missing parent! %s\n", sprint(toPrint))
	}

	fmt.Printf("[")
	printNode(*toPrint.left, false)
	fmt.Printf(",")
	printNode(*toPrint.right, false)
	fmt.Printf("]")
	if start {
		fmt.Println()
	}
}

func sprint(toPrint node) string {
	if toPrint.value > -1 {
		return fmt.Sprintf("%d", toPrint.value)
	}
	ret := "[" + sprint(*toPrint.left) + "," + sprint(*toPrint.right) + "]"
	return ret
}
