/*
����˼·��
�������ֻ��һ���� a�� ��ȫ���о��������� a��
����������� a, b�� ������ֿ��� ��a ׷�ӵ� b ��ȫ���� ���� [b, a]�� ��b׷�ӵ�a��ȫ���о��� [a, b]
�ã���ô������ a,b,c��
�� a �ó��� ׷�ӵ� b��c��ȫ���еõ� [c,b,a] �� [b,c,a]��
�ٰ�b�ó��� ׷�ӵ� a, c��ȫ���еõ� [c,a,b]��[a,c,b]��
�ٰ�c�ó��� ׷�ӵ� a, b��ȫ���еõ� [b,a,c]��[a,b,c]��
�������ơ�����

���Ե��ƹ�ʽ ����������ÿ���ó�һ��������׷�ӵ�ʣ�������ȫ������
�ݹ�ĳ����ǣ�����ֻ��һ����ʱ�����Լ�����
4����(a,b,c,d)
a,b,c,d ȫ����[b,c,d,a],[b,d,c,a],[c,b,d,a],[d,b,c,a],[c,d,b,a],[d,c,b,a]
			  [b,c,a,d],[b,d,a,c],[d,b,a,c],[d,c,a,b],[c,d,a,b],[c,b,a,d]

ʱ�临�Ӷ�O(N*N!)
�ռ临�Ӷ�O(N)
������⣬ʹ���˵ݹ飬ʱ�临�ӶȱȽϸ�
*/

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}

	res := [][]int{}

	for i, num := range nums {
		// ��num�� nums �ó�ȥ �õ�tmp
		tmp := make([]int, len(nums)-1)
		copy(tmp[0:], nums[0:i])
		copy(tmp[i:], nums[i+1:])

		// sub �ǰ�num �ó�ȥ��������ʣ�����ݵ�ȫ����
		sub := permute(tmp)
		for _, s := range sub {
			res = append(res, append(s, num))
		}
	}
	return res
}
