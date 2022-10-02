package moves

import "github.com/skjdfhkskjds/Chess/chess-logic/board"

// IMPLEMENT EN PASSANT
func GeneratePawnAttacks(side, square int) *board.Bitboards {
	attackBoard := board.NewBitboard(0)
	pieceBoard := board.NewBitboard(0)
	pieceBoard.SetBit(square)

	// white team
	if side == 0 {
		// checks whether the "attacked square" is out of bounds
		if ((pieceBoard.Board >> 7) & board.NOT_A_FILE) != 0 {
			attackBoard.Board |= (pieceBoard.Board >> 7)
		}
		if ((pieceBoard.Board >> 9) & board.NOT_H_FILE) != 0 {
			attackBoard.Board |= (pieceBoard.Board >> 9)
		}
	} else { // black team
		// checks whether the "attacked square" is out of bounds
		if ((pieceBoard.Board << 7) & board.NOT_H_FILE) != 0 {
			attackBoard.Board |= (pieceBoard.Board << 7)
		}
		if ((pieceBoard.Board << 9) & board.NOT_A_FILE) != 0 {
			attackBoard.Board |= (pieceBoard.Board << 9)
		}
	}
	attackBoard.Piece = board.PAWN
	return attackBoard
}
