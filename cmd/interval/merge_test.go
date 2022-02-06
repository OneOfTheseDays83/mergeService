package interval

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerge(t *testing.T) {
	input := [][]uint64{[]uint64{1, 3}, []uint64{8, 15}, []uint64{15, 18}, []uint64{2, 6}}

	t.Run("test 1", func(t *testing.T) {
		result, err := Merge(input)

		assert.Nil(t, err)
		solution := [][]uint64{[]uint64{1, 6}, []uint64{8, 18}}
		assert.Equal(t, solution, result)
	})

	// TODO extend
}
