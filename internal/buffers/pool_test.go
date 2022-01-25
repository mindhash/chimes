package buffers
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewBuffer(t *testing.T) {
	b, err := NewBuffer()
	assert.NoError(t, err, "failed to allocate buffer")
	assert.NotEqual(t, b.Cap(), 0, "buffer is empty")
}

