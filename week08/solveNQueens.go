/*
����λ����Ļ���
ʱ�临�Ӷ�:O(N!),����N�ǻʺ�������

�ռ临�Ӷ�:O(N),����N�ǻʺ�����.����ʹ��λ�����ʾ,��˴洢�ʺ���Ϣ�Ŀռ临�Ӷ���O(1),�ռ临�Ӷ���Ҫȡ���ڵݹ���ò����ͼ�¼ÿ�з��õĻʺ�����±������,�ݹ���ò������ᳬ��N,����ĳ���ΪN
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

