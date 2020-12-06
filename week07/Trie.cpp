/*
 * 
 *�ܽ�Trie �ļ������ʣ�

1.Trie ����״�͵��ʵĲ����ɾ��˳���޹أ�Ҳ����˵�������������һ�鵥�ʣ�Trie ����״����Ψһ�ġ�

2.���һ����һ������Ϊ L �ĵ��ʣ����� next ����Ĵ������Ϊ L+1���� Trie �а������ٸ������޹ء�

3.Trie ��ÿ������ж�������һ����ĸ�����Ǻܺķѿռ�ġ���� Trie �ĸ߶�Ϊ n����ĸ��Ĵ�СΪ m���������� Trie �л�������ǰ׺��ͬ�ĵ��ʣ��ǿռ临�ӶȾ�Ϊ O(m^n) 
 * 
 * Ӧ�ó�����һ�ν�������β�ѯ
 *
 *
 */

class Trie {
private:
    bool isEnd;
    Trie* next[26];
public:
    Trie() {
        isEnd = false;
        memset(next, 0, sizeof(next));
    }
    
    void insert(string word) {
        Trie* node = this;
        for (char c : word) {
            if (node->next[c-'a'] == NULL) {
                node->next[c-'a'] = new Trie();
            }
            node = node->next[c-'a'];
        }
        node->isEnd = true;
    }
    
    bool search(string word) {
        Trie* node = this;
        for (char c : word) {
            node = node->next[c - 'a'];
            if (node == NULL) {
                return false;
            }
        }
        return node->isEnd;
    }
    
    bool startsWith(string prefix) {
        Trie* node = this;
        for (char c : prefix) {
            node = node->next[c-'a'];
            if (node == NULL) {
                return false;
            }
        }
        return true;
    }
};


