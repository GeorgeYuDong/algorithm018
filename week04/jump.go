/*
解题思路
思路：若数组只有一个，返回0
每一次都跳到最好的位置（即：该位置数字+位置下标 最大）这样能保证跳跃次数最少。
*/

func jump(nums []int) int {
    if len(nums)==1{
		return 0
	}
	length := len(nums)
	//思路：最大可跳跃的距离
	for i,step:=0,0;;step++{
		//为0说明最后一步不用走了
		if nums[i] == 0{
			return step
		}
		//最后一步
		if i + nums[i] >= length-1{
			return step +1
		}
		//下一步应该走到走哪
		max,maxJ := 0,i+1
		for j:=i+1;j<=nums[i] + i;j++{
			if nums[j] + j > max{//找到综合位置最好的地方
				max = nums[j] + j
				//最后的哪个max即i应该跳到的位置
				maxJ = j
			}
		}
		i = maxJ
	}
}

