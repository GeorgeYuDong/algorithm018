//快慢指针解法
//时间复杂度N，空间复杂度1

func removeDuplicates(nums []int) int {
	n := len(nums)
    if n<2{
        return n
    }

    left, right := 1,1
    for right < n{
        if nums[right] != nums[right-1]{
            nums[left] = nums[right]
            left++
        }
        right++
    }
    return left
}
