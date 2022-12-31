package internal

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapLeftValue(t *testing.T) {
	e := NewLeft[int, error](1)
	afterMap := MapLeft(e, func(v int) string { return fmt.Sprint(v + 1) })

	assert.Equal(t, "2", afterMap.Left.value)
}

func TestMapLeftError(t *testing.T) {
	e := NewLeft[error, int](errors.New("something went wrong"))
	afterMap := MapLeft(e, func(v error) string { return v.Error() })

	assert.Equal(t, "something went wrong", afterMap.Left.value)
}

func TestMapRightValue(t *testing.T) {
	e := NewRight[int, error](errors.New("something went wrong"))
	afterMap := MapRight(e, func(v error) string { return v.Error() })

	assert.Equal(t, "something went wrong", afterMap.Right.value)
}
