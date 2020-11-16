/*
�������û�з�ת���� nums[left] <= nums[right]���� nums[left] ������Сֵ��ֱ�ӷ��ء�

������鷭ת����Ҫ�ҵ������еڶ����ֵĵ�һ��Ԫ�أ�

1.�� nums[left] <= nums[mid]��˵������ [left,mid] ��������������СԪ��һ������������������ֱ���ų�����ˣ��� left = mid+1���� [mid+1,right] ��������

2.����˵������ [left,mid] ������������СԪ��һ��������������ˣ��� right = mid���� [left,mid] ��������

3.[left,right] ��ʾ��ǰ���������䡣right ����ʱ�ᱻ��Ϊ mid ������ mid-1����Ϊ mid �޷����ų���

ʱ�临�Ӷ�O(logN),�ռ临�Ӷ�O��1)
*/
func findMin (nums []int) int {
    left, right := 0, len(nums)-1
    for left <= right { // ʵ�����ǲ�������ѭ������ left==right ʱֱ�ӷ���
        if nums[left] <= nums[right] { // ��� [left,right] ������ֱ�ӷ���
            return nums[left]
        }
        mid := left + (right-left)>>1
        if nums[left] <= nums[mid] { // [left,mid] �������������� [mid+1,right] ����
            left = mid + 1
        }else {
            right = mid // [left,mid] ���������� [left,mid] ����
        }
    }
    return -1
}
