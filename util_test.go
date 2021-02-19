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

func TestMakeLegalBoard(t *testing.T) {
	board := Board{
		0x0000000800000000,
		0x00001C141C000000,
	}

	legalBoard := makeLegalBoard(board)

	assert.Equal(t, legalBoard, BitBoard(0x002A0022002A0000))
}
