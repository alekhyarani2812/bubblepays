package learning

import (
	"testing"

	"github.com/bubblepays/log"
	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate_True(t *testing.T) {
	nums := []int{22, 5, 7, 9, 4, 6, 22, 65}
	result := containsDuplicate(nums)
	assert.True(t, result, "Expected true, but got false")
}

func TestContainsDuplicate_False(t *testing.T) {
	nums := []int{22, 5, 7, 9, 4, 6, 65}
	result := containsDuplicate(nums)
	assert.False(t, result, "Expected false, but got true")
}

func TestContainsDuplicate(t *testing.T) {
	// Test case 1
	t.Run("Duplicate numbers", func(t *testing.T) {
		nums := []int{22, 5, 7, 9, 4, 6, 22, 65}
		result := containsDuplicate(nums)
		assert.True(t, result, "Expected true, but got false")
	})

	// Test case 2
	t.Run("Unique numbers", func(t *testing.T) {
		nums := []int{22, 5, 7, 9, 4, 6, 65}
		result := containsDuplicate(nums)
		assert.False(t, result, "Expected false, but got true")
	})
}

// Approach 2
func TestContainsDuplicateApproach2(t *testing.T) {

	tests := []struct {
		name             string
		expected         bool
		inputData        []int
		testErrorMessage string
	}{
		{
			name:             "Duplicate numbers",
			expected:         true,
			inputData:        []int{22, 5, 7, 9, 4, 6, 22, 65},
			testErrorMessage: "Expected true, but got false",
		},
		{
			name:             "Unique numbers",
			expected:         false,
			inputData:        []int{22, 5, 7, 9, 4, 6, 65},
			testErrorMessage: "Expected false, but got true",
		},
	}
	//Run Test case
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := containsDuplicate(test.inputData)
			if test.expected != result {
				t.Errorf(test.testErrorMessage)
			}
		})
	}
}

func TestClimbStairs(t *testing.T) {
	n := 4
	result := ClimbStairs(n)
	//result := RecClimbStairs(n)
	log.Infof("result : %d is number of ways to climb to the top ", result)
}

func TestProducerConsumer(t *testing.T) {
	ProducerConsumerChGoRoutines()
}

func TestShapeImpls(t *testing.T) {
	ShapeImpls()
}

func TestChanneleLinks(t *testing.T) {
	//ChanneleLinks()
	//ChanneleLinksGo()
	//ChanneleLinksGor()
	//ChanneleLinksGoRepeat()
	StrLteral()
}
