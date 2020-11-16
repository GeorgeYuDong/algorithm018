/*方法一：深度优先搜索
对起始节点进行深度优先搜索，搜索到节点则将节点设置为0，并继续深度搜索
*/

func numIslands(grid [][]byte) int {
	// 如果岛屿为0，则直接返回
	if len(grid) == 0 {
		return 0
	}

	row, col := len(grid), len(grid[0])
	var count int
	// 遍历所有节点
	for x := 0; x < row; x++ {
		for y := 0; y < col; y++ {
			// 如果节点为岛屿，则深度搜索
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
	// 深度优先搜索, 递归退出条件, 非常重要。
	// 不符合条件的节点，不向下递归
	if x < 0 || x >= row || y < 0 || y >= col || grid[x][y] == '0' {
		return
	}
	// 将自身设置为0
	grid[x][y] = 0
	// 向下层(四周)继续搜索，不符合条件的节点通常不进行判断，而是递归到下一层判断
	dfs(grid, x+1, y)
	dfs(grid, x-1, y)
	dfs(grid, x, y-1)
	dfs(grid, x, y+1)
}

/*
方法二：广度优先搜索
对起始节点进行广度优先搜索

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
		// 判断是否越界，如果越界则不继续进行向下搜索
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
