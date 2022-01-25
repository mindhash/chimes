package linked_list

import (
	"sync"
	"errors"
)
const DefaultMaxListSize = 64

var (
	ListEmptyError = errors.New("list is empty")
	KeyDoesNotExistError = errors.New("key does not exist")
)
 
type DoubleList struct {
	
	lock sync.Mutex

	first *ListItem
	last *ListItem
	size int32
	MaxSize int32
}

// New creates a new double list with a given size (as max)
func New(size int32) *DoubleList {
	
	l := DoubleList{
		MaxSize: size,
	}

	if l.MaxSize == 0 {
		l.MaxSize = DefaultMaxListSize
	}

	return &l
}
 

func (dl *DoubleList) pluckLast(withLock bool) *ListItem {
	var item *ListItem

	if dl.last != nil {
		item = dl.last
		if item.next != nil {
			//unexpected raise error
		}
		
		// make prev item last in the list
		if item.prev != nil {
			if withLock {
				dl.lock.Lock()
				defer func() {
					dl.lock.Unlock()
				}()
			}
			prevItem := item.prev
			prevItem.next = nil
			dl.last = prevItem
		}

		// reset links 
		item.prev = nil
		item.next = nil 
	}

	return item
}

func (dl *DoubleList) pluck(item *ListItem, withLock bool) *ListItem{
	if withLock {
		dl.lock.Lock()

		defer func() {
			dl.lock.Unlock()
		}()
	}
	
	if item == nil {
		return nil
	}

	if (item.prev == nil && item.next == nil) {
		// first item. nothing to do
		return item
	}

	if item.next != nil {
		// item is in the middle 
		nextItem := item.next 
		prevItem := item.prev

		prevItem.next = nextItem
		nextItem.prev = prevItem

		item.prev = nil
		item.next = nil
		
	} else {
		// last item 
		prevItem := item.prev
		
		// move prev item to end of the list
		prevItem.next = nil
		dl.last = prevItem

		item.prev = nil
	}
	
	return item 

}

func (dl *DoubleList) moveToStart(item *ListItem, withLock bool) *ListItem{
	if withLock {
		// acquire write lock only todo
		dl.lock.Lock()

		defer func() {
			dl.lock.Unlock()
		}()

	}
	
	if item == nil  {
		return nil
	}

	if dl.first != nil {
		
		if dl.first == item {
			// already first item
			return item
		}

		// current first.prev -> new item
		dl.first.prev = item
		
		// new item.next -> current first 
		item.next = dl.first
	}
	
	// empty prev for new item
	item.prev = nil

	// new first -> item
	dl.first = item
	
	// if this is first item then mark as last as well
	if dl.last == nil {
		dl.last = item
	}

	return item
}


