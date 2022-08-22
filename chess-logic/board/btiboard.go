package board

import (
	"fmt"
)

type BitBoard interface {
	GetBit(i int) bool
	Print_BitBoard(board uint64)
}

type BitBoards struct {
	board uint64
	piece byte
}

func (board BitBoards) GetBit(i int) bool {
	return (board.board & (1 << i)) != 0
}

func (board BitBoards) Print_BitBoard() {
	print("\n")
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			square := rank*8 + file
			if file == 0 {
				print("  %d ", 8-rank)
			}
			if board.GetBit(square) {
				print(" %d", 1)
			} else {
				print(" %d", 0)
			}
		}
		print("\n")
	}
}
