/*
 * ʱ�临�Ӷ�O��1��������ʱ����n��λ1���йء������£�n������λ����1������32λ����������ʱ����O��1����
 * �ռ临�Ӷ�O��1��
 *
 */
public int hammingWeight(int n) {
    int sum = 0;
    while (n != 0) {
        sum++;
        n &= (n - 1);
    }
    return sum;
}
