/**
思路:
从起点词出发，每次变动一个字母，经过 n 次变换，变成终点词，希望 n 尽量小。
我们需要找出邻接关系，比如hit它变一个字母会变成_it、h_t、hi_形式的单词。
然后看这个新词是否存在于单词表，如果存在，就找到了一个下一层的转换词。
同时，要避免重复访问，比如hit->hot->hit这样变回来，只会徒增转换的长度。
因此，要将下一个转换词从单词表中删除（单词表的单词是唯一的）。
可能下一层的单词有多个，都要maintain下来考察，哪一条转换路径先遇到终点词，它就最短。

整理一下:
因为思路是由一个结点带出下一层的邻接点，用BFS，把单词看作节点。
维护一个队列，让起点词入列，level 为 1，然后出列考察。
将它的每个字符变成26字母之一，逐个看是否在单词表，如果在，这个新词为下一层的转变词。
将它入列，它的 level +1，并从单词表中删去这个词。
出列入列…重复，当出列的单词和终点词相同，说明遇到了终点词，返回它的 level。
当队列为空时，BFS结束，始终没有遇到终点词，没有路径通往终点，返回 0。
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
