/*
 *����˼·�����Trie���������������ڣ�������������������
 *
 *
 *
 *
 *
 */

class Solution {
class TrieNode{ //���Ľڵ���
        TrieNode [] next;
        String isWord;
        TrieNode(){
            this.isWord = "";
            this.next = new TrieNode [26];
        }
    }
    class Trie{//���࣬����2�������������в��뵥���Լ���鵱ǰ����Ƿ����ĳ���ַ����ӽڵ�
        private TrieNode root;
        Trie(){
            this.root = new TrieNode();
        }
        public void insert(String x){
            TrieNode curr = root;
            for(int i=0;i<x.length();i++){
                TrieNode node = curr.next[x.charAt(i)-'a'];
                if(node==null){
                    node = new TrieNode();
                    curr.next[x.charAt(i)-'a'] = node;
                }
                curr = node;
                if(i==x.length()-1)     curr.isWord = x;
            }
        }
        public boolean contains(char x,TrieNode root){
            if(root.next[x-'a']!=null)  return true;
            else return false;
        }
    }
    List<String> res = new ArrayList<>();
    Trie tree = new Trie();
    int []x = new int []{-1,1,0,0};
    int [] y = new int []{0,0,-1,1};
    public List<String> findWords(char[][] board, String[] words) {
        for(String word:words){//����words������
            tree.insert(word);
        }
        for (int i = 0; i < board.length; i++) {
            for (int j = 0; j < board[0].length; j++) {
                if(tree.contains(board[i][j],tree.root)){//������ĸ��ڵ���ӽڵ�����ַ�����ʼ�����Ƿ����������
                    dfs(board,i,j,tree.root);
                }
            }
        }
        return res;
    }
    private void dfs(char[][] board,int i, int j, TrieNode root){
        TrieNode node = root.next[board[i][j]-'a'];
        if(!node.isWord.equals("")){//������ʱ�ǲ�Ϊ�գ�����ӣ��������ÿ�
            res.add(node.isWord);
            node.isWord = "";
        }
        char ss = board[i][j];
        board[i][j] = '.';//��ֹ��η���
        for(int index=0;index<4;index++){
            int newi = i+x[index];
            int newj = j+y[index];
            if(newi<0||newj<0||newi>=board.length||newj>=board[0].length||board[newi][newj]=='.')    continue;
            if(tree.contains(board[newi][newj],node))   
				dfs(board,newi,newj,node);
        }
        board[i][j] = ss;//����
    }
}

