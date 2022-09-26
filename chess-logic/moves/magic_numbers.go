package moves

import (
	"fmt"
	"math/bits"

	"github.com/skjdfhkskjds/Chess/chess-logic/board"
)

var random_seed uint32 = 1804289383

func GetRandomU32() uint32 {
	// get current state
	number := random_seed
	// XOR shift algorithm
	number ^= number << 13
	number ^= number >> 17
	number ^= number << 5

	random_seed = number
	return number
}

func GetRandomU64() uint64 {
	// init random u64 values by slicing 16 bits from the MS1B side
	n1 := uint64(GetRandomU32() & 0xFFFF)
	n2 := uint64(GetRandomU32() & 0xFFFF)
	n3 := uint64(GetRandomU32() & 0xFFFF)
	n4 := uint64(GetRandomU32() & 0xFFFF)

	rand := n1 | (n2 << 16) | (n3 << 32) | (n4 << 48)
	return rand
}

func GenerateMagicNumberCandidate() uint64 {
	rand1 := GetRandomU64()
	rand2 := GetRandomU64()
	rand3 := GetRandomU64()
	return rand1 & rand2 & rand3
}

func GenerateMagicNumbers(square int, relevant_bits int, bishop bool) uint64 {
	var occupancies [4096]*board.Bitboards
	var attacks [4096]uint64
	var used_attacks [4096]uint64
	var attack_mask *board.Bitboards

	if bishop {
		attack_mask = GenerateBishopMasks(square)
	} else {
		attack_mask = GenerateRookMasks(square)
	}
	occupancy_indices := 1 << relevant_bits

	for index := 0; index < occupancy_indices; index++ {
		occupancies[index] = attack_mask.SetOccupancies(index, relevant_bits)
		if bishop {
			attacks[index] = GenerateBishopAttacks(square, *occupancies[index]).Board
		} else {
			attacks[index] = GenerateRookAttacks(square, *occupancies[index]).Board
		}
	}

	// testing magic numbers
	for random_count := 0; random_count < 100000000; random_count++ {
		magic_number := GenerateMagicNumberCandidate()
		if bits.OnesCount64((attack_mask.Board*magic_number)&0xFF00000000000000) < 6 {
			continue
		}
		used_attacks = [len(used_attacks)]uint64{}
		var index int
		var fail bool

		for index, fail = 0, false; !fail && index < occupancy_indices; index++ {
			magic_index := int(((occupancies[index].Board * magic_number) & 0xffffffffffffffff) >> (64 - relevant_bits))
			if used_attacks[magic_index] == uint64(0) {
				used_attacks[magic_index] = attacks[index]
			} else if used_attacks[magic_index] != attacks[index] {
				fail = true
			}
		}

		if !fail {
			return magic_number
		}
	}
	print("Magic search failed")
	return uint64(0)
}

func InitMagicNumbers() {
	print("Rooks: \n")
	for square := 0; square < 64; square++ {
		fmt.Printf("0x%x,\n", GenerateMagicNumbers(square, board.RelevantRookOccupancyBitCount[square], board.Rook))
	}
	print("Bishops: \n")
	for square := 0; square < 64; square++ {
		fmt.Printf("0x%x,\n", GenerateMagicNumbers(square, board.RelevantBishopOccupancyBitCount[square], board.Bishop))
	}
}
