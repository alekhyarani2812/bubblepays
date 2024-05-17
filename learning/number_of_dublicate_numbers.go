package learning

import (
	"fmt"
)

func DuplicateNumbers(numberArray []int) ([]int) {
	counts := make(map[int]int)
	duplecates := []int{}
	//count each number in the array
	for _, num := range numberArray {
		counts[num]++
		fmt.Printf("num : ,counts[num]: %v %v",num,counts[num])
	}
	for num, count := range counts {
		if count > 1 {
			duplecates = append(duplecates, num)
		}
	}
	return duplecates
}
