package lru_20260316

type CacheNode struct {
	key   int
	value int
}

type DLinkNode struct {
	cNode CacheNode
	pre   *DLinkNode
	next  *DLinkNode
}

type DLinkList struct {
	headNode *DLinkNode
	tailNode *DLinkNode
}

func NewDLinkList() *DLinkList {
	return &DLinkList{
		headNode: &DLinkNode{},
		tailNode: &DLinkNode{},
	}
}

func (d *DLinkList) insertToHead(key, value int) {
	node := &DLinkNode{
		cNode: CacheNode{
			key:   key,
			value: value,
		},
		pre:  nil,
		next: nil,
	}
	node.pre = d.headNode
	node.next = d.headNode.next
	node.pre.next = node
	node.next.pre = node
}

func (d *DLinkList) moveToHead(node *DLinkNode) {
	node.pre.next = node.next
	node.next.pre = node.pre

	node.pre = d.headNode
	node.next = d.headNode.next
	node.pre.next = node
	node.next.pre = node
}

func (d *DLinkList) removeTail() *DLinkNode {
	node := d.tailNode.pre
	node.pre.next = node.next
	node.next.pre = node.pre
	return node
}

type HashMapNode struct {
	node *DLinkNode
	next *HashMapNode
}

type HashMap struct {
	capacity int
	headers  []*HashMapNode
}

func NewHashMap(capacity int) *HashMap {
	hashMap := &HashMap{capacity: capacity}
	hashMap.headers = make([]*HashMapNode, 31)
	for index := range hashMap.headers {
		hashMap.headers[index] = &HashMapNode{}
	}
	return hashMap
}

func (h *HashMap) hash(key int) int {
	return key % len(h.headers)
}

func (h *HashMap) get(key int) *DLinkNode {
	header := h.headers[h.hash(key)]

	for next := header.next; next != nil; next = next.next {
		if next.node.cNode.key == key {
			return next.node
		}
	}
	return nil
}

func (h *HashMap) remove(key int) {
	header := h.headers[h.hash(key)]
	for current := header; current.next != nil; current = current.next {
		if current.next.node.cNode.key == key {
			current.next = current.next.next
		}
	}
}

func (h *HashMap) put(key int, node *DLinkNode) {
	hashNode := &HashMapNode{node: node}
	header := h.headers[h.hash(key)]
	hashNode.next = header.next
	header.next = hashNode
}

type LRUCache struct {
	capacity int
	size     int
	hash     *HashMap
	linkList *DLinkList
}

func NewLRUCache(capacity int) *LRUCache {
	l := &LRUCache{
		capacity: capacity,
		hash:     NewHashMap(capacity),
		linkList: NewDLinkList(),
	}
	l.linkList.headNode.next = l.linkList.tailNode
	l.linkList.tailNode.pre = l.linkList.headNode
	return l
}

func (l *LRUCache) Put(key int, value int) {
	if l.Get(key) != -1 {
		return
	}
	if l.size >= l.capacity {
		node := l.linkList.removeTail()
		l.hash.remove(node.cNode.key)
		l.size--
	}
	l.linkList.insertToHead(key, value)
	l.hash.put(key, l.linkList.headNode.next)
	l.size++
}

func (l *LRUCache) Get(key int) int {
	if node := l.hash.get(key); node != nil {
		l.linkList.moveToHead(node)
		return node.cNode.value
	}
	return -1
}
