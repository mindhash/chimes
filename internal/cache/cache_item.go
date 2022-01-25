package cache
import (
	"bytes"
)

type CacheItem interface {
	GetKey() int
	GetData() []bytes
	HasKey(key int) bool
}

