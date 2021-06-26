package 其他

/**
运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制 。
实现 LRUCache 类：

LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。


进阶：你是否可以在 O(1) 时间复杂度内完成这两种操作？



示例：

输入
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]

解释
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4


提示：

1 <= capacity <= 3000
0 <= key <= 3000
0 <= value <= 104
最多调用 3 * 104 次 get 和 put

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lru-cache
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//只写了O(N)的  O（1）据说要双链表
//链表还是要画呀 不画还是比较难理清的

type LRUCache struct {
	Capacity int
	Head     *DoubleList
	End      *DoubleList
	m        map[int]*DoubleList
	Size     int
}

//双链表
type DoubleList struct {
	Next *DoubleList
	Prev *DoubleList
	Val  int
	Key  int
}

func Constructor(capacity int) LRUCache {
	Head := &DoubleList{nil, nil, 0, 0}
	End := &DoubleList{nil, nil, 0, 0}
	Head.Next = End
	End.Prev = Head
	m := make(map[int]*DoubleList)
	l := LRUCache{capacity, Head, End, m, 0}
	return l
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.m[key]; ok {
		v.Prev.Next = v.Next
		v.Next.Prev = v.Prev
		v.Prev = this.Head
		v.Next = this.Head.Next
		v.Next.Prev = v
		this.Head.Next.Prev = v
		this.Head.Next = v
		return v.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.m[key]; ok {
		//如果本身有 则把这个移到头部 然后修改值
		v.Val = value
		v.Prev.Next = v.Next
		v.Next.Prev = v.Prev
		v.Prev = this.Head
		v.Next = this.Head.Next
		v.Next.Prev = v
		this.Head.Next.Prev = v
		this.Head.Next = v
		this.m[key] = v
	} else {
		//不存在
		this.Size++
		if this.Size > this.Capacity {
			//超过容量 移除末尾的
			removeNode := this.End.Prev
			removeNode.Prev.Next = this.End
			this.End.Prev = removeNode.Prev
			removeNode.Next = nil
			removeNode.Prev = nil
			delete(this.m, removeNode.Key)
			this.Size--
		}
		node := &DoubleList{nil, nil, value, key}
		node.Prev = this.Head
		node.Next = this.Head.Next
		this.Head.Next = node
		node.Next.Prev = node
		this.m[key] = node
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
