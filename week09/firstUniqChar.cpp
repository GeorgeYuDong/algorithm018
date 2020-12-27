/*
两次遍历，第一次统计每个字符出现次数，第二次遍历只出现一次的字符
时间复杂度：O（N）
*/
class Solution {
public:
    int firstUniqChar(string s) {
        unordered_map<int, int> frequency;
        for (char ch: s) {
            ++frequency[ch];
        }
        for (int i = 0; i < s.size(); ++i) {
            if (frequency[s[i]] == 1) {
                return i;
            }
        }
        return -1;
    }
};
