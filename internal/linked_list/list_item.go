package linked_list
import ( 
	"bytes"
)

type ListItem struct {
	key int
	data *bytes.Buffer
	next *ListItem
	prev *ListItem
}

func (li *ListItem) SetKey(k int) {
	li.key = k
}

func (li *ListItem) SetBytes(b []byte) {
	li.data.Write(b)
}

func (li *ListItem) Reset() {
	if li != nil {

		li.ResetData()
		li.ResetKey()
	}
}

func (li *ListItem) ResetKey() {
	li.key = 0
}

func (li *ListItem) ResetData() {
	li.data.Reset()
}

func (li *ListItem) GetKey() int {
	return li.key
}

func (li *ListItem) GetData() []byte {
	return li.data.Bytes()
}

func (li *ListItem) HasKey(key int) bool {
	if li.key == key {
		return true
	}
	return false
}
