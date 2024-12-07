package main

import "fmt"

func main()  {
	var n uint8
	
	fmt.Scan(&n)

	if n % 2 == 0 {
		fmt.Println("b")
	} else {
		fmt.Println("a")
	}
}