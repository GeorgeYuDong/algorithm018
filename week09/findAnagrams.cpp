class Solution {
public:
    vector<int> findAnagrams(string s, string p) {
        unordered_map<char,int>need,window;
        for(char c:p)need[c]++;
        int left=0,right=0;//左开右闭，len=right-left
        int valid=0;
        vector<int>res;
        while(right<s.size()){
            char c=s[right];
            right++;//右闭
            if(need.count(c)!=0){//need没有直接跳过，有的话需要window[c]++实现更新，万一更新后window[c]==need[c]还要valid++
                window[c]++;
                if(window[c]==need[c])valid++;
            }
            while(valid==need.size()){//left收缩条件
                if(right-left==p.size())res.push_back(left);//滑动窗口长度等于p时就是答案
                char d=s[left];
                left++;
                if(need.count(d)!=0){//need没有直接跳过，有的话主要是将window[d]--实现更新，万一本来window[d]==need[d]，--之后valid就少了，left收缩结束
                    if(window[d]==need[d])valid--;
                    window[d]--;
                }
            }
        }
        return res;
    }
