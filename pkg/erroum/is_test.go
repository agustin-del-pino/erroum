package erroum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestInCaseOfAny tests the evaluation of error != nil for retuning some value.
func TestInCaseOfAny(t *testing.T) {
	err1 := New("error 1")

	t.Run("with error", func(t *testing.T) {
		assert.Equal(t, 500, InCaseOfAny(err1, 500, 200))
	})
	t.Run("with nil error", func(t *testing.T) {
		assert.Equal(t, 200, InCaseOfAny(nil, 500, 200))
	})
}

// TestInCaseOf tests the evaluation of error -> error for retuning some value.
func TestInCaseOf(t *testing.T) {
	err1 := New("error 1")
	err2 := New("error 2")

	t.Run("with same errors", func(t *testing.T) {
		assert.Equal(t, 500, InCaseOf(err1, err1, 500, 200))
	})
	t.Run("with different errors", func(t *testing.T) {
		assert.Equal(t, 400, InCaseOf(err1, err2, 500, 400))
	})
	t.Run("with nil error to evaluate", func(t *testing.T) {
		assert.Equal(t, 200, InCaseOf(nil, err2, 500, 200))
	})
	t.Run("with nil error to evaluate against", func(t *testing.T) {
		assert.Equal(t, 200, InCaseOf(err1, nil, 500, 200))
	})
	t.Run("with nil error at both", func(t *testing.T) {
		assert.Equal(t, 200, InCaseOf(nil, nil, 500, 200))
	})
}

// TestInCaseOfSome tests the evaluation of error -> n*errors for retuning some value.
func TestInCaseOfSome(t *testing.T) {
	err1 := New("error 1")
	err2 := New("error 2")
	err3 := New("error 3")

	t.Run("with same errors", func(t *testing.T) {
		assert.Equal(t, 500, InCaseOfSome(err1, 500, 200, err1))
	})
	t.Run("with many errors", func(t *testing.T) {
		t.Run("with match", func(t *testing.T) {
			assert.Equal(t, 500, InCaseOfSome(err1, 500, 200, err2, err1))
		})

		t.Run("without match", func(t *testing.T) {
			assert.Equal(t, 200, InCaseOfSome(err1, 500, 200, err2, err3))
		})

		t.Run("with nil error to evaluate", func(t *testing.T) {
			assert.Equal(t, 200, InCaseOfSome(nil, 500, 200, err2, err3))
		})

		t.Run("with nil error to evaluate against", func(t *testing.T) {
			assert.Equal(t, 200, InCaseOfSome(err1, 500, 200, nil, nil))
		})

		t.Run("with some nil error to evaluate against", func(t *testing.T) {
			assert.Equal(t, 500, InCaseOfSome(err1, 500, 200, nil, nil, nil, err1))
		})

		t.Run("with nil error at both", func(t *testing.T) {
			assert.Equal(t, 200, InCaseOfSome(nil, 500, 200, nil, nil))
		})
	})
}
