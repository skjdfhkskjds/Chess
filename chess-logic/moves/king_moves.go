package moves

import "github.com/skjdfhkskjds/Chess/chess-logic/board"

func GenerateKingAttacks(square int) *board.Bitboards {
	attackBoard := board.NewBitboard()
	pieceBoard := board.NewBitboard()
	pieceBoard.SetBit(square)

	// Top Row
	if (pieceBoard.Board >> 8) != 0 {
		attackBoard.Board |= (pieceBoard.Board >> 8)
	}
	if ((pieceBoard.Board >> 7) & board.NOT_A_FILE) != 0 {
		attackBoard.Board |= (pieceBoard.Board >> 7)
	}
	if ((pieceBoard.Board >> 9) & board.NOT_H_FILE) != 0 {
		attackBoard.Board |= (pieceBoard.Board >> 9)
	}

	// Bottom Row
	if (pieceBoard.Board << 8) != 0 {
		attackBoard.Board |= (pieceBoard.Board << 8)
	}
	if ((pieceBoard.Board << 7) & board.NOT_H_FILE) != 0 {
		attackBoard.Board |= (pieceBoard.Board << 7)
	}
	if ((pieceBoard.Board << 9) & board.NOT_A_FILE) != 0 {
		attackBoard.Board |= (pieceBoard.Board << 9)
	}

	// Middle Left, Middle Right
	if ((pieceBoard.Board << 1) & board.NOT_A_FILE) != 0 {
		attackBoard.Board |= (pieceBoard.Board << 1)
	}
	if ((pieceBoard.Board >> 1) & board.NOT_H_FILE) != 0 {
		attackBoard.Board |= (pieceBoard.Board >> 1)
	}

	attackBoard.Piece = board.KING
	return attackBoard
}
