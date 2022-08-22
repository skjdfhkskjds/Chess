package main

import (
	"github.com/skjdfhkskjds/Chess/board"
)

func main() {
	bitboard := *board.NewBitboard(uint64(0))

	// testing
	bitboard.SetBit(board.E2)
	bitboard.PopBit(board.E2)
	bitboard.Print_Bitboard()

	// for n := 8; n >= 1; n-- {
	// 	fmt.Printf("A%d; B%d; C%d; D%d; E%d; F%d; G%d; H%d;\n", n, n, n, n, n, n, n, n)
	// }
}
