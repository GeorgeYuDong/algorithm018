/*
解题思路：
如果数组只有一个数 a， 那全排列就是它本身 a，
如果有两个数 a, b。 把它拆分开， 把a 追加到 b 的全排列 就是 [b, a]。 把b追加到a的全排列就是 [a, b]
好，那么三个数 a,b,c。
把 a 拿出来 追加到 b，c的全排列得到 [c,b,a] 和 [b,c,a]。
再把b拿出来 追加到 a, c的全排列得到 [c,a,b]和[a,c,b]。
再把c拿出来 追加到 a, b的全排列得到 [b,a,c]和[a,b,c]。
依次类推。。。

所以递推公式 就是数组中每次拿出一个，把它追加到剩余数组的全排列中
递归的出口是，数组只有一个数时，把自己返回
4个数(a,b,c,d)
a,b,c,d 全排列[b,c,d,a],[b,d,c,a],[c,b,d,a],[d,b,c,a],[c,d,b,a],[d,c,b,a]
			  [b,c,a,d],[b,d,a,c],[d,b,a,c],[d,c,a,b],[c,d,a,b],[c,b,a,d]

时间复杂度O(N*N!)
空间复杂度O(N)
容易理解，使用了递归，时间复杂度比较高
*/

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}

	res := [][]int{}

	for i, num := range nums {
		// 把num从 nums 拿出去 得到tmp
		tmp := make([]int, len(nums)-1)
		copy(tmp[0:], nums[0:i])
		copy(tmp[i:], nums[i+1:])

		// sub 是把num 拿出去后，数组中剩余数据的全排列
		sub := permute(tmp)
		for _, s := range sub {
			res = append(res, append(s, num))
		}
	}
	return res
}
