package moves

import "github.com/skjdfhkskjds/Chess/board"

type Pawn_Attacks struct {
	attacks [2][64]uint64
}

func GeneratePawnAttacks(square, side int) *board.Bitboards {
	// attackBoard := board.NewBitboard()
	pieceBoard := board.NewBitboard()
	pieceBoard.SetBit(square)

	// white team
	if side == 0 {

	} else {

	}
	return pieceBoard
}
