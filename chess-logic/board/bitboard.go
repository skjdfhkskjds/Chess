package board

import (
	"fmt"
)

type Bitboards struct {
	Board uint64
	Piece int
}

// Initializes Bitboard
func NewBitboard() *Bitboards {
	return &Bitboards{
		Board: 0,
	}
}

// Get, Set, Pop Operators

func (board *Bitboards) GetBit(square int) bool {
	return (board.Board & (uint64(1) << square)) != 0
}

func (board *Bitboards) SetBit(square int) {
	board.Board |= (uint64(1) << square)
}

func (board *Bitboards) PopBit(square int) {
	if board.GetBit(square) {
		board.Board ^= (uint64(1) << square)
	}
}

// Complex Operations
func (board *Bitboards) GetLeastSignificantFirstBitIndex() int {
	b := board.Board
	const debruijn64 uint64 = 0x03f79d71b4cb0a89
	return Index64[((b&-b)*debruijn64)>>58]
}

func (board *Bitboards) SetOccupancies(index int, bits_in_mask int) *Bitboards {
	occupancy_board := NewBitboard()
	new := NewBitboard()
	new.Board = board.Board
	for count := 0; count < bits_in_mask; count++ {
		square := new.GetLeastSignificantFirstBitIndex()
		new.PopBit(square)
		if index&(1<<count) != 0 {
			occupancy_board.Board |= (uint64(1) << square)
		}
	}
	return occupancy_board
}

func (board *Bitboards) PrintBitboard() {
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

	fmt.Printf("    Bitboard: %d\n\n", board.Board)
}
