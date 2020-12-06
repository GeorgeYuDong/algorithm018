/*

�����㷨��ʽ��
result = []
func backtrack(·����ѡ���б�) {
	if ����������� {
		result.add(·��)
	}
	return

	for ѡ�� in ѡ���б� {
		��ѡ��
		backtrack(·����ѡ���б�)
		����ѡ��
	}
}

*/

// ȫ�ֱ���,������
var result [][]string

// ���ݺ���
// board�� ����
// path��ѡ���б�
func backtrack(board [][]bool, path [][]byte) {
	// ��������
	if len(path) == len(board) {
		t := make([]string, len(path))
		for k, bs := range path {
			t[k] = string(bs)
		}
		result = append(result, t)
	}

	for i := 0; i < len(board); i++ {
		// ���Ϸ����,�޳�(��֦)
		if !isValid(board, len(path), i) {
			continue
		}
		// ��ӡĬ��λ��
		bs := printLine(len(board))
		// ���ûʺ�
		bs[i] = 'Q'
		// ��������
		board[len(path)][i] = true
		// ����ѡ��·��
		path = append(path, bs)
		// ������һ�ξ���
		backtrack(board, path)
		// ����ѡ��
		path = path[:len(path)-1]
		board[len(path)][i] = false
	}
}

// �Ƿ����� board[row][col] λ�÷��ûʺ�
// �ʺ󲻿����������ҶԽ���ͬʱ����
func isValid(board [][]bool, row, col int) bool {
	// ������Ƿ��лʺ��ͻ
	for i := 0; i < row; i++ {
		if board[i][col] == true {
			return false
		}
	}
	// ������Ƿ��лʺ��ͻ
	for j := 0; j < len(board); j++ {
		if board[row][j] == true {
			return false
		}
	}
	// ���Խ���: "\"
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == true {
			return false
		}
	}
	// ���Խ���: "/"
	for i, j := row, col; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == true {
			return false
		}
	}

	return true
}

// ��ӡÿ��Ĭ�����,����'.'
func printLine(n int) []byte {
	bs := make([]byte, n)
	for i := 0; i < n; i++ {
		bs[i] = '.'
	}
	return bs
}

func solveNQueens(n int) [][]string {
	// ��ձ���
	result = [][]string{}
	board := make([][]bool, n)
	for i := 0; i < n; i++ {
		board[i] = make([]bool, n)
	}
	backtrack(board, [][]byte{})
	return result
}

