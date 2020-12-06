/*
 * ʹ�ò��鼯
 * ʱ�临�Ӷ�O(n3)
 *
 *
 *
 */ 
public class Solution {
    public int findCircleNum(int[][] M) {
        Union un = new Union(M.length);
        for(int i=0;i<M.length;i++){
            for(int j=i+1;j<M.length;j++){
                if(M[i][j]==1)  un.union(i, j);
            }
        }
        return un.getCount();
    }
    class Union{
        private int [] group;//��������
        private int n;//��ʼ��С
        private int count;//ӵ�е�����
        Union(int n){
            this.n = n;
            this.group = new int [n];
            for (int i = 0; i < group.length; i++) {
                group[i] = i;//��ʼ��
            }
            this.count = n;
        }
        public int find(int p){
            //�ҵ����ĸ��ڵ�
            if(p==group[p]){
                return p;
            }
            return group[p] = find(group[p]);
        }
        public void union(int p,int q){//�ϲ�2����
            int root1 = find(p);
            int root2 = find(q);
            if(root1!=root2){
                group[root1] = root2;
                this.count--;
            }
        }
        public int getCount(){
            return count;
        }
    }
}
