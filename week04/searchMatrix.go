/*
����˼·
����˼·�Ƚ��ر������ڶ��ֲ��ҡ����Թ涨����ʼ��Ϊ���»��������㣬��������ͬ�����߼ȿ��������ֿ��Լ���

�㷨���̣�

�涨����Ϊ��ʼ��
����(i,j)����target��ֱ�ӷ���true
����(i,j)С��target�������ߣ���j++
����(i,j)����target�������ߣ���i--
���Ѿ������˱߽磬����ѭ����˵��δ�ҵ�
*/
func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix)==0{
        return false
    }
    row:=len(matrix)
    col:=len(matrix[0])
    //1. �涨��ʼ��
    i,j:=row-1,0
    for i>=0 && j<col{
        if matrix[i][j]==target{ //2.�ҵ��򷵻�
            return true
        }else if matrix[i][j]<target{ //3.С��target�����Ҳ���
            j++
        }else{                      //4.����target�����ϲ���
            i--
        }
    }

    return false
}

