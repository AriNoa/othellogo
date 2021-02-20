package othellogo

// CoordinateToBitBoard returns BitBoard that is flagged only at the specified coordinates
func CoordinateToBitBoard(x int, y int) BitBoard {
	var bb BitBoard = 0x8000000000000000

	bb = bb >> x
	bb = bb >> (y * 8)

	return bb
}

// MakeLegalBoard returns BitBoard with flags only on the squares where the player can be placed
func (board Board) MakeLegalBoard() BitBoard {
	horizontalWatchBoard := board.Opponent & 0x7e7e7e7e7e7e7e7e
	verticalWatchBoard := board.Opponent & 0x00FFFFFFFFFFFF00
	allSideWatchBoard := board.Opponent & 0x007e7e7e7e7e7e00

	blankBoard := ^(board.Player | board.Opponent)

	var legalBoard BitBoard

	getNegativeStridedBoard := func(watchBoard BitBoard, shift int) BitBoard {
		nextOpponentBoard := watchBoard & (board.Player << shift)
		for i := 0; i < 5; i++ {
			nextOpponentBoard |= horizontalWatchBoard & (nextOpponentBoard << shift)
		}

		return blankBoard & (nextOpponentBoard << shift)
	}

	getPositiveStridedBoard := func(watchBoard BitBoard, shift int) BitBoard {
		nextOpponentBoard := watchBoard & (board.Player >> shift)
		for i := 0; i < 5; i++ {
			nextOpponentBoard |= horizontalWatchBoard & (nextOpponentBoard >> shift)
		}

		return blankBoard & (nextOpponentBoard >> shift)
	}

	// left
	legalBoard |= getNegativeStridedBoard(horizontalWatchBoard, 1)
	// left top
	legalBoard |= getNegativeStridedBoard(allSideWatchBoard, 9)
	// top
	legalBoard |= getNegativeStridedBoard(verticalWatchBoard, 8)
	// right top
	legalBoard |= getNegativeStridedBoard(allSideWatchBoard, 7)

	// right
	legalBoard |= getPositiveStridedBoard(horizontalWatchBoard, 1)
	// right bottom
	legalBoard |= getPositiveStridedBoard(allSideWatchBoard, 9)
	// bottom
	legalBoard |= getPositiveStridedBoard(verticalWatchBoard, 8)
	// left bottom
	legalBoard |= getPositiveStridedBoard(allSideWatchBoard, 7)

	return legalBoard
}

// CanPutPoint returns true if possible
func (board Board) CanPutPoint(x int, y int) bool {
	bb := CoordinateToBitBoard(x, y)

	return (bb & board.MakeLegalBoard()) == bb
}
