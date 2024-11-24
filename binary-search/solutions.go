package binary_search

// recursive + two pointer(?)
func BinarySearch(nums []int, target int, left int, right int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	if nums[mid] == target {
		return mid
	}

	if nums[mid] < target {
		return BinarySearch(nums, target, mid+1, right)
	} else {
		return BinarySearch(nums, target, left, mid-1)
	}
}
