package board

import (
	"fmt"
)

// type Bitboard interface {
// 	GetBit(i int) bool
// 	Print_Bitboard(board uint64)
// }

type Bitboards struct {
	board uint64
	// piece byte
}

// Initializes Bitboard

func NewBitboard() *Bitboards {
	return &Bitboards{
		board: 0,
	}
}

// Get, Set, Pop Operators

func (board *Bitboards) GetBit(square int) bool {
	return (board.board & (uint64(1) << square)) != 0
}

func (board *Bitboards) SetBit(square int) {
	board.board |= (uint64(1) << square)
}

func (board *Bitboards) PopBit(square int) {
	if board.GetBit(square) {
		board.board ^= (uint64(1) << square)
	}
}

// Complex Operations

func (board *Bitboards) Print_Bitboard() {
	print("\n")
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			square := rank*8 + file
			if file == 0 {
				fmt.Printf("  %d ", 8-rank)
			}
			if board.GetBit(square) {
				fmt.Printf(" %d", 1)
			} else {
				fmt.Printf(" %d", 0)
			}
		}
		print("\n")
	}
	print("\n     a b c d e f g h\n\n")

	fmt.Printf("    Bitboard: %d\n\n", board.board)
}
