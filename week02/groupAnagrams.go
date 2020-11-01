/*����˼·
�������飬Ȼ��������е��ַ��������Ƚ���תΪ�Զ���bytes���ͣ�Ȼ�����������תΪ�ַ����������������ַ������������ַ���δ���ֹ���¼��������ַ����Լ��ַ�����ret��λ�ã����뵽ret�У�������뵽ret���Ѿ����ֵĵط���
*/



type bytes []byte
// �Զ����������
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
        sort.Sort(kByte) // ���ַ�������
        k := string(kByte) 
        if idx, ok := m[k]; !ok {
            m[k] = len(ret) // ��¼��������ַ����Լ��ַ�����ret��λ��
            ret = append(ret, []string{str})
        } else { // �Ѿ����ֹ���������������ret�У���ret��λ��Ϊidx
            ret[idx] = append(ret[idx], str)
        }
    }
    return ret
}

