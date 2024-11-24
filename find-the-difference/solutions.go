package find_the_difference

import "fmt"

func findTheDifference(s string, t string) byte {
	sumS, sumT := 0, 0

	for i := len(t) - 1; i >= 0; i-- {
		sumT = sumT + int(t[i])

		if i != len(t) - 1 {
			sumS = sumS + int(s[i])
		}
	}

	return byte(sumT - sumS)
}

func main()  {
	fmt.Printf("%v\n", findTheDifference("abcd", "abcde"))
}