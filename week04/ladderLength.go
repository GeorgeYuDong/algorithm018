/**
˼·:
�����ʳ�����ÿ�α䶯һ����ĸ������ n �α任������յ�ʣ�ϣ�� n ����С��
������Ҫ�ҳ��ڽӹ�ϵ������hit����һ����ĸ����_it��h_t��hi_��ʽ�ĵ��ʡ�
Ȼ������´��Ƿ�����ڵ��ʱ�������ڣ����ҵ���һ����һ���ת���ʡ�
ͬʱ��Ҫ�����ظ����ʣ�����hit->hot->hit�����������ֻ��ͽ��ת���ĳ��ȡ�
��ˣ�Ҫ����һ��ת���ʴӵ��ʱ���ɾ�������ʱ�ĵ�����Ψһ�ģ���
������һ��ĵ����ж������Ҫmaintain�������죬��һ��ת��·���������յ�ʣ�������̡�

����һ��:
��Ϊ˼·����һ����������һ����ڽӵ㣬��BFS���ѵ��ʿ����ڵ㡣
ά��һ�����У����������У�level Ϊ 1��Ȼ����п��졣
������ÿ���ַ����26��ĸ֮һ��������Ƿ��ڵ��ʱ�����ڣ�����´�Ϊ��һ���ת��ʡ�
�������У����� level +1�����ӵ��ʱ���ɾȥ����ʡ�
�������С��ظ��������еĵ��ʺ��յ����ͬ��˵���������յ�ʣ��������� level��
������Ϊ��ʱ��BFS������ʼ��û�������յ�ʣ�û��·��ͨ���յ㣬���� 0��
*/
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap := map[string]bool{}
	for _,w := range wordList{
		wordMap[w] = true
	}
	queue := []string{}
	queue = append(queue, beginWord)
	level := 1
	for len(queue) != 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			word := queue[0]
			queue = queue[1:]
			if word == endWord {
				return level
			}
			for c := 0; c < len(word); c++ {
				for j := 'a'; j <= 'z'; j++ {
					newWord := word[:c] + string(j) + word[c+1:]
					if wordMap[newWord] == true {
						queue = append(queue, newWord)
						wordMap[newWord] = false
					}
				}
			}
		}
		level++
	}
	return 0
}
