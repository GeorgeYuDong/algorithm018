/*����һ�������������
����ʼ�ڵ������������������������ڵ��򽫽ڵ�����Ϊ0���������������
*/

func numIslands(grid [][]byte) int {
	// �������Ϊ0����ֱ�ӷ���
	if len(grid) == 0 {
		return 0
	}

	row, col := len(grid), len(grid[0])
	var count int
	// �������нڵ�
	for x := 0; x < row; x++ {
		for y := 0; y < col; y++ {
			// ����ڵ�Ϊ���죬���������
			if grid[x][y] == '1' {
				dfs(grid,x,y)
				count++
			}
		}
	}
	return count
}

func dfs(grid [][]byte, x, y int) {
	row, col := len(grid), len(grid[0])
	// �����������, �ݹ��˳�����, �ǳ���Ҫ��
	// �����������Ľڵ㣬�����µݹ�
	if x < 0 || x >= row || y < 0 || y >= col || grid[x][y] == '0' {
		return
	}
	// ����������Ϊ0
	grid[x][y] = 0
	// ���²�(����)���������������������Ľڵ�ͨ���������жϣ����ǵݹ鵽��һ���ж�
	dfs(grid, x+1, y)
	dfs(grid, x-1, y)
	dfs(grid, x, y-1)
	dfs(grid, x, y+1)
}

/*
�������������������
����ʼ�ڵ���й����������

*/
func numIslands(grid [][]byte) int {
    row,col := len(grid),len(grid[0])
    ans := 0
    if row == 0 {
        return ans
    }

    for i:=0;i<row;i++ {
        for j:=0;j<col;j++ {
            if grid[i][j] == '1' {
				//  dfs(grid,i,j)
                bfs(grid,i,j)
                ans++
            }
        }
    }
    return ans
}

func bfs(grid [][]byte,x,y int)  {
	row,col := len(grid),len(grid[0])
	if x<0||x>=row||y<0||y>=col || grid[x][y] == '0' {
		return
	}

	queue :=[][]int{[]int{x,y}}
	for len(queue) > 0 {
		tx,ty := queue[0][0],queue[0][1]
		queue = queue[1:]

		if tx > 0 && grid[tx-1][ty] == '1' {
			queue = append(queue,[]int{tx-1,ty})
			grid[tx-1][ty] = '0'
		}
		// �ж��Ƿ�Խ�磬���Խ���򲻼���������������
		if tx <row-1  && grid[tx+1][ty] == '1' {
			queue = append(queue,[]int{tx+1,ty})
			grid[tx+1][ty] = '0'
		}
		if ty > 0 && grid[tx][ty-1] == '1' {
			queue = append(queue,[]int{tx,ty-1})
			grid[tx][ty-1] = '0'
		}
		if ty < col-1 && grid[tx][ty+1] == '1' {
			queue = append(queue,[]int{tx,ty+1})
			grid[tx][ty+1] = '0'
		}
	}
}
