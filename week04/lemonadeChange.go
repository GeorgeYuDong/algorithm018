// 算法1：
// 如果给5元，不用找
// 如果给10元，找5元，否则找不开
// 如果给20元，优先找10+5，再找5+5+5，否则找不开
// 时间复杂度O(N),空间复杂度O(1)
func lemonadeChange1(bills []int) bool {
	five, ten := 0, 0

	for _, v := range bills {
		if v == 5 {
			five ++
		} else if v == 10 {
			if five >= 1 {
				five --
				ten ++
			} else {
				return false
			}
		} else if v == 20 {
			if five >= 1 && ten >= 1 {
				five --
				ten --
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}

	return true
}
