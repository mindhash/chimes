package hash_list
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewHashList(t *testing.T) {
	hashList := New(2, 5)
	assert.Equal(t, len(hashList.buckets), 2, "expected 2 buckets")
}

func TestStoreInHashList(t *testing.T) {
	hashList := New(2, 5)

	item, err := hashList.Store(323, []byte("data"))
	assert.NoError(t, err, "failed to store in hash")
	assert.NotNil(t, item, "failed to store an item in hash list")
}

func TestFindInHashList(t *testing.T) {
	hashList := New(2, 5)

	item, err := hashList.Store(323, []byte("data"))
	assert.NoError(t, err, "failed to store in hash")
	assert.NotNil(t, item, "failed to store an item in hash list")

	found, err := hashList.Find(323)
	assert.NoError(t, err, "expected key 323 in the hash list")
	assert.NotNil(t, found, "failed to find a valid item in hash list")
}