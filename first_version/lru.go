package first_version

type CacheNode struct {
	key   int
	value int
}

type DLinkNode struct {
	cNode *CacheNode
	pre   *DLinkNode
	next  *DLinkNode
}

type DLinkList struct {
	headNode *DLinkNode
	tailNode *DLinkNode
}

func (l *DLinkList) moveToHead(node *DLinkNode) {
	node.pre.next = node.next
	node.next.pre = node.pre

	node.pre = l.headNode
	node.next = l.headNode.next
	node.pre.next = node
	node.next.pre = node
	return
}

func (l *DLinkList) insertToHead(key, value int) {
	node := &DLinkNode{
		cNode: &CacheNode{
			key:   key,
			value: value,
		},
		pre:  nil,
		next: nil,
	}

	node.pre = l.headNode
	node.next = l.headNode.next
	l.headNode = node
	node.next.pre = node
	return
}

func (l *DLinkList) removeTail() *DLinkNode {
	node := l.tailNode.pre
	node.pre.next = node.next
	node.next.pre = node.pre
	return node
}

type HashMapNode struct {
	dNode *DLinkNode
	next  *HashMapNode
}

type HashMap struct {
	capacity int
	heads    []*HashMapNode
	size     int
}

func (h *HashMap) hash(key int) int {
	return key % len(h.heads)
}

func (h *HashMap) get(key int) *DLinkNode {
	index := h.hash(key)
	current := h.heads[index]
	for current != nil {
		if current.dNode.cNode.key == key {
			return current.dNode
		}
		current = current.next
	}
	return nil
}

func (h *HashMap) put(key int, node *DLinkNode) {
	hNode := &HashMapNode{
		dNode: node,
	}
	index := h.hash(key)
	head := h.heads[index]
	hNode.next = head.next
	head.next = hNode
	h.size++
	return
}

func (h *HashMap) remove(key int) *DLinkNode {
	index := h.hash(key)
	current := h.heads[index]
	var prev *HashMapNode
	for current != nil {
		if current.dNode.cNode.key == key {
			if prev == nil {
				h.heads[index] = current.next
			} else {
				prev.next = current.next
			}
			h.size--
		}
		prev = current
	}

	return nil
}

func NewHashMap(capacity int) *HashMap {
	return &HashMap{
		capacity: capacity,
		heads:    make([]*HashMapNode, 31),
		size:     0,
	}
}

type LRUCache struct {
	capacity int
	hash     *HashMap
	linkList *DLinkList
}

func (c *LRUCache) Get(key int) int {
	if node := c.hash.get(key); node != nil {
		c.linkList.moveToHead(node)
		return node.cNode.value
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if c.Get(key) != -1 {
		return
	}
	if c.hash.size >= c.hash.capacity {
		node := c.linkList.removeTail()
		c.hash.remove(node.cNode.key)
	}
	c.linkList.insertToHead(key, value)
	c.hash.put(key, c.linkList.headNode.next)
	return
}

func NewLRUCache(capacity int) *LRUCache {
	cache := &LRUCache{
		capacity: capacity,
		hash:     NewHashMap(capacity),
		linkList: &DLinkList{
			headNode: &DLinkNode{},
			tailNode: &DLinkNode{},
		},
	}
	cache.linkList.headNode.next = cache.linkList.tailNode
	cache.linkList.tailNode.pre = cache.linkList.headNode
	return cache
}
