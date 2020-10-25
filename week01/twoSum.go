//时间复杂度N，空间复杂度N
//N为数组元素个数





func twoSum(nums []int, target int) []int {
	result := []int{};
	m := make(map[int]int)
	for i,k := range nums{
		if value,exist := m[target-k];exist{
			result = append(result, value)
			result = append(result, i)
		}
		m[k] = i;
	}
	return result
}
