package erroum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestIsAny tests the evaluation of error != nil.
func TestIsAny(t *testing.T) {
	err1 := New("error 1")

	t.Run("with error", func(t *testing.T) {
		assert.True(t, IsAny(err1))
	})
	t.Run("with nil error", func(t *testing.T) {
		assert.False(t, IsAny(nil))
	})
}

// TestIs tests the evaluation of error -> error
func TestIs(t *testing.T) {
	err1 := New("error 1")
	err2 := New("error 2")

	t.Run("with same errors", func(t *testing.T) {
		assert.True(t, Is(err1, err1))
	})
	t.Run("with different errors", func(t *testing.T) {
		assert.False(t, Is(err1, err2))
	})
	t.Run("with nil error to evaluate", func(t *testing.T) {
		assert.False(t, Is(nil, err2))
	})
	t.Run("with nil error to evaluate against", func(t *testing.T) {
		assert.False(t, Is(err1, nil))
	})
	t.Run("with nil error at both", func(t *testing.T) {
		assert.False(t, Is(nil, nil))
	})
}

// TestIsSome tests the evaluation of error -> n*errors
func TestIsSome(t *testing.T) {
	err1 := New("error 1")
	err2 := New("error 2")
	err3 := New("error 3")

	t.Run("with same errors", func(t *testing.T) {
		assert.True(t, IsSome(err1, err1))
	})
	t.Run("with many errors", func(t *testing.T) {
		t.Run("with match", func(t *testing.T) {
			assert.True(t, IsSome(err1, err2, err1))
		})

		t.Run("without match", func(t *testing.T) {
			assert.False(t, IsSome(err1, err2, err3))
		})

		t.Run("with nil error to evaluate", func(t *testing.T) {
			assert.False(t, IsSome(nil, err2, err3))
		})

		t.Run("with nil error to evaluate against", func(t *testing.T) {
			assert.False(t, IsSome(err1, nil, nil))
		})

		t.Run("with some nil error to evaluate against", func(t *testing.T) {
			assert.True(t, IsSome(err1, nil, nil, nil, err1))
		})

		t.Run("with nil error at both", func(t *testing.T) {
			assert.False(t, IsSome(nil, nil, nil))
		})
	})
}
