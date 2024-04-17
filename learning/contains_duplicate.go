package learning

import "fmt"

func containsDuplicate(nums []int) bool {
	intMap := make(map[int]int)
	for _,n := range nums {
		_, exists := intMap[n]
		if exists {
			fmt.Println(n)
			return true
		}
		intMap[n]=n
	}
	return false
}
