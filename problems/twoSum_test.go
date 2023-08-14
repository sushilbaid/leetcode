package problems

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwosum(t *testing.T) {

	tests := []struct {
		nums, expectResult []int
		target             int
	}{
		{[]int{2, 7, 11, 15}, []int{0, 1}, 9},
		{[]int{3, 2, 4}, []int{1, 2}, 6},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test%d", i+1), func(t *testing.T) {
			a := assert.New(t)
			result := twoSum(test.nums, test.target)
			a.Equal(test.expectResult, result)
		})
	}
}
