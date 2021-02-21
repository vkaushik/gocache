package gocache

// Cache is the interface for In-Memory-Cache
type Cache interface {
	Get(key interface{}) (value interface{}, cacheFound bool)
	Set(key, value interface{})
	Free()
}
