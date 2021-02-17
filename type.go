package othellogo

// BitBoard is a bit representation of the board
type BitBoard uint64

// Board represents the board of the player and the opponent
type Board struct {
	Player   BitBoard
	Opponent BitBoard
}
