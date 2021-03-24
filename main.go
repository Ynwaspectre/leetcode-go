package main

import (
	"fmt"
	"sort"
)

func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	leftMin := nums[0]
	for j := 1; j < len(nums)-1; j++ {
		if nums[j] < leftMin {
			leftMin = nums[j]
		} else {
			//在j+1 到len(nums)-1中找比leftMin大的最小值
			rightSort := make([]int, len(nums)-j-1)
			copy(rightSort, nums[j+1:])
			sort.Ints(rightSort)
			//二分查找比leftMin大的最小值
			rightIndex := find(leftMin, rightSort)
			if rightIndex == -1 {
				continue
			}
			if rightSort[rightIndex] < nums[j] {
				return true
			}
		}
	}
	return false
}

func find(leftMin int, nums []int) int {
	l := 0
	r := len(nums) - 1
	if r == 0 {
		if nums[0] > leftMin {
			return 0
		}
	}
	for l < r {
		mid := (l + r + 1) / 2
		if nums[mid] > leftMin {
			if nums[mid-1] == leftMin {
				return mid
			}
			r = mid - 1
		} else if nums[mid] < leftMin {
			if nums[mid+1] == leftMin {
				return mid + 2
			}
			l = mid + 1
		} else {
			if mid+1 >= len(nums) {
				return -1
			} else {
				return mid + 1
			}
		}
	}
	return -1
}

func main() {

	//nums := []int{-2, 1, 2, -2, 1, 2}
	nums := []int{3, 1, 4, 2}
	fmt.Print(find132pattern(nums))

}
