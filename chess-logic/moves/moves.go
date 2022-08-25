package moves

import "github.com/skjdfhkskjds/Chess/chess-logic/board"

type AttackMasks struct {
	PawnAttacks   [2][64]*board.Bitboards
	KnightAttacks [64]*board.Bitboards
	KingAttacks   [64]*board.Bitboards
}

func InitAttackMasks() *AttackMasks {
	var attacks AttackMasks
	for square := 0; square < 64; square++ {
		attacks.PawnAttacks[board.WHITE][square] = GeneratePawnAttacks(board.WHITE, square)
		attacks.PawnAttacks[board.BLACK][square] = GeneratePawnAttacks(board.BLACK, square)

		attacks.KnightAttacks[square] = GenerateKnightAttacks(square)

		attacks.KingAttacks[square] = GenerateKingAttacks(square)
	}
	return &attacks
}
