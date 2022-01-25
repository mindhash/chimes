package linked_list
import (
	"log"
	"errors"
	"sync/atomic"
	"github.com/mindhash/chimes/internal/buffers"
)

// Store initiates a new item and stores at the start of list
func (dl *DoubleList) Store(key int, data []byte) (item *ListItem, err error) {
	return dl.newItem(key, data)
}

// newItem stores a given data object with key at the start of the list 
func (dl *DoubleList) newItem(key int, b []byte) (*ListItem, error) {
	// todo: make sure newItem is thread safe or runs only once
	
	var item *ListItem
	var err error 
	log.Println("size of list:", dl.size, dl.MaxSize)

	if dl.size < dl.MaxSize {
		// log.Println("free list available. picking...")

		// free space available
		item, err = dl.createItemFromFreeList(key, b)

		if err != nil {
			if !errors.Is(err, buffers.BufferLimitExceeded) {
				return nil, err
			}
		}
		
		if item != nil {
			// buffer limit has not reached, we have an item
			atomic.AddInt32(&dl.size, 1)
			return item, err
		}
		
	}

	// free list is full. pick a LRU buffer
	return dl.createFromEndOfList(key, b)
}

// createItemFromFreeList creates an item by picking a buffer from freelist
func (dl *DoubleList) createItemFromFreeList(key int, b []byte) (*ListItem, error) {
		
	var item *ListItem

	// ask for a new buffer 
	emptyBuf, err := buffers.NewBuffer(b)
	if err != nil {
		return nil, err
	}

	// received a new buffer, lets create an item
	// emptyBuf.Write(b)
	item = &ListItem{
		key: key,
		data: emptyBuf, 
	}

	dl.lock.Lock()
	
	defer func() {
		dl.lock.Unlock()
	}()
	
	// move the item to start of list 
	item = dl.moveToStart(item, false)
	
	return item, nil 
}

// createFromEndOfList creates an item by picking existing buffer (LRU) from end of the list
func (dl *DoubleList) createFromEndOfList(key int, b []byte) (*ListItem, error) {
	var item *ListItem
	var err error
	// new  buffers are not free. lets pick existing 
	dl.lock.Lock()
	
	defer func() {
		dl.lock.Unlock()
	}() 

	// get from the end of the list 
	item = dl.pluckLast(false)
	
	// empty the contents from item 
	item.Reset()

	item.SetKey(key)
	// write new content 
	item.SetBytes(b)
	
	// move the item to start of list 
	item = dl.moveToStart(item, false) 
	return item, err
}


