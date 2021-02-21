package lru

// NewCache returns a new instance of lru.Cache
func NewCache(capacity int) *Cache {
	if capacity == 0 {
		return nil
	}
	return &Cache{
		capacity:       capacity,
		keyToCacheNode: make(map[interface{}]*cacheNode, capacity),
		nodes:          newCacheNodes(capacity),
	}
}

// Cache is the thread-safe implementation In-Memory Cache with Least Recently Used eviction policy
type Cache struct {
	capacity       int
	keyToCacheNode map[interface{}]*cacheNode
	nodes          *cacheNodes
}

// Get returns the cached value against given key and cacheFound if the key is found in cache.e False value of cacheFound means a Cache-Miss.
// TimeComplexity: O(1), SpaceComplexity: O(1)
func (c *Cache) Get(key interface{}) (value interface{}, cacheFound bool) {
	node, cacheFound := c.keyToCacheNode[key]
	if cacheFound {
		c.nodes.markMostRecentlyUsed(node)
		value = node.value
	}
	return
}

// Set caches the key, value pair provided.
// TimeComplexity: O(1), SpaceComplexity: O(1)
func (c *Cache) Set(key, value interface{}) {
	existingNode, keyIsPresent := c.keyToCacheNode[key]
	if keyIsPresent {
		c.nodes.markMostRecentlyUsed(existingNode)
		existingNode.value = value
		return
	}

	if len(c.keyToCacheNode) == c.capacity { // cache is full
		node := c.nodes.removeLeastRecentlyUsed() // LRU eviction
		delete(c.keyToCacheNode, node.key)
		node.free() // to avoid memory leak
	}

	node := newCacheNode(key, value)
	c.keyToCacheNode[key] = node
	c.nodes.append(node)
}

// Free clears the cache. Once cleared, cache is not recoverable.
func (c *Cache) Free() {
	c.nodes.free()
	for _, node := range c.keyToCacheNode {
		node.free()
	}
	c.keyToCacheNode = nil
}
