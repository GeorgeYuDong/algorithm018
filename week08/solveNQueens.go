/*
基于位运算的回溯
时间复杂度:O(N!),其中N是皇后数量。

空间复杂度:O(N),其中N是皇后数量.由于使用位运算表示,因此存储皇后信息的空间复杂度是O(1),空间复杂度主要取决于递归调用层数和记录每行放置的皇后的列下标的数组,递归调用层数不会超过N,数组的长度为N
*/

var solutions [][]string

func solveNQueens(n int) [][]string {
    solutions = [][]string{}
    queens := make([]int, n)
    for i := 0; i < n; i++ {
        queens[i] = -1
    }
    solve(queens, n, 0, 0, 0, 0)
    return solutions
}

func solve(queens []int, n, row, columns, diagonals1, diagonals2 int) {
    if row == n {
        board := generateBoard(queens, n)
        solutions = append(solutions, board)
        return
    }
    availablePositions := ((1 << n) - 1) & (^(columns | diagonals1 | diagonals2))
    for availablePositions != 0 {
        position := availablePositions & (-availablePositions)
        availablePositions = availablePositions & (availablePositions - 1)
        column := bits.OnesCount(uint(position - 1))
        queens[row] = column
        solve(queens, n, row + 1, columns | position, (diagonals1 | position) >> 1, (diagonals2 | position) << 1)
        queens[row] = -1
    }
}

func generateBoard(queens []int, n int) []string {
    board := []string{}
    for i := 0; i < n; i++ {
        row := make([]byte, n)
        for j := 0; j < n; j++ {
            row[j] = '.'
        }
        row[queens[i]] = 'Q'
        board = append(board, string(row))
    }
    return board
}

