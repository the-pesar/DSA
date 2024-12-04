package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func emptyAction(k uint8) {
	if k != s {
		return
	}

	if x == "L" {
		x = "R"
	} else {
		x = "L"
	}
}

func movingAction(k uint8, selfhand, desthand string) {
	var next uint8
	if k == 3 {
		next = 1
	} else {
		next = k + 1
	}

	if k == s && selfhand == x {
		s = next
		x = desthand
	} else if next == s && desthand == x {
		s = k
		x = selfhand
	}
}

var scanner = bufio.NewScanner(os.Stdin)
var (
	s uint8
	x string
)

func main() {
	var n uint8

	fmt.Scan(&n)

	fmt.Scanf("%d %s", &s, &x)

	for i := uint8(0); i < n; i++ {
		scanner.Scan()
		line := scanner.Text()

		line = strings.ReplaceAll(line, " ", "")

		a := string(line[0])

		k, _ := strconv.ParseUint(string(line[1]), 10, 8)
		switch a {
		case "1":
			emptyAction(uint8(k))
		case "2":
			movingAction(uint8(k), string(line[2]), string(line[3]))
		default:
			os.Exit(1)
		}
	}
	fmt.Printf("%v %v", s, x)
}
