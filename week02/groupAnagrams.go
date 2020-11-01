/*解题思路
遍历数组，然后对数组中的字符串排序（先将其转为自定义bytes类型，然后对其排序，再转为字符串，即是排序后的字符串），若该字符串未出现过记录拍完序的字符串以及字符串在ret的位置，加入到ret中，否则加入到ret中已经出现的地方。
*/



type bytes []byte
// 自定义排序规则
func (b bytes) Len() int {
    return len(b)
}
func (b bytes) Less(i,j int) bool {
    return b[i] < b[j]
}
func (b bytes) Swap(i, j int) {
    b[i], b[j] = b[j], b[i]
}
func groupAnagrams(strs []string) [][]string {
    ret := [][]string{}
    m := make(map[string]int)
    for _, str := range strs {
        kByte := bytes(str)
        sort.Sort(kByte) // 将字符串排序
        k := string(kByte) 
        if idx, ok := m[k]; !ok {
            m[k] = len(ret) // 记录拍完序的字符串以及字符串在ret的位置
            ret = append(ret, []string{str})
        } else { // 已经出现过，将其放入出现在ret中，在ret的位置为idx
            ret[idx] = append(ret[idx], str)
        }
    }
    return ret
}

