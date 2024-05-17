package learning

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDuplicateNumbers(t *testing.T) {
	numbers := []int{5, 3, 7, 5, 3, 3, 4, 3, 2, 5, 2}
	nums := DuplicateNumbers(numbers)
	sort.Ints(nums)
	assert.Equal(t, []int{2,3,5}, nums, 
		fmt.Sprintln("Expcted an array of dublicate numbers. But got empty array"))
}
func TestDuplicateNumbersArray(t *testing.T) {
	//numbers := []int{3, 6, 1, 3, 3, 4, 3, 3}
	numbers := []int{3, 6, 1, 3}
	nums := DuplicateNumbers(numbers)
	assert.Equal(t, []int{3}, nums, 
		fmt.Sprintln("Expcted an array of dublicate numbers. But got empty array"))
}
func TestDuplicateNumbers_WithoutDuplecate(t *testing.T) {
	numbers := []int{2, 3, 1}
	nums := DuplicateNumbers(numbers)
	assert.Equal(t, []int{}, nums, fmt.Sprintf("Expcted an empty array. But got numbers : '%v'", nums))
}
func TestDuplicateNumbers_SingleNumber(t *testing.T) {
	numbers := []int{4}
	nums := DuplicateNumbers(numbers)
	assert.Equal(t, []int{}, nums, fmt.Sprintf("Expcted an empty array. But got numbers : '%v'", nums))
}

