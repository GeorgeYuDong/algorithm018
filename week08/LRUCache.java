class Node {
    public int key, val;
    public Node next, prev;
    public Node(int k, int v) {
        this.key = k;
        this.val = v;
    }
};

class DoubleList {  
    // ������ͷ����ӽڵ� x��ʱ�� O(1)
    public void addFirst(Node x);

    // ɾ�������е� x �ڵ㣨x һ�����ڣ�
    // ������˫�����Ҹ�����Ŀ�� Node �ڵ㣬ʱ�� O(1)
    public void remove(Node x);
    
    // ɾ�����������һ���ڵ㣬�����ظýڵ㣬ʱ�� O(1)
    public Node removeLast();
    
    // ���������ȣ�ʱ�� O(1)
    public int size();
};

class LRUCache {
    // key -> Node(key, val)
    private HashMap<Integer, Node> map;
    // Node(k1, v1) <-> Node(k2, v2)...
    private DoubleList cache;
    // �������
    private int cap;
    
    public LRUCache(int capacity) {
        this.cap = capacity;
        map = new HashMap<>();
        cache = new DoubleList();
    }
    
    public int get(int key) {
        if (!map.containsKey(key))
            return -1;
        int val = map.get(key).val;
        // ���� put �����Ѹ�������ǰ
        put(key, val);
        return val;
    }
    
    public void put(int key, int val) {
        // �Ȱ��½ڵ� x ������
        Node x = new Node(key, val);
        
        if (map.containsKey(key)) {
            // ɾ���ɵĽڵ㣬�µĲ嵽ͷ��
            cache.remove(map.get(key));
            cache.addFirst(x);
            // ���� map �ж�Ӧ������
            map.put(key, x);
        } else {
            if (cap == cache.size()) {
                // ɾ���������һ������
                Node last = cache.removeLast();
                map.remove(last.key);
            }
            // ֱ����ӵ�ͷ��
            cache.addFirst(x);
            map.put(key, x);
        }
    }
}

