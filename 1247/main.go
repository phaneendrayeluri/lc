package main

import (
	"log"
)

func main() {
	log.Printf("Total Swaps: %d", minimumSwap("xx", "yy"))
	log.Printf("Total Swaps: %d", minimumSwap("xy", "yx"))
}

func minimumSwap(s1 string, s2 string) int {
	noOfSwaps := 0
	var x byte
	var y byte

	for i := 0; i <= len(s1)-1; i++ {
		if s1[i] != s2[i] {
			switch s1[i] {
			case 120:
				if x != 0 {
					noOfSwaps++
					x = 0
				} else {
					x = s1[i]
				}
			case 121:
				if y != 0 {
					noOfSwaps++
					y = 0
				} else {
					y = s1[i]
				}
			}
		}
	}

	if x == 0 && y == 0 {
		return noOfSwaps
	}
	if x != 0 && y != 0 {
		return noOfSwaps + 2
	}

	return -1
}
