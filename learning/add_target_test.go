package learning

import (
	"testing"

	"github.com/bubblepays/log"
	"github.com/stretchr/testify/assert"
)

func TestAddTarget(t *testing.T) {
	nums := []int{10, 1, 11, -1}
	target := 10
	result := AddTwoIntTarget(nums, target)
	tt := "hi"
	assert.Equal(t, "hi", tt)
	assert.Equal(t, 2, result[0])
	assert.Equal(t, 3, result[1])

	nums = []int{10, 1, 11, -1}
	target = 10
	result = AddTwoIntTarget(nums, target)
	log.Info(result)
}

func TestAddTargetArray(t *testing.T) {
	nums := []int{2,3,4,8,6,5}
	target := 10
	result := AddTwoIntTarget(nums, target)
	tt := "hi"
	assert.Equal(t, "hi", tt)
	//assert.Equal(t, 2, result[0])
	//assert.Equal(t, 3, result[1])
	log.Infof(" result: %v ", result)

}
