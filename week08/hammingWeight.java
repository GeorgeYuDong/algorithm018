/*
 * 时间复杂度O（1）。运行时间与n中位1的有关。最坏情况下，n中所有位都是1，对于32位整数，运行时间是O（1）。
 * 空间复杂度O（1）
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
