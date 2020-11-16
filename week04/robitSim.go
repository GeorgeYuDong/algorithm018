
/*
����˼·
1���涨�ƶ�����Ϊ����0��1��2��3����ǰ����curDir��ʼ����Ϊ0��������
2������ָ������commands
3������ת������-1/-2���ı���Ӧ�ķ���curDir
1�������-1������
curDir = (curDir+1) % 4
2�������-2������
curDir = (curDir+3) % 4

����������һ�㣬����-1ʱ�� 1->2 / 2->3 / 3->4 / 4-> 1

4���������ǰ�������ÿǰ��һ������Ҫ�ж��£�ǰ�����λ����û���ϰ���еĻ���ֹͣǰ��
5��ÿһ��ǰ�������������Ϊ������λ�ƣ���Ҫ���������µ�ŷʽ������û�г������ֵ
*/

func robotSim(commands []int, obstacles [][]int) int {
	// ������������0123
	result, curDir, mObstacles := float64(0),0, make(map[string]bool) // ��� ��ǰ���� �ϰ����ϣ
	x, y, stepX, stepY := 0, 0, []int{0,1,0,-1}, []int{1,0,-1,0}      // ��ǰ��λ�� �Լ� xy���ϸ��������ƶ��Ĵ�С
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
				// �����ϰ���Ͳ�Ҫ�ƶ���
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
