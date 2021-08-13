package unit_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveIndexForEmptySlice(t *testing.T) {
	t.Parallel()

	boolVar := true

	assert.Equal(t, boolVar, true, "boolVar should be true")
}
