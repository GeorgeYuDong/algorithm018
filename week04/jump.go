/*
����˼·
˼·��������ֻ��һ��������0
ÿһ�ζ�������õ�λ�ã�������λ������+λ���±� ��������ܱ�֤��Ծ�������١�
*/

func jump(nums []int) int {
    if len(nums)==1{
		return 0
	}
	length := len(nums)
	//˼·��������Ծ�ľ���
	for i,step:=0,0;;step++{
		//Ϊ0˵�����һ����������
		if nums[i] == 0{
			return step
		}
		//���һ��
		if i + nums[i] >= length-1{
			return step +1
		}
		//��һ��Ӧ���ߵ�����
		max,maxJ := 0,i+1
		for j:=i+1;j<=nums[i] + i;j++{
			if nums[j] + j > max{//�ҵ��ۺ�λ����õĵط�
				max = nums[j] + j
				//�����ĸ�max��iӦ��������λ��
				maxJ = j
			}
		}
		i = maxJ
	}
}

