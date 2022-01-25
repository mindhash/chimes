package cache

type Cache interface {
	// Fetch fetches an item for given key 
	Find(key int) (item CacheItem, err error)
	
	// Store saves a given item to cache
	Store(key int, data []byte) (item CacheItem, err error)
}