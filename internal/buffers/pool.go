package buffers

import (
	"bytes"
	"errors"
)

var (
	BufferLimitExceeded = errors.New("buffer limit exceeded")
	BufferSize = 256
)


type FreeList struct {
	// Max pool length
	MaxLength int

	// current length
	Length int
	
}

func  NewBuffer(data []byte) (*bytes.Buffer, error) {
	// check the length first 
	// if max length reached return bufferlimitexceeded error

	// return newe buffer 
	return bytes.NewBuffer(make([]byte, 0, BufferSize)), nil
}

