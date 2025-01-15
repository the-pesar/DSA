package main

import (
	"fmt"
	"strconv"
)

// two pointer
func AnagramNumber(number int) bool {
	strNumber := strconv.Itoa(number)

	var leftPtr, rightPtr = 0, len(strNumber) - 1

	for (leftPtr != rightPtr && leftPtr < rightPtr) {
		if strNumber[leftPtr] != strNumber[rightPtr] {
			return false
		}

		leftPtr++
		rightPtr--
	}

	return true
}

func main() {
	fmt.Println(AnagramNumber(21312))
}