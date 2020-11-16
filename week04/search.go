/*
����˼·
�ⷨһ����������
�ο��ٷ��Ľ���˼·������ϸ���
˼·��������
1.�ö��ַ��ҵ�pivot(��),����pivot����С����ߵ������ұߵ���
2.ȷ��target��pivot����һ�ߣ����[0,pivot-1]�������ö��ֲ��ң��ұ�[pivot,len(nums)-1]�������ö��ֲ���
���Ա�����ѵ�˼·Ҳ�����㣺
a.�����ҵ�pivot
b.�����ҵ�target��pivot����һ�ߣ�����������׽����ֻ���ж�nums[0]��target�Ĺ�ϵ���ɡ�nums[0]��target��pivot�ұ�
nums[0]С��target��pivot��ߣ����԰�����Сֵ�Ķ���ȥ˼����
�����ص�����ö��ַ���pivot�ķ�ʽ��
(1)�����������飬��ʱһ����nums[left]<nums[right],��ʱpivot=0
(2)���ڷ��������飬��nums=[]
*/

func search(nums []int, target int) int {
    if len(nums)==0{
        return -1
    }
    pivot:=pivotIndex(nums)
    left,right:=0,0

    if nums[0]>target || pivot==0{
        left,right=pivot,len(nums)-1
    }else{
        left,right=0,pivot-1
    }

    for left<=right{
        mid:=(left+right)/2
        if nums[mid]==target{
            return mid
        }else if nums[mid]>target{
            right=mid-1
        }else{
            left=mid+1
        }
    }

    return -1
}

func pivotIndex(nums []int) int{
    left:=0
    right:=len(nums)-1
    if nums[left]<=nums[right]{
        return 0
    }
    for left<=right{
        mid:=(left+right)/2
        if nums[mid+1]>nums[mid]{
            if nums[left]<nums[mid]{
                left=mid+1
            }else{
                right=mid
            }
        }else{
            return mid+1
        }
    }
    return 0
}

/*
�ⷨ�������ݵ������䣬����target����
����һ�����������ͨ�ýⷨ��
�㷨˼·��

1.��mid��left����ֵ�ж�mid�Ǳ��ǵ�����������
2.���ڵ������䣬������˵��ֵ�������ж�target�Ƿ�϶�(ע���������)�ڴ������ڣ�Ȼ�����left��right������˵�����෴�������ڡ�
*/
func search(nums []int, target int) int {
    if len(nums)<=0{
        return -1
    }

    left:=0
    right:=len(nums)-1
    mid:=0

    for left<=right{
        mid=(left+right)/2
        if nums[mid]==target{
            return mid
        }
        //1 ��������ǵ�������
        if nums[mid]>=nums[left]{
            if nums[mid]>target && nums[left]<=target{ //1.1 [left,mid]�������,ͨ���˵���Կ϶�target�Ƿ��ڴ�����
                right=mid-1
            }else{                                     //1.2 ���ڣ�˵����[mid,right]����
                left=mid+1
            }
        }else{
            if nums[mid]<target && nums[right]>=target{
                left=mid+1
            }else{
                right=mid-1
            }
        }
    }

    return -1
}

