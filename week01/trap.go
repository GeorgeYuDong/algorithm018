func trap(height []int) int {
	var left, right, leftMax, rightMax, res int
	right = len(height) - 1
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				//��������������
				leftMax = height[left]
			} else {
				//�ұ߱ض������ӵ�ˮ�����ԣ���������ֵС�ڵ���leftMax�ģ�ȫ������ˮ��
				res += leftMax - height[left]
			}
			left++
		} else {
			if height[right] > rightMax { 
				//�����ұ��������
				rightMax = height[right] 
			} else {
				//��߱ض������ӵ�ˮ�����ԣ���������ֵС�ڵ���rightMax�ģ�ȫ������ˮ��
				res += rightMax - height[right] 
			}
			right--
		}
	}
	return res
}
