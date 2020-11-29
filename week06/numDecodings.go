/**
根据当前迭代位置i的数字与i-1的数字不同的情况，需要对dp[i]作出不同的解。

1.如果i是0，i-1不是1也不是2，那这是个非法的字符串啊(解码失败); 否则dp[i] = dp[i-2]，同时也得更新dp[i-1] = dp[i-2]，因为
2.如果i小于等于6，恰好i-1是1或者2，再如果i大于6，前面数字是1，那dp[i] = dp[i-1] + dp[i-2]
3.剩下的情况dp[i] = dp[i-1]

对于空间复杂度为O(1)的优化实现而言，维护一个cur（代表dp[i]）和pre(代表dp[i-1]) 就行

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
