package moves

import "github.com/skjdfhkskjds/Chess/chess-logic/board"

func GenerateKnightAttacks(square int) *board.Bitboards {
	attackBoard := board.NewBitboard(0)
	pieceBoard := board.NewBitboard(0)
	pieceBoard.SetBit(square)

	// 17, 15, 10, 6 bits in each direction
	// Upwards 4 moves (top left, top right, top-middle left, top-middle right)
	if ((pieceBoard.Board >> 17) & board.NOT_H_FILE) != 0 {
		attackBoard.Board |= (pieceBoard.Board >> 17)
	}
	if ((pieceBoard.Board >> 15) & board.NOT_A_FILE) != 0 {
		attackBoard.Board |= (pieceBoard.Board >> 15)
	}
	if ((pieceBoard.Board >> 10) & board.NOT_GH_FILE) != 0 {
		attackBoard.Board |= (pieceBoard.Board >> 10)
	}
	if ((pieceBoard.Board >> 6) & board.NOT_AB_FILE) != 0 {
		attackBoard.Board |= (pieceBoard.Board >> 6)
	}

	// Bottom 4 moves (bottom left, bottom right, bottom-middle left, bottom-middle right)
	if ((pieceBoard.Board << 17) & board.NOT_A_FILE) != 0 {
		attackBoard.Board |= (pieceBoard.Board << 17)
	}
	if ((pieceBoard.Board << 15) & board.NOT_H_FILE) != 0 {
		attackBoard.Board |= (pieceBoard.Board << 15)
	}
	if ((pieceBoard.Board << 10) & board.NOT_AB_FILE) != 0 {
		attackBoard.Board |= (pieceBoard.Board << 10)
	}
	if ((pieceBoard.Board << 6) & board.NOT_GH_FILE) != 0 {
		attackBoard.Board |= (pieceBoard.Board << 6)
	}

	attackBoard.Piece = board.KNIGHT
	return attackBoard
}
