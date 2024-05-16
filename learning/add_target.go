package learning

import "github.com/bubblepays/log"

//AddTwoIntTarget nums = [10, 1, 11, -1], target = 10
func AddTwoIntTarget(nums []int, target int) []int {
	numIndices := make(map[int]int)
	for i, num := range nums {
		// Calculate the complement needed to reach the target
		complement := target - num
		log.Infof("num  : %v, complement: %v", num, complement)
		// Check if the complement exists in the map
		if index, ok := numIndices[complement]; ok {
			// Found a pair that adds up to the target
			log.Infof("numIndices[complement] after: %v ", numIndices[complement])
			log.Infof("numIndices[complement] %v ", numIndices[complement])
			return []int{index, i}
		}
		numIndices[num] = i
	}
	return nil
}

func AddTwoIntegersArray(nums []int, target int) (int, int) {
	var i int
	for i, num := range nums {
		// Calculate the complement needed to reach the target
		complement := target - num
		log.Infof("num  : %v, complement: %v ", num, complement)
		if complement == nums[i] {

		}
	}
	return i, i + 1
}
