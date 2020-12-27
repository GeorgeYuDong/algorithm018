class Solution {
public:
    vector<int> findAnagrams(string s, string p) {
        unordered_map<char,int>need,window;
        for(char c:p)need[c]++;
        int left=0,right=0;//���ұգ�len=right-left
        int valid=0;
        vector<int>res;
        while(right<s.size()){
            char c=s[right];
            right++;//�ұ�
            if(need.count(c)!=0){//needû��ֱ���������еĻ���Ҫwindow[c]++ʵ�ָ��£���һ���º�window[c]==need[c]��Ҫvalid++
                window[c]++;
                if(window[c]==need[c])valid++;
            }
            while(valid==need.size()){//left��������
                if(right-left==p.size())res.push_back(left);//�������ڳ��ȵ���pʱ���Ǵ�
                char d=s[left];
                left++;
                if(need.count(d)!=0){//needû��ֱ���������еĻ���Ҫ�ǽ�window[d]--ʵ�ָ��£���һ����window[d]==need[d]��--֮��valid�����ˣ�left��������
                    if(window[d]==need[d])valid--;
                    window[d]--;
                }
            }
        }
        return res;
    }
