/*
解题思路
开始最开始的时候就是满足第一个条件，如果上来就踩到***，就直接标记后结束游戏
如果没踩到，就要扫描周围的8个元素，给出结果值，周围没有则返回B，周围有则返回对应的元素
对于最后一个条件，实际上是要保证只有当前元素为B的时候，即周围为空的时候，才继续扫描周围的元素，否则标记当前元素之后就不再追加当前元素的周围元素
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

