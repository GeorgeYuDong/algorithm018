
/*时间复杂度O（N）,空间复杂度O（1）*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// if len(s) < 1 {  // 空串多的情况效率好
	// 	return true
	// }
    s_arr, t_arr := [26]int{}, [26]int{}
	for i := 0; i < len(s); i++ {
		// s[i]-'a'将数值限定在26之内
		s_arr[s[i]-'a']++
		t_arr[t[i]-'a']++
	}
    if s_arr!=t_arr{
        return false
    }
	return true
}

