package moves

import (
	"math/bits"

	"github.com/skjdfhkskjds/Chess/chess-logic/board"
)

type AttackMasks struct {
	PawnAttacks   [2][64]*board.Bitboards
	KnightAttacks [64]*board.Bitboards
	KingAttacks   [64]*board.Bitboards
	BishopMasks   [64]*board.Bitboards
	RookMasks     [64]*board.Bitboards
	BishopAttacks [64][512]*board.Bitboards
	RookAttacks   [64][4096]*board.Bitboards
}

func InitAttackMasks() *AttackMasks {
	var attacks AttackMasks
	for square := 0; square < 64; square++ {
		// initialize pawn attack tables
		attacks.PawnAttacks[board.WHITE][square] = GeneratePawnAttacks(board.WHITE, square)
		attacks.PawnAttacks[board.BLACK][square] = GeneratePawnAttacks(board.BLACK, square)

		// initialize knight attack tables
		attacks.KnightAttacks[square] = GenerateKnightAttacks(square)

		// initialize king attack tables
		attacks.KingAttacks[square] = GenerateKingAttacks(square)
	}
	attacks.InitSlidingPiecesAttacks(board.Bishop)
	attacks.InitSlidingPiecesAttacks(board.Rook)

	return &attacks
}

func (attacks *AttackMasks) InitSlidingPiecesAttacks(bishop bool) {
	for square := 0; square < 64; square++ {
		// initialize bishop masks
		attacks.BishopMasks[square] = GenerateBishopMasks(square)

		// initialize rook masks
		attacks.RookMasks[square] = GenerateRookMasks(square)

		// initialize sliding pieces attack tables
		var attack_mask *board.Bitboards
		if bishop {
			attack_mask = attacks.BishopMasks[square]
		} else {
			attack_mask = attacks.RookMasks[square]
		}

		relevant_bits_count := bits.OnesCount64(attack_mask.Board)
		occupancy_indices := (1 << relevant_bits_count)
		for index := 0; index < occupancy_indices; index++ {
			if bishop {
				state := attack_mask.SetOccupancies(index, relevant_bits_count)
				magic_index := (state.Board * board.BishopMagicNumbers[square]) >> (64 - board.RelevantBishopOccupancyBitCount[square])
				attacks.BishopAttacks[square][magic_index] = GenerateBishopAttacks(square, *state)
			} else {
				state := attack_mask.SetOccupancies(index, relevant_bits_count)
				magic_index := (state.Board * board.RookMagicNumbers[square]) >> (64 - board.RelevantRookOccupancyBitCount[square])
				attacks.RookAttacks[square][magic_index] = GenerateRookAttacks(square, *state)
			}
		}
	}
}

func (attacks *AttackMasks) GetBishopAttacks(square int, state uint64) *board.Bitboards {
	// get bishop attacks given current board state
	state &= attacks.BishopMasks[square].Board
	state *= board.BishopMagicNumbers[square]
	state >>= 64 - board.RelevantBishopOccupancyBitCount[square]
	return attacks.BishopAttacks[square][state]
}

func (attacks *AttackMasks) GetRookAttacks(square int, state uint64) *board.Bitboards {
	// get rook attacks given current board state
	state &= attacks.RookMasks[square].Board
	state *= board.RookMagicNumbers[square]
	state >>= 64 - board.RelevantRookOccupancyBitCount[square]
	return attacks.RookAttacks[square][state]
}

func (attacks *AttackMasks) GetQueenAttacks(square int, state uint64) *board.Bitboards {
	bishopState := state
	rookState := state

	// get bishop attacks given current board state
	bishopState &= attacks.BishopMasks[square].Board
	bishopState *= board.BishopMagicNumbers[square]
	bishopState >>= 64 - board.RelevantBishopOccupancyBitCount[square]

	attackBoard := attacks.BishopAttacks[square][bishopState]

	// get rook attacks given current board state
	rookState &= attacks.RookMasks[square].Board
	rookState *= board.RookMagicNumbers[square]
	rookState >>= 64 - board.RelevantRookOccupancyBitCount[square]
	attackBoard.Board |= attacks.RookAttacks[square][rookState].Board

	return attackBoard
}
