
/*ʱ�临�Ӷ�O��N��,�ռ临�Ӷ�O��1��*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// if len(s) < 1 {  // �մ�������Ч�ʺ�
	// 	return true
	// }
    s_arr, t_arr := [26]int{}, [26]int{}
	for i := 0; i < len(s); i++ {
		// s[i]-'a'����ֵ�޶���26֮��
		s_arr[s[i]-'a']++
		t_arr[t[i]-'a']++
	}
    if s_arr!=t_arr{
        return false
    }
	return true
}

