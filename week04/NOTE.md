学习笔记

寻找一个半有序数组中间无序的地方:
边界的判断条件变为根据nums[0] < nums[mid]来决定左侧还是右侧有序。

func find(nums []int) int {
	
	if nums[0] < nums[len(nums) - 1] {
		return -1		
	}

	l, h := 0, len(nums) - 1

	for l < h {
		mid = l + (h - l) / 2			
		if nums[mid] > nums[mid + 1] {
			return mid + 1	
		}
		if nums[mid - 1] > nums[mid] {
			return mid		
		}
		if nums[0] < nums[mid] {
			l = mid + 1	
		} else{
			h = mid - 1 	
		}
	}
	return -1
}

