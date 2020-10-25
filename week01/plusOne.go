

func plusOne(digits []int) []int {
    var result []int
    addon := 0
    for i := len(digits) - 1;i >= 0; i-- {
        digits[i]+=addon
        addon = 0
        if i == len(digits) - 1 {
            digits[i]++
        }
        if digits[i] == 10 {
            addon = 1
            digits[i] = digits[i] % 10
        }
    }
    if addon == 1 {
        result = make([]int, 1)
        result[0] = 1
        result = append(result,digits...)
    }else{
        result = digits
    }
    return result
}

