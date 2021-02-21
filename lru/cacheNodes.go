package lru

func newCacheNodes(capacity int) *cacheNodes {
	return &cacheNodes{capacity: capacity}
}

// cacheNodes maintains a doubly linked list of cacheNode.
// Front of list is the least recently used node and last is the most recently used node.
type cacheNodes struct {
	first    *cacheNode
	last     *cacheNode
	capacity int
	size     int
}

// markMostRecentlyUsed adds the node to the end of cacheNodes
// assumption: it's for internal use and the edge cases are taken care by the higher component.
func (c *cacheNodes) markMostRecentlyUsed(node *cacheNode) {
	// if just one node in collection
	if c.size == 1 || c.last == node {
		return
	}
	if node == c.first {
		c.first = node.next
		c.first.prev = nil
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	node.next = nil
	node.prev = c.last
	c.last.next = node
	c.last = node
}

// removeLeastRecentlyUsed deletes the node from front of cacheNodes and returns it
// assumption: it's for internal use and the edge cases are taken care by the higher component.
func (c *cacheNodes) removeLeastRecentlyUsed() (node *cacheNode) {
	defer func() { c.size-- }()
	node = c.first
	c.first = c.first.next
	if c.first == nil { // just one node in collection
		c.last = nil
	} else {
		c.first.prev = nil
	}

	return
}

// append adds the cacheNode to last
// assumption: it's for internal use and the edge cases are taken care by the higher component.
func (c *cacheNodes) append(node *cacheNode) {
	defer func() { c.size++ }()
	if c.size == 0 { // very first node
		c.first, c.last = node, node
		return
	}
	node.prev = c.last
	c.last.next = node
	c.last = node
}

// free to clear the cacheNodes
func (c *cacheNodes) free() {
	c.first, c.last = nil, nil
	c.size = 0
}
