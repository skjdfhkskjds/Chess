package main

import (
	"github.com/skjdfhkskjds/Chess/chess-logic/board"
	"github.com/skjdfhkskjds/Chess/chess-logic/moves"
)

func main() {
	// masks := moves.InitAttackMasks()

	// Pawn Attack Masks
	// for _, side := range masks.PawnAttacks {
	// 	for _, square := range side {
	// 		square.Print_Bitboard()
	// 	}
	// }

	// Knight Attack Masks
	// for _, square := range masks.KnightAttacks {
	// 	square.Print_Bitboard()
	// }

	// testing
	// for rank := 0; rank < 8; rank++ {
	// 	for file := 0; file < 8; file++ {
	// 		square := rank*8 + file
	// 		if file > 1 {
	// 			bitboard.SetBit(square)
	// 		}
	// 	}
	// }
	// bitboard.SetBit(board.H4)
	// bitboard.PopBit(board.E2)

	bitboard := moves.GenerateRookAttacks(board.D4)
	bitboard.Print_Bitboard()

	// for n := 8; n >= 1; n-- {
	// 	fmt.Printf("A%d; B%d; C%d; D%d; E%d; F%d; G%d; H%d;\n", n, n, n, n, n, n, n, n)
	// }
}
