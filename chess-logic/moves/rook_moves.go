package moves

import "github.com/skjdfhkskjds/Chess/chess-logic/board"

func GenerateRookMasks(square int) *board.Bitboards {
	attackBoard := board.NewBitboard()
	pieceBoard := board.NewBitboard()
	pieceBoard.SetBit(square)

	target_rank := square / 8
	target_file := square % 8

	for rank := target_rank + 1; rank < 7; rank++ {
		attackBoard.Board |= (uint64(1) << (rank*8 + target_file))
	}
	for rank := target_rank - 1; rank > 0; rank-- {
		attackBoard.Board |= (uint64(1) << (rank*8 + target_file))
	}

	for file := target_file + 1; file < 7; file++ {
		attackBoard.Board |= (uint64(1) << (target_rank*8 + file))
	}
	for file := target_file - 1; file > 0; file-- {
		attackBoard.Board |= (uint64(1) << (target_rank*8 + file))
	}

	attackBoard.Piece = board.ROOK
	return attackBoard
}

func GenerateRookAttacks(square int, b board.Bitboards) *board.Bitboards {
	attackBoard := board.NewBitboard()
	pieceBoard := board.NewBitboard()
	pieceBoard.SetBit(square)

	target_rank := square / 8
	target_file := square % 8

	for rank := target_rank + 1; rank <= 7; rank++ {
		attackBoard.Board |= (uint64(1) << (rank*8 + target_file))
		if ((uint64(1) << (rank*8 + target_file)) & b.Board) != 0 {
			break
		}
	}
	for rank := target_rank - 1; rank >= 0; rank-- {
		attackBoard.Board |= (uint64(1) << (rank*8 + target_file))
		if ((uint64(1) << (rank*8 + target_file)) & b.Board) != 0 {
			break
		}
	}

	for file := target_file + 1; file <= 7; file++ {
		attackBoard.Board |= (uint64(1) << (target_rank*8 + file))
		if ((uint64(1) << (target_rank*8 + file)) & b.Board) != 0 {
			break
		}
	}
	for file := target_file - 1; file >= 0; file-- {
		attackBoard.Board |= (uint64(1) << (target_rank*8 + file))
		if ((uint64(1) << (target_rank*8 + file)) & b.Board) != 0 {
			break
		}
	}

	attackBoard.Piece = board.ROOK
	return attackBoard
}
