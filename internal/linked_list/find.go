package linked_list

import "log"

func (dl *DoubleList)	Find(key int) (item *ListItem, err error){
	var result *ListItem
	var current *ListItem

	current = dl.first 

	if dl.first == nil {
		return nil, ListEmptyError
	}
	
	// loop through list until hasKey returns true
	for {	
		
		if current.HasKey(key) {
			result = current
			break
		} 

		if current.next != nil {
			current = current.next
			
		} else {
			// reached end of the list. item doesn't exist 
			return nil, KeyDoesNotExistError
		}
	}
	
	if dl.first == current {
		// already first no plucking or moving
		return current, nil
	}
	
	// acquire lock 
	dl.lock.Lock()

	defer func() {
		dl.lock.Unlock()
	}()

	// found the node in the list. pluck it now	
	result = dl.pluck(result, false)

	// move the item to the top of the list
	result = dl.moveToStart(result, false)
	
	// return the item 
	return current, nil
}
