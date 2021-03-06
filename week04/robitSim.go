
/*
解题思路
1、规定移动方向为：上0右1下2左3，当前方向curDir初始方向为0，即向上
2、遍历指令数组commands
3、碰到转向命令-1/-2，改变相应的方向curDir
1）如果是-1，向右
curDir = (curDir+1) % 4
2）如果是-2，向左
curDir = (curDir+3) % 4

可以试着算一算，比如-1时候， 1->2 / 2->3 / 3->4 / 4-> 1

4、如果碰到前进命令，就每前进一步，都要判断下，前进后的位置有没有障碍物，有的话，停止前进
5、每一次前进命令结束后，因为发生了位移，都要计算下最新的欧式距离有没有超过最大值
*/

func robotSim(commands []int, obstacles [][]int) int {
	// 方向：上右下左0123
	result, curDir, mObstacles := float64(0),0, make(map[string]bool) // 结果 当前方向 障碍物哈希
	x, y, stepX, stepY := 0, 0, []int{0,1,0,-1}, []int{1,0,-1,0}      // 当前的位置 以及 xy轴上各个方向移动的大小
	for _, v := range obstacles {
		mObstacles[fmt.Sprintf("%d-%d", v[0], v[1])] = true
	}
	for _, v := range commands {
		switch v {
		case -1:
			curDir = (curDir+1) % 4
		case -2:
			curDir = (curDir+3) % 4
		default:
			for i := 0; i < v;i++ {
				tempX, tempY := x + stepX[curDir], y + stepY[curDir]
				// 碰到障碍物，就不要移动了
				if mObstacles[fmt.Sprintf("%d-%d", tempX, tempY)] {
					break
				}
				x, y = tempX, tempY
				result = math.Max(float64(x*x+ y*y), result)
			}
		}
	}
	return int(result)
}
