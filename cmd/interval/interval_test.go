package interval

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntervals_Sort(t *testing.T) {
	input := [][]uint64{[]uint64{1, 3}, []uint64{6, 15}, []uint64{15, 18}, []uint64{2, 6}}

	t.Run("sorted", func(t *testing.T) {
		toTest, err := NewIntervals(input)
		assert.Nil(t, err)

		toTest.Sort()

		inputSorted := [][]uint64{[]uint64{1, 3}, []uint64{2, 6}, []uint64{6, 15}, []uint64{15, 18}}
		sorted, err := NewIntervals(inputSorted)
		assert.Nil(t, err)

		assert.Equal(t, toTest, sorted)
	})
}

func TestInterval_Overlaps(t *testing.T) {
	input := [][]uint64{[]uint64{1, 3}, []uint64{6, 15}, []uint64{15, 18}, []uint64{2, 6}}

	t.Run("overlapping", func(t *testing.T) {
		toTest, err := NewIntervals(input)
		assert.Nil(t, err)

		assert.True(t, toTest[0].Overlaps(toTest[3]))
	})

	t.Run("not overlapping", func(t *testing.T) {
		toTest, err := NewIntervals(input)
		assert.Nil(t, err)

		assert.False(t, toTest[0].Overlaps(toTest[1]))
	})
}

func TestInterval_Merge(t *testing.T) {
	input := [][]uint64{[]uint64{1, 3}, []uint64{1, 2}, []uint64{15, 18}, []uint64{2, 6}}

	t.Run("merge", func(t *testing.T) {
		toTest, err := NewIntervals(input)
		assert.Nil(t, err)

		toTest[0].Merge(toTest[3])

		assert.Equal(t, toTest[0].end, toTest[3].end)
	})

	t.Run("merge 2", func(t *testing.T) {
		toTest, err := NewIntervals(input)
		assert.Nil(t, err)

		toTest[0].Merge(toTest[1])

		assert.Equal(t, toTest[0].end, toTest[0].end)
	})
}
