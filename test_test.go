package Replace_Element_With_Biggest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	var nums = []int{17, 18, 5, 4, 6, 1}         // Input array
	var expectedNums = []int{18, 6, 6, 6, 1, -1} // The expected answer with correct length.
	var result = replaceElements(nums)           // Calls your implementation

	assert.Equal(t, result, expectedNums)
}
func replaceElements(nums []int) []int {
	if len(nums) < 2 {
		return []int{-1}
	}

	for i := 0; i < len(nums)-1; i++ {
		maximum := nums[i+1]
		for j := i + 1; j < len(nums); j++ {
			if maximum < nums[j] {
				maximum = nums[j]
			}
		}
		nums[i] = maximum
	}
	nums[len(nums)-1] = -1
	return nums
}
