package main

import (
	"github.com/skjdfhkskjds/Chess/chess-logic/board"
	"github.com/skjdfhkskjds/Chess/chess-logic/moves"
)

func main() {
	// INITIALIZE ROUTINE
	attacks := moves.InitAttackMasks()

	// game := board.NewGame()
	// game := board.ParseFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1 ")

	// game.PrintGame()
	// game.Occupancies[board.WHITE].PrintBitboard()
	// game.Occupancies[board.BLACK].PrintBitboard()
	// game.Occupancies[board.BOTH].PrintBitboard()

	bitboard := board.NewBitboard(0)
	bitboard.SetBit(board.B6)
	bitboard.SetBit(board.D6)
	bitboard.SetBit(board.F6)
	bitboard.SetBit(board.G4)
	bitboard.SetBit(board.B4)
	bitboard.SetBit(board.B2)
	bitboard.SetBit(board.D2)
	bitboard.SetBit(board.F2)
	attack := attacks.GetQueenAttacks(board.D4, bitboard.Board)
	// // // bitboard.SetBit(board.D7)
	// // // bitboard.SetBit(board.E7)
	// // // bitboard.SetBit(board.F7)
	// // // bitboard.SetBit(board.G7)
	// // // bitboard.SetBit(board.H7)
	attack.PrintBitboard()

	// fmt.Printf(" %s\n", board.ASCII_PIECES[board.WP])
	// fmt.Printf(" %s\n", board.UNICODE_PIECES[board.WP])
	// print("Rooks: \n")
	// for square := 0; square < 64; square++ {
	// 	fmt.Printf("%x\n", board.RookMagicNumbers[square])
	// }
	// print("Bishops: \n")
	// for square := 0; square < 64; square++ {
	// 	fmt.Printf("%x\n", board.BishopMagicNumbers[square])
	// }

	// engine.InitMagicNumbers()
	// num :=
	// bitboard := board.NewBitboard(0)
	// bitboard.SetBit(board.C5)
	// bitboard.SetBit(board.F2)
	// bitboard.SetBit(board.G7)
	// bitboard.SetBit(board.B2)
	// bitboard.SetBit(board.G5)
	// bitboard.SetBit(board.E2)
	// bitboard.SetBit(board.E7)
	// bitboard.PrintBitboard()
	// bishopTest := attacks.GetBishopAttacks(board.D4, bitboard.Board)
	// bishopTest.PrintBitboard()
	// rookTest := attacks.GetRookAttacks(board.E5, bitboard.Board)
	// rookTest.PrintBitboard()

	// fmt.Printf("%d\n", uint64(engine.GetRandomU32()))
	// fmt.Printf("%d\n", uint64(engine.GetRandomU32()&0xFFFF))
	// fmt.Printf("%d\n", engine.GetRandomU32())
	// fmt.Printf("%d\n", engine.GetRandomU32())
	// fmt.Printf("%d\n", engine.GetRandomU32())
	// bitboard2 := board.NewBitboard(0)
	// bitboard2.Board = uint64(engine.GetRandomU32() & 0xFFFF)
	// bitboard2.PrintBitboard()

	// bitboard3 := board.NewBitboard(0)
	// bitboard3.Board = engine.GetRandomU64()
	// bitboard3.PrintBitboard()

	// bitboard5 := board.NewBitboard(0)
	// bitboard5.Board = engine.GetRandomU64() & engine.GetRandomU64() & engine.GetRandomU64()
	// bitboard5.PrintBitboard()

	// bitboard4 := board.NewBitboard(0)
	// bitboard4.Board = engine.GenerateMagicNumberCandidate()
	// bitboard4.PrintBitboard()
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
	// 		bitboard := moves.GenerateRookMasks(square)
	// 		fmt.Printf(" %d,", bitboard.CountBits())
	// 	}
	// 	print("\n")GenerateMagicNumbers
	// }
	// bitboard.SetBit(board.H4)
	// bitboard.PopBit(board.E2)

	// attacks := moves.GenerateRookMasks(board.A1)
	// bitboard := board.NewBitboard(0)
	// bitboard.SetBit(board.D1)
	// bitboard.SetBit(board.B4)
	// bitboard.SetBit(board.G4)
	// print(board.CountBits(bitboard.Board), "\n")
	// print(bits.OnesCount64(bitboard.Board))
	// bitboard.PrintBitboard()

	// for n := 8; n >= 1; n-- {
	// 	fmt.Printf("A%d; B%d; C%d; D%d; E%d; F%d; G%d; H%d;\n", n, n, n, n, n, n, n, n)
	// }
}
