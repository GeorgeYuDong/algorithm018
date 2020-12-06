func leftmove (v int) int {
    return 0x1 << v
}

func byte2int(b byte) int {
    return int(b-'0')
}

func int2byte(i int) byte {
    return byte('0' + i)
}

var row, col, box [9]int // 9�� 9�� 9��
var success bool // �������״̬���

// ��������
func SearchAndRollback(step int, board [][]byte) {
    // ����81���������
    if step >= 81 {
        success = true
        return
    }

    var i, j = step/9, step%9

    // �ǿո�ֱ�Ӳ�����һ��
    if board[i][j] != '.' {
        SearchAndRollback(step+1, board)
        return
    }

    // 1-9 ����һ������������
    var idx = i/3*3 + j/3
    for k := 1; k < 10; k++ {
        var v = leftmove(k)
        // ɸ��������������������������
        if row[i] & v > 0 || col[j] & v > 0 || box[idx] & v > 0 {
            continue
        }

        // ���
        board[i][j] = int2byte(k)
        row[i] |= v
        col[j] |= v
        box[idx] |= v

        SearchAndRollback(step+1, board)
        if success {
            break // ����������������ֹ
        }

        // ����������������
        row[i] ^= v
        col[j] ^= v
        box[idx] ^= v
        board[i][j] = '.'
    }
}

func solveSudoku(board [][]byte)  {
    row, col, box, success = [9]int{}, [9]int{}, [9]int{}, false // ״̬����

    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            // ��������Ѹ��������֣���������
            if board[i][j] != '.' {
                var idx = i/3*3 + j/3
                var v = leftmove(byte2int(board[i][j]))
                row[i] |= v
                col[j] |= v
                box[idx] |= v
            }
        }
    }

    // �ӵ�һ��ʼ����
    SearchAndRollback(0, board)
}

