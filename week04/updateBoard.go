/*
����˼·
��ʼ�ʼ��ʱ����������һ����������������Ͳȵ�***����ֱ�ӱ�Ǻ������Ϸ
���û�ȵ�����Ҫɨ����Χ��8��Ԫ�أ��������ֵ����Χû���򷵻�B����Χ���򷵻ض�Ӧ��Ԫ��
�������һ��������ʵ������Ҫ��ֻ֤�е�ǰԪ��ΪB��ʱ�򣬼���ΧΪ�յ�ʱ�򣬲ż���ɨ����Χ��Ԫ�أ������ǵ�ǰԪ��֮��Ͳ���׷�ӵ�ǰԪ�ص���ΧԪ��
*/

func updateBoard(board [][]byte, click []int) [][]byte {
    xLength := len(board)
    yLength := 0
    if xLength > 0 {
        yLength = len(board[0])
    }

    slice := [][]int{click}
    traceMap := map[string]bool{}
    for i:=0; i<len(slice); i++ {
        x, y := slice[i][0], slice[i][1]

        if board[x][y] == 'M' {
            board[x][y] = 'X'
            return board
        }

        i := x
        if x > 0 {
            i = x-1
        }

        count := 0
        for ; i<xLength&&i<=x+1; i++ {
            j := y
            if y > 0 {
                j = y - 1
            }
            for ; j<yLength&&j<=y+1; j++ {
                if board[i][j] == 'M' {
                    count++
                }
            }
        }

        if board[x][y] == 'E' {
            if count == 0 {
                board[x][y] = 'B'
                i := x
                if x > 0 {
                    i = x-1
                }
                for ; i<xLength&&i<=x+1; i++ {
                    j := y
                    if y > 0 {
                        j = y - 1
                    }
                    for ; j<yLength&&j<=y+1; j++ {
                        if board[i][j] == 'E' {
                            _, ok := traceMap[strconv.Itoa(i) + ":" + strconv.Itoa(j)]
                            if !ok {
                                slice = append(slice, []int{i, j})
                                traceMap[strconv.Itoa(i) + ":" + strconv.Itoa(j)] = true
                            }
                        }
                    }
                }


            } else {
                board[x][y] = []byte(strconv.Itoa(count))[0]
            }
        }
    }

    return board
}

