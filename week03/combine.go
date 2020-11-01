/*



*/

func combine(n int, k int) (ans [][]int) {
	// ��ʼ��
	// �� temp �� [0, k - 1] ÿ��λ�� i ����Ϊ i + 1���� [0, k - 1] �� [1, k]
	// ĩβ��һλ n + 1 ��Ϊ�ڱ�
	temp := []int{}
	for i := 1; i <= k; i++ {
		temp = append(temp, i)
	}
	temp = append(temp, n+1)

	for j := 0; j < k; {
		comb := make([]int, k)
		copy(comb, temp[:k])
		ans = append(ans, comb)
		// Ѱ�ҵ�һ�� temp[j] + 1 != temp[j + 1] ��λ�� t
		// ������Ҫ�� [0, t - 1] �����ڵ�ÿ��λ�����ó� [1, t]
		for j = 0; j < k && temp[j]+1 == temp[j+1]; j++ {
			temp[j] = j + 1
		}
		// j �ǵ�һ�� temp[j] + 1 != temp[j + 1] ��λ��
		temp[j]++
	}
	return
}

