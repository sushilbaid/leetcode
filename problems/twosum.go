package problems

import "sort"

type item struct {
	index, value int
}

func twoSum(nums []int, target int) []int {
	n := len(nums)
	a := make([]*item, n)
	for i, v := range nums {
		a[i] = &item{i, v}
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i].value < a[j].value
	})
	s, e := -1, -1
outer:
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			sum := a[i].value + a[j].value
			if sum > target {
				break
			} else if sum == target {
				s, e = a[i].index, a[j].index
				break outer
			}
		}
	}
	return []int{s, e}
}
