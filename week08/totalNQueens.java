/*
 * 时间复杂度:O(N!)，其中N是皇后数量。
空间复杂度:O(N),其中N是皇后数量.由于使用位运算表示,因此存储皇后信息的空间复杂度是O(1),空间复杂度主要取决于递归调用层数,递归调用层数不会超过N。
 *
 */

class Solution {
    public int totalNQueens(int n) {
        return solve(n, 0, 0, 0, 0);
    }

    public int solve(int n, int row, int columns, int diagonals1, int diagonals2) {
        if (row == n) {
            return 1;
        } else {
            int count = 0;
            int availablePositions = ((1 << n) - 1) & (~(columns | diagonals1 | diagonals2));
            while (availablePositions != 0) {
                int position = availablePositions & (-availablePositions);
                availablePositions = availablePositions & (availablePositions - 1);
                count += solve(n, row + 1, columns | position, (diagonals1 | position) << 1, (diagonals2 | position) >> 1);
            }
            return count;
        }
    }
}
