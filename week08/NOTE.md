
学习笔记

排序代码：

#define MAX 9



//交换两个记录的位置
void swap(int *a, int *b){
    int temp;
    temp = *a;
    *a = *b;
    *b = temp;
}

//直接插入排序函数
void InsertSort(int a[], int n)
{
    for(int i= 1; i<n; i++) {

		//若第 i 个元素大于 i-1 元素则直接插入；反之，需要找到适当的插入位置后在插入。
        if(a[i] < a[i-1]) {

            int j= i-1;
            int x = a[i];

            while(j>-1 && x < a[j]){  

			//采用顺序查找方式找到插入的位置，在查找的同时，将数组中的元素进行后移操作，给插入元素腾出空间

                a[j+1] = a[j];
                j--;
            }

			//插入到正确位置
            a[j+1] = x;          
		}
    }
}

//折半插入排序
void BInsertSort(int a[],int size){

    int i,j,low = 0,high = 0,mid;
    int temp = 0;

    for (i=1; i<size; i++) {

        low=0;
        high=i-1;
        temp=a[i];

        //采用折半查找法判断插入位置，最终变量 low 表示插入位置
        while (low<=high) {

            mid=(low+high)/2;

            if (a[mid]>temp) {
                high=mid-1;
            } else {
                low=mid+1;
            }
        }

        //有序表中插入位置后的元素统一后移
        for (j=i; j>low; j--) {
            a[j]=a[j-1];
        }

		//插入元素
        a[low]=temp;
    }
}

//冒泡排序
void BubbleSort(int array[], int n)
{
	int i, j;

    int key;

    //有多少记录，就需要多少次冒泡，当比较过程，所有记录都按照升序排列时，排序结束
    for (i = 0; i < n; i++){

		//每次开始冒泡前，初始化 key 值为 0
        key=0;

        //每次起泡从下标为 0 开始，到 n-i 结束
        for (j = 0; j+1 < n-i; j++){

            if (array[j] > array[j+1]){
                key=1;
                swap(&array[j], &array[j+1]);
            }

        }

        //如果 key 值为 0，表明表中记录排序完成
        if (key==0) {
            break;
        }
    }
}



