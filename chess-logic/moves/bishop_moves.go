package moves

import "github.com/skjdfhkskjds/Chess/chess-logic/board"

func GenerateBishopMasks(square int) *board.Bitboards {
	attackBoard := board.NewBitboard(0)
	pieceBoard := board.NewBitboard(0)
	pieceBoard.SetBit(square)

	target_rank := square / 8
	target_file := square % 8

	// Diagonal Down and Right
	for rank, file := target_rank+1, target_file+1; rank <= 6 && file <= 6; rank, file = rank+1, file+1 {
		attackBoard.Board |= uint64(1) << (rank*8 + file)
	}
	// Diagonal Up and Right
	for rank, file := target_rank-1, target_file+1; rank >= 1 && file <= 6; rank, file = rank-1, file+1 {
		attackBoard.Board |= uint64(1) << (rank*8 + file)
	}
	// Diagonal Up and Left
	for rank, file := target_rank+1, target_file-1; rank <= 6 && file >= 1; rank, file = rank+1, file-1 {
		attackBoard.Board |= uint64(1) << (rank*8 + file)
	}
	// Diagonal Down and Left
	for rank, file := target_rank-1, target_file-1; rank >= 1 && file > 0; rank, file = rank-1, file-1 {
		attackBoard.Board |= uint64(1) << (rank*8 + file)
	}

	attackBoard.Piece = board.BISHOP
	return attackBoard
}

func GenerateBishopAttacks(square int, b board.Bitboards) *board.Bitboards {
	attackBoard := board.NewBitboard(0)
	pieceBoard := board.NewBitboard(0)
	pieceBoard.SetBit(square)

	target_rank := square / 8
	target_file := square % 8

	// Diagonal Down and Right
	for rank, file := target_rank+1, target_file+1; rank <= 7 && file <= 7; rank, file = rank+1, file+1 {
		attackBoard.Board |= uint64(1) << (rank*8 + file)
		if ((uint64(1) << (rank*8 + file)) & b.Board) != 0 {
			break
		}
	}
	// Diagonal Up and Right
	for rank, file := target_rank-1, target_file+1; rank >= 0 && file <= 7; rank, file = rank-1, file+1 {
		attackBoard.Board |= uint64(1) << (rank*8 + file)
		if ((uint64(1) << (rank*8 + file)) & b.Board) != 0 {
			break
		}
	}
	// Diagonal Up and Left
	for rank, file := target_rank-1, target_file-1; rank >= 0 && file >= 0; rank, file = rank-1, file-1 {
		attackBoard.Board |= uint64(1) << (rank*8 + file)
		if ((uint64(1) << (rank*8 + file)) & b.Board) != 0 {
			break
		}
	}
	// Diagonal Down and Left
	for rank, file := target_rank+1, target_file-1; rank <= 7 && file >= 0; rank, file = rank+1, file-1 {
		attackBoard.Board |= uint64(1) << (rank*8 + file)
		if ((uint64(1) << (rank*8 + file)) & b.Board) != 0 {
			break
		}
	}

	attackBoard.Piece = board.BISHOP
	return attackBoard
}
