package othellogo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOriginCoordinateToBitBoard(t *testing.T) {
	bb := coordinateToBitBoard(0, 0)

	assert.Equal(t, bb, BitBoard(0x8000000000000000))
}

func TestEdgeCoordinateToBitBoard(t *testing.T) {
	bb := coordinateToBitBoard(7, 7)

	assert.Equal(t, bb, BitBoard(0x0000000000000001))
}
