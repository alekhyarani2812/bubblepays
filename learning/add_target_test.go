package learning

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

)
func TestAddTarget(t *testing.T) {
	nums := []int{10, 1, 11, -1}
	target := 10
	result := AddTwoIntTarget(nums, target)
	//fmt.Println(result)
	tt := "hi"
	assert.Equal(t, "hi", tt)
	assert.Equal(t, 2, result[0])
	assert.Equal(t, 3, result[1])

	nums = []int{10, 1, 11, 0}
	target = 10
	result = AddTwoIntTarget(nums, target)
	fmt.Println(result)
}
