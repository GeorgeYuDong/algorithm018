/*
如果数组没有翻转，即 nums[left] <= nums[right]，则 nums[left] 就是最小值，直接返回。

如果数组翻转，需要找到数组中第二部分的第一个元素：

1.若 nums[left] <= nums[mid]，说明区间 [left,mid] 连续递增，则最小元素一定不在这个区间里，可以直接排除。因此，令 left = mid+1，在 [mid+1,right] 继续查找

2.否则，说明区间 [left,mid] 不连续，则最小元素一定在这个区间里。因此，令 right = mid，在 [left,mid] 继续查找

3.[left,right] 表示当前搜索的区间。right 更新时会被设为 mid 而不是 mid-1，因为 mid 无法被排除。

时间复杂度O(logN),空间复杂度O（1)
*/
func findMin (nums []int) int {
    left, right := 0, len(nums)-1
    for left <= right { // 实际上是不会跳出循环，当 left==right 时直接返回
        if nums[left] <= nums[right] { // 如果 [left,right] 递增，直接返回
            return nums[left]
        }
        mid := left + (right-left)>>1
        if nums[left] <= nums[mid] { // [left,mid] 连续递增，则在 [mid+1,right] 查找
            left = mid + 1
        }else {
            right = mid // [left,mid] 不连续，在 [left,mid] 查找
        }
    }
    return -1
}
