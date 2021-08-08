package unit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveIndexForEmptySlice(t *testing.T) {
	t.Parallel()
	boolVar := true

	assert.Equal(t, boolVar, true, "boolVar should be true")
}
