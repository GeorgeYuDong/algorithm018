/*
ʱ�临�Ӷ�O��NlogN��,NΪ���鳤��
�ռ临�Ӷ�O��N����NΪ���鳤��

*/


func reversePairs(nums []int) int {
    n := len(nums)
    if n <= 1 {
        return 0
    }

    n1 := append([]int(nil), nums[:n/2]...)
    n2 := append([]int(nil), nums[n/2:]...)
    cnt := reversePairs(n1) + reversePairs(n2) // �ݹ���Ϻ�n1 �� n2 ��Ϊ����

    // ͳ����Ҫ��ת�� (i,j) ������
    // ���� n1 �� n2 ��Ϊ���򣬿���������ָ��ͬʱ����
    j := 0
    for _, v := range n1 {
        for j < len(n2) && v > 2*n2[j] {
            j++
        }
        cnt += j
    }

    // n1 �� n2 �鲢���� nums
    p1, p2 := 0, 0
    for i := range nums {
        if p1 < len(n1) && (p2 == len(n2) || n1[p1] <= n2[p2]) {
            nums[i] = n1[p1]
            p1++
        } else {
            nums[i] = n2[p2]
            p2++
        }
    }
    return cnt
}

