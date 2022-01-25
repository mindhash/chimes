package hash_list

import (
	"fmt"
	"github.com/mindhash/chimes/internal/linked_list"
)

type HashList struct {
	noOfBuckets int32
	bucketSize int32
	buckets map[int]*linked_list.DoubleList 
}

func New(noOfBuckets int32, bucketSize int32) *HashList{
  hl:= HashList{
		noOfBuckets: noOfBuckets,
		bucketSize: bucketSize,
	}

	buckets := make(map[int]*linked_list.DoubleList, noOfBuckets)

	for b:=0; b < int(noOfBuckets); b++ {
		bufferList := linked_list.New(bucketSize)
		buckets[b] = bufferList
	}

	hl.buckets = buckets
	return &hl
}

func (h *HashList) findBucket(key int) int {
	return key % int(h.noOfBuckets) 
}

func (h *HashList) Store(key int, data []byte) (item *linked_list.ListItem, err error) {
	bucketID := h.findBucket(key)

	if bucket, ok := h.buckets[bucketID]; ok {
		return bucket.Store(key, data)
	} 

	return nil, fmt.Errorf("bucket not initialised for given hash key")
}


func (h *HashList) Find(key int) (item *linked_list.ListItem, err error){ 
	bucketID := h.findBucket(key)

	if bucket, ok := h.buckets[bucketID]; ok {
		return bucket.Find(key)
	} 
	
	return nil, fmt.Errorf("failed to identify bucket for given key")
}