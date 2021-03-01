package othellogo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOriginCoordinateToBitBoard(t *testing.T) {
	bb := CoordinateToBitBoard(0, 0)

	assert.Equal(t, bb, BitBoard(0x8000000000000000))
}

func TestEdgeCoordinateToBitBoard(t *testing.T) {
	bb := CoordinateToBitBoard(7, 7)

	assert.Equal(t, bb, BitBoard(0x0000000000000001))
}

func TestBoardMakeLegalBoard(t *testing.T) {
	board := Board{
		0x0000000800000000,
		0x00001C141C000000,
	}

	legalBoard := board.MakeLegalBoard()

	assert.Equal(t, legalBoard, BitBoard(0x002A0022002A0000))
}

func TestBoardCanPutLegalPoint(t *testing.T) {
	board := Board{
		0x0000000800000000,
		0x00001C141C000000,
	}

	assert.True(t, board.CanPutPoint(2, 1))
}

func TestBoardCanPutIllegalPoint(t *testing.T) {
	board := Board{
		0x0000000800000000,
		0x00001C141C000000,
	}

	assert.False(t, board.CanPutPoint(0, 0))
}

func TestBoardReverse(t *testing.T) {
	board := Board{
		0x0000000800000000,
		0x00001C141C000000,
	}

	reversed := Board{
		0x0020100800000000,
		0x00000C141C000000,
	}

	board.Reverse(2, 1)

	assert.Equal(t, board, reversed)
}

func TestBoardTurnChange(t *testing.T) {
	board := Board{
		0x0000000800000000,
		0x00001C141C000000,
	}

	changed := Board{
		0x00001C141C000000,
		0x0000000800000000,
	}

	board.TurnChange()

	assert.Equal(t, board, changed)
}
