//bfs+dfs(�����˫��bfs��Ч�������)
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
    //�ֵ����wordList�еĵ��ʷ���hash���У�������ң�
    dict:=make(map[string]bool,0)
    for _,v:=range wordList{
        dict[v]=true
    }
    //���endWord����hash���У���ʾ������ת���б����ؿռ���
    if !dict[endWord]{
        return [][]string{}
    }
    //����һ�����ʷ���hash���У�����ʵ���ڽӱ�����ͼ��
    dict[beginWord]=true
    //�����ڽӱ�
    graph:=make(map[string][]string,0)
    //ִ��bfs�������ҵ�ÿ���㵽endWord�ľ���
    distance:=bfs(endWord,dict,graph)
    res:=make([][]string,0)//������
    //ִ��dfs����
    dfs(beginWord,endWord,&res,[]string{},distance,graph)
    return res
}

//����ʵ�ַ�ʽһ��������ƫ�������������ģ�壩
func dfs(beginWord string,endWord string,res *[][]string,path []string,distance map[string]int,graph map[string][]string){
    //���ݹ�����
    if beginWord==endWord{
        path=append(path,beginWord) //����endWord�ڵ�
        tmp:=make([]string,len(path))
        copy(tmp,path)
        (*res)=append((*res),tmp)
        path=path[:len(path)-1] //�Ƴ�endWord�ڵ�
        return
    }
    //�������ͼ
    for _,v:=range graph[beginWord]{
        //����ͼʱ�����ž������յ�Խ��Խ���ķ�����У���ǰ�ڵ�ľ���϶�Ҫ����һ�������1��
        if distance[beginWord]==distance[v]+1{
            path=append(path,beginWord)
            dfs(v,endWord,res,path,distance,graph)
            //���ݣ�ִ��������������ʱ��������ݻ�ȥ��
            path=path[:len(path)-1]
        }
    }
}
//����ʵ�ַ�ʽ����
func dfs(beginWord string,endWord string,res *[][]string,path []string,distance map[string]int,graph map[string][]string){
    path=append(path,beginWord)
    //���ݹ�����
    if beginWord==endWord{
        tmp:=make([]string,len(path))
        copy(tmp,path)
        (*res)=append((*res),tmp)
        return
    }
    //�������ͼ
    for _,v:=range graph[beginWord]{
        //����ͼʱ�����ž������յ�Խ��Խ���ķ�����У���ǰ�ڵ�ľ���϶�Ҫ����һ�������1��
        if distance[beginWord]==distance[v]+1{
            dfs(v,endWord,res,path,distance,graph)
        }
    }
    //���ݣ�ִ��������������ʱ��������ݻ�ȥ��
    path=path[:len(path)-1]
}


//���յ����������bfs������ÿһ���㵽���յ�ľ���
func bfs(endWord string,dict map[string]bool,graph map[string][]string)map[string]int{
    distance:=make(map[string]int,0) //ÿ���㵽�յ�ľ���
    queue:=make([]string,0)
    queue=append(queue,endWord)
    distance[endWord]=0 //��ʼֵ
    for len(queue)!=0{
        cursize:=len(queue)
        for i:=0;i<cursize;i++{
            word:=queue[0]
            queue=queue[1:]
            //�ҵ���word��һλ��ͬ�ĵ����б�
            expansion:=expand(word,dict)
            for _,v:=range expansion{
                //�����ڽӱ�
                //�����Ǵ�beginWord��endWord�����ڽӱ���bfs�Ǵ�endWord��ʼ�����Թ���ʱ������������
                //��graph[v]=append(graph[v],word)������graph[word]=append(graph[word],v)
                graph[v]=append(graph[v],word)
                //��ʾû���ʹ�
                if _,ok:=distance[v];!ok{
                    distance[v]=distance[word]+1 //�����һ
                    queue=append(queue,v) //�����
                }
            }
        }
    }
    return distance
}

//����ڽӵ�
func expand(word string,dict map[string]bool)[]string{
    expansion:=make([]string,0) //����word���ڽӵ�
    //��word��ÿһλ��ʼ
    chs:=[]rune(word)
    for i:=0;i<len(word);i++{
        tmp:=chs[i] //���浱ǰλ�����������и�λ
        for c:='a';c<='z';c++{
            //���һ����ֱ��������֮������tmp������Ϊchs[i]�ڱ�
            if tmp==c{ 
                continue
            }
            chs[i]=c
            newstr:=string(chs)
            //�µ�����dict�в����ڣ�������
            if dict[newstr]{
                expansion=append(expansion,newstr)
            }
        }
        chs[i]=tmp //����λ��λ
    }
    return expansion
}

