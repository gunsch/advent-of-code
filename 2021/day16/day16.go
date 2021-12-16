package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
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

var hexMap = map[rune]byte{
	'0': 0,
	'1': 0x1,
	'2': 0x2,
	'3': 0x3,
	'4': 0x4,
	'5': 0x5,
	'6': 0x6,
	'7': 0x7,
	'8': 0x8,
	'9': 0x9,
	'A': 0xA,
	'B': 0xB,
	'C': 0xC,
	'D': 0xD,
	'E': 0xE,
	'F': 0xF,
}

type bitReader struct {
	byteArr []byte
	bitPosition uint
}

var totalVersionScore = 0

func main() {

	// regex_insertion := regexp.MustCompile(`(?P<x1>[A-Z][A-Z]) -> (?P<x3>[A-Z])`)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanRunes)

	// D2FE28

	// row := 0
	reader := bitReader{
		byteArr: []byte{},
		bitPosition: 0,
	}
	for scanner.Scan() {
		input := scanner.Text()
		fmt.Println(input)
		reader.byteArr = append(reader.byteArr, hexMap[[]rune(input)[0]])
	}
	fmt.Println(reader.byteArr)

	readPackets(&reader)
	fmt.Printf("total version score = %d\n", totalVersionScore)
}

// Returns number of bits read, and value
func readPackets(reader *bitReader) (uint, uint) {
	startPosition := reader.bitPosition
	// Read 3 bits
	packetVersion := readBits(reader, 3)
	totalVersionScore += int(packetVersion)
	packetTypeId := readBits(reader, 3)
	packetValue := uint(0)
	// fmt.Printf("version: %d, type ID: %d\n", packetVersion, packetTypeId)
	for {

		if packetTypeId == 4 {
			literalValue := uint(0)
			for {
				literalValue = literalValue << 4
				byteValue := readBits(reader, 5)
				literalValue += byteValue % 16
				// End of the line
				if byteValue & 16 == 0 {
					break
				}
			}

			packetValue = literalValue
			// fmt.Printf("Read literal: %d\n", literalValue)
			break
		}

		subPacketValues := []uint{}

		lengthTypeId := readBits(reader, 1)
		if lengthTypeId == 0 {
			totalLengthOfBits := readBits(reader, 15)
			fmt.Printf("typeid=0, totalLengthOfBits=%d\n", totalLengthOfBits)

			subBitsRead := uint(0)
			for subBitsRead < totalLengthOfBits {
				bitsRead, subPacketValue := readPackets(reader)
				subBitsRead += bitsRead
				subPacketValues = append(subPacketValues, subPacketValue)
			}
		} else {
			numberOfSubPackets := readBits(reader, 11)
			fmt.Printf("typeid=1, numberOfSubPackets=%d\n", numberOfSubPackets)
			for i := 0; i < int(numberOfSubPackets); i++ {
				_, subPacketValue := readPackets(reader)
				subPacketValues = append(subPacketValues, subPacketValue)
			}
		}
	
		if packetTypeId == 0 {
			packetValue = sum(subPacketValues)
		} else if packetTypeId == 1 {
			packetValue = product(subPacketValues)
		} else if packetTypeId == 2 {
			packetValue = minimum(subPacketValues)
		} else if packetTypeId == 3 {
			packetValue = maximum(subPacketValues)
		} else if packetTypeId == 5 {
			packetValue = greaterThan(subPacketValues)
		} else if packetTypeId == 6 {
			packetValue = lessThan(subPacketValues)
		} else if packetTypeId == 7 {
			packetValue = equalTo(subPacketValues)
		} else {
			log.Fatalf("unknown packet type id: %d\n", packetTypeId)
		}

		fmt.Printf("type=%d, subpacket=%v, output=%d\n", packetTypeId, subPacketValues, packetValue)
		break
	}

	bitsRead := reader.bitPosition - startPosition
	// fmt.Printf("version=%d, bitsRead=%d, packetValue=%d\n", packetVersion, bitsRead, packetValue)
	return bitsRead, packetValue
}

func readBits(reader *bitReader, bitsToRead uint) uint {
	value := uint(0)
	bitsLeftToRead := bitsToRead

	for {
		if bitsLeftToRead == 0 {
			break
		}

		currentByteIdx := reader.bitPosition / 4
		currentByte := reader.byteArr[currentByteIdx]
		numBitsReadInCurrentByte := reader.bitPosition % 4
		numBitsRemainingInCurrentByte := 4 - numBitsReadInCurrentByte

		// fmt.Printf("value=%d, currentByte=%d/%d, bitPosition=%d\n", value, currentByteIdx, currentByte, reader.bitPosition)

		if bitsLeftToRead < numBitsRemainingInCurrentByte {
			// Shift it to make room, even if awkward number of bytes
			value = value << bitsLeftToRead

			valueRead := readBitsFromByte(currentByte, bitsLeftToRead, numBitsReadInCurrentByte)
			value += uint(valueRead)
			// fmt.Printf("(a) readBitsFromByte(%d, %d, %d), valueRead=%d, value=%d\n", currentByte, bitsLeftToRead, numBitsReadInCurrentByte, valueRead, value)
			reader.bitPosition += uint(bitsLeftToRead)
			break
		}

		if numBitsRemainingInCurrentByte < 4 {
			// I think this should only happen on the first run-through here
			if value > 0 {
				log.Fatalf("numBitsremainingInCurrentByte=%d, but value=%d\n", numBitsRemainingInCurrentByte, value)
			}

			valueRead := readBitsFromByte(currentByte, numBitsRemainingInCurrentByte, numBitsReadInCurrentByte)
			reader.bitPosition += numBitsRemainingInCurrentByte
			bitsLeftToRead -= numBitsRemainingInCurrentByte

			value = uint(valueRead)
			// fmt.Printf("(b) readBitsFromByte(%d, %d, %d). valueRead=%d, value=%d\n", currentByte, numBitsRemainingInCurrentByte, numBitsReadInCurrentByte, valueRead, value)
			continue
		}

		if bitsLeftToRead >= 4 && numBitsRemainingInCurrentByte == 4 {
			// Shift it ahead, to make room for the new!
			value = value << 4

			valueRead := readBitsFromByte(currentByte, 4, 0)
			value += uint(valueRead)
			reader.bitPosition += 4
			bitsLeftToRead -= 4
			// fmt.Printf("(c) readBitsFromByte(%d, 4, 0), valueRead=%d, value=%d\n", currentByte, valueRead, value)
			continue
		}

		log.Fatalf("oops! bitsLeftToRead=%d, numBitsRemainingInCurrentByte=%d\n", bitsLeftToRead, numBitsRemainingInCurrentByte)
	}

	return value
}

func readBitsFromByte(byteVal byte, nbits, start uint) byte {
	if nbits > 4 - start {
		log.Fatalf("tried to read %d bits after pos %d\n", nbits, start)
	}

	if start > 0 {
		// Zero out left of it
		byteVal = byteVal % byte(math.Pow(2, float64(4 - start)))
		byteVal = byteVal << start
	}

	if nbits < 4 {
		byteVal = byteVal >> (4 - nbits)
	}

	return byteVal
}



func equalTo(subPacketValues []uint) uint {
	if subPacketValues[0] == subPacketValues[1] {
		return 1
	} else {
		return 0
	}
}

func lessThan(subPacketValues []uint) uint {
	if subPacketValues[0] < subPacketValues[1] {
		return 1
	} else {
		return 0
	}
}

func greaterThan(subPacketValues []uint) uint {
	if subPacketValues[0] > subPacketValues[1] {
		return 1
	} else {
		return 0
	}
}

func maximum(subPacketValues []uint) uint {
	maximum := uint(0)
	for _, v := range subPacketValues {
		if v > maximum {
			maximum = v
		}
	}
	return maximum
}

func minimum(subPacketValues []uint) uint {
	minimum := uint(18446744073709551615)
	for _, v := range subPacketValues {
		if v < minimum {
			minimum = v
		}
	}
	return minimum
}

func product(subPacketValues []uint) uint {
	product := uint(1)
	for _, v := range subPacketValues {
		product = product * v
	}
	return product
}

func sum(subPacketValues []uint) uint {
	sum := uint(0)
	for _, v := range subPacketValues {
		sum += v
	}
	return sum
}
