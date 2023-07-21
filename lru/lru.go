package lru

type LRUCache struct {
	capacity int
	list *DoublyLinkedList
	cache map[string]*DoublyLinkedListNode
}

type cacheItem struct {
	key   string
	value any
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		list:     NewDoublyLinkedList(),
		cache:    make(map[string]*DoublyLinkedListNode),
	}
}

func (c *LRUCache) Get(key string) (value any, ok bool) {
	if node, ok := c.cache[key]; ok {
		c.list.MoveToFront(node)
		return node.Value.(*cacheItem).value, true
	}
	return nil, false
}

func (c *LRUCache) Set(key string, value any) {
	if node, ok := c.cache[key]; ok {
		node.Value.(*cacheItem).value = value
		c.list.MoveToFront(node)
		return
	}

	newNode := c.list.PushFront(&cacheItem{key: key, value: value})
	c.cache[key] = newNode

	if c.list.Len() > c.capacity {
		tail := c.list.RemoveTail()
		delete(c.cache, tail.Value.(*cacheItem).key)
	}
}

type DoublyLinkedListNode struct {
	Value      any
	Prev, Next *DoublyLinkedListNode
}

type DoublyLinkedList struct {
	head, tail *DoublyLinkedListNode
}

func NewDoublyLinkedList() *DoublyLinkedList {
	head := &DoublyLinkedListNode{}
	tail := &DoublyLinkedListNode{}

	head.Next = tail
	tail.Prev = head

	return &DoublyLinkedList{
		head: head,
		tail: tail,
	}
}

func (l *DoublyLinkedList) PushFront(v any) *DoublyLinkedListNode {
	newNode := &DoublyLinkedListNode{Value: v}
	newNode.Next = l.head.Next
	newNode.Prev = l.head
	l.head.Next.Prev = newNode
	l.head.Next = newNode

	return newNode
}

func (l *DoublyLinkedList) Len() int {
	dist := 0
	curr := l.head.Next

	for curr != l.tail {
		dist++
		curr = curr.Next
	}

	return dist
}

func (l *DoublyLinkedList) RemoveTail() *DoublyLinkedListNode {
	if l.tail.Prev == l.head {
		return nil
	}

	tail := l.tail.Prev
	l.RemoveNode(tail)
	return tail
}

func (l *DoublyLinkedList) AppendTail(value any) *DoublyLinkedListNode {
	newNode := &DoublyLinkedListNode{Value: value}

	newNode.Prev = l.tail.Prev
	newNode.Next = l.tail

	l.tail.Prev.Next = newNode
	l.tail.Prev = newNode

	return newNode
}

func (l *DoublyLinkedList) RemoveNode(node *DoublyLinkedListNode) {
	if node == l.head || node == l.tail {
		return
	}

	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (l *DoublyLinkedList) MoveToFront(node *DoublyLinkedListNode) {
	l.RemoveNode(node)
	l.PushFront(node.Value)
}
