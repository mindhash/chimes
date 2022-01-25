package linked_list

import (
	"bytes"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	
	third := &ListItem {
		key: 3,
		data: bytes.NewBufferString("third"),
	}

	second := &ListItem {
		key: 2,
		data: bytes.NewBufferString("second"),
	}

	first:= &ListItem{key: 1,
		data: bytes.NewBufferString("first"),
	}
	
	first.next = second
	
	second.prev = first
	second.next = third

	third.prev = second

	dl := DoubleList{
		first: first,
		last: third,
	}

	found, err := dl.Find(second.key)
	assert.NoError(t, err, "failed to find given key")

	assert.NotNil(t, found, "item not found and no error raised")

	assert.Equal(t, found.key, second.key, "the item found is incorrect")
	
}

type keyVal struct {
	key int
	val []byte
}

var testData = []keyVal{
	{key: 1, val: []byte("data"),},
	{key: 2, val: []byte("data"),},
	{key: 3, val: []byte("data"),},
}

func TestStore(t *testing.T) {

	// case 1: test simple store and retrieve
	dl := New(2)

	storedItem, err := dl.Store(testData[0].key, testData[0].val)
	assert.NoError(t, err, "failed to store given data")
	assert.NotNil(t, storedItem, "item not stored and no error raised")

	found, err := dl.Find(testData[0].key)
	assert.NoError(t, err, "failed to find given key")
	assert.NotNil(t, found, "item not found and no error raised")

	assert.Equal(t, found.key, testData[0].key, "the item found is incorrect")
	
}

func TestStoreMemoryRead(t *testing.T) {
	// case 2: test store in absence of list space and retrieve
	dl := New(1)

	for _, testCase := range testData {
		storedItem, err := dl.Store(testCase.key, testCase.val)
		assert.NoError(t, err, "failed to store given data")
		assert.NotNil(t, storedItem, "item not stored and no error raised")
	}

	notfound, err := dl.Find(testData[0].key)
	assert.Nil(t, notfound, "unexpected item [0] in the list")
	assert.Error(t, err, "expected error for item[0] search")
}