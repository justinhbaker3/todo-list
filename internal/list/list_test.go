package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	l := NewList("TestList")
	err := l.Append(&Item{
		Title:       "TestItem",
		Description: "test",
	})
	assert.NoError(t, err)
	assert.Equal(t, 1, len(l.Items))
}
