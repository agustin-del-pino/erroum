package erroum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNew tests the creation of a new error.
func TestNew(t *testing.T) {
	t.Run("with description", func(t *testing.T) {
		err := New("error")

		assert.NotNil(t, err)
		assert.Equal(t, "error", err.Error())
	})

	t.Run("without description", func(t *testing.T) {
		assert.PanicsWithValue(t, "empty errors are not admitted", func() { New("") })
	})
}

// TestFrom tests the creation of an error from another.
func TestFrom(t *testing.T) {
	err1 := New("error 1")
	err2 := New("error 2")
	t.Run("with only one error", func(t *testing.T) {
		err := From("->", "error", err1)

		assert.NotNil(t, err)
		assert.Equal(t, "error->error 1", err.Error())
	})

	t.Run("with many errors", func(t *testing.T) {
		err := From("->", "error", err1, err2)

		assert.NotNil(t, err)
		assert.Equal(t, "error->error 1->error 2", err.Error())
	})

	t.Run("with only description", func(t *testing.T) {
		err := From("->", "error")

		assert.NotNil(t, err)
		assert.Equal(t, "error", err.Error())
	})

	t.Run("without description", func(t *testing.T) {
		assert.PanicsWithValue(t, "empty errors are not admitted", func() { From("->", "") })
	})
}

// TestMerge tests the creation of an error from merging.
func TestMerge(t *testing.T) {
	err1 := New("error 1")
	err2 := New("error 2")
	t.Run("with only one error", func(t *testing.T) {
		err := Merge("->", err1)

		assert.NotNil(t, err)
		assert.Equal(t, "error 1", err.Error())
	})

	t.Run("with many errors", func(t *testing.T) {
		err := Merge("->", err1, err2)

		assert.NotNil(t, err)
		assert.Equal(t, "error 1->error 2", err.Error())
	})

	
	t.Run("with nil errors", func(t *testing.T) {
		err := Merge("->", err1, err2, nil, err1, nil)

		assert.NotNil(t, err)
		assert.Equal(t, "error 1->error 2->error 1", err.Error())
	})

	t.Run("without error", func(t *testing.T) {
		assert.PanicsWithValue(t, "cannot merger from no error", func() { Merge("->", nil) })
	})
}
