package lru

func newCacheNode(key, value interface{}) *cacheNode {
	return &cacheNode{key: key, value: value}
}

// cacheNode wraps the cached value
type cacheNode struct {
	key   interface{}
	value interface{}
	next  *cacheNode
	prev  *cacheNode
}

// free to avoid any memory leak
func (c *cacheNode) free() {
	c.key, c.value, c.next, c.prev = nil, nil, nil, nil
}
