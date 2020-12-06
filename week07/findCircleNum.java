/*
 * 使用并查集
 * 时间复杂度O(n3)
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
        private int [] group;//分组数组
        private int n;//初始大小
        private int count;//拥有的组数
        Union(int n){
            this.n = n;
            this.group = new int [n];
            for (int i = 0; i < group.length; i++) {
                group[i] = i;//初始化
            }
            this.count = n;
        }
        public int find(int p){
            //找到结点的父节点
            if(p==group[p]){
                return p;
            }
            return group[p] = find(group[p]);
        }
        public void union(int p,int q){//合并2个组
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
