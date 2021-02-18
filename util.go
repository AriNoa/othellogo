package othellogo

// coordinateToBitBoard returns BitBoard that is flagged only at the specified coordinates
func coordinateToBitBoard(x int, y int) BitBoard {
	var bb BitBoard = 0x8000000000000000

	bb = bb >> x
	bb = bb >> (y * 8)

	return bb
}
