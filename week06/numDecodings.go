/**
���ݵ�ǰ����λ��i��������i-1�����ֲ�ͬ���������Ҫ��dp[i]������ͬ�Ľ⡣

1.���i��0��i-1����1Ҳ����2�������Ǹ��Ƿ����ַ�����(����ʧ��); ����dp[i] = dp[i-2]��ͬʱҲ�ø���dp[i-1] = dp[i-2]����Ϊ
2.���iС�ڵ���6��ǡ��i-1��1����2�������i����6��ǰ��������1����dp[i] = dp[i-1] + dp[i-2]
3.ʣ�µ����dp[i] = dp[i-1]

���ڿռ临�Ӷ�ΪO(1)���Ż�ʵ�ֶ��ԣ�ά��һ��cur������dp[i]����pre(����dp[i-1]) ����

*/
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	pre, cur, tmp := 1, 1, 0
	for i := 1; i < len(s); i++ {
		switch {
		case s[i] == '0' && s[i-1] != '1' && s[i-1] != '2':
			return 0
		case s[i] == '0':
			cur = pre
		case (s[i] <= '6' && (s[i-1] == '1' || s[i-1] == '2')) || (s[i] > '6' && s[i-1] == '1'):
			tmp = cur
			cur = cur + pre
			pre = tmp
		default:
			pre = cur
		}
	}
	return cur
}
