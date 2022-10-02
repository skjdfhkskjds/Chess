package board

import "fmt"

// Game Functions
func (game *Game) PrintGame() {
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			square := rank*8 + file
			if file == 0 {
				fmt.Printf("  %d ", 8-rank)
			}
			piece := -1
			for i, board := range game.Boards {
				if board.GetBit(square) {
					piece = i
				}
			}
			if piece == -1 {
				print(" . ")
			} else {
				// fmt.Printf(" %s\n", ASCII_PIECES[piece])
				fmt.Printf(" %s ", UNICODE_PIECES[piece])
			}
		}
		print("\n")
	}
	print("\n     a  b  c  d  e  f  g  h\n\n\n")
	side := "White"
	if game.Side == 1 {
		side = "Black"
	}
	enpassant := "N/A"
	if game.Enpassant != no_square {
		enpassant = IndexesToCoordinates[game.Enpassant]
	}
	castling := "----"
	if (game.Castle & Wk) != 0 {
		castling = castling[:0] + "K" + castling[1:]
	}
	if (game.Castle & Wq) != 0 {
		castling = castling[:1] + "Q" + castling[2:]
	}
	if (game.Castle & Bk) != 0 {
		castling = castling[:2] + "k" + castling[3:]
	}
	if (game.Castle & Bq) != 0 {
		castling = castling[:3] + "q" + castling[4:]
	}
	print("    Side to Move:     " + side + "\n")
	print("    Enpassant Square: " + enpassant + "\n")
	print("    Castling Rights:  " + castling + "\n")
}

func (game *Game) SetupBitboards(start bool) {
	startBoards := [12]uint64{}
	if start {
		startBoards = GameStartBoards
	}
	for i := range game.Boards {
		game.Boards[i] = NewBitboard(startBoards[i])
	}
}

func (game *Game) SetupOccupancies(start bool) {
	startBoards := [3]uint64{}
	if start {
		startBoards = GameStartOccupancies
	}
	for i := range game.Occupancies {
		game.Occupancies[i] = NewBitboard(startBoards[i])
	}
}

func NewGame() *Game {
	var new Game
	new.SetupBitboards(true)
	new.SetupOccupancies(true)
	new.Side = WHITE
	new.Enpassant = no_square
	new.Castle = 0
	new.Castle |= Wk
	new.Castle |= Wq
	new.Castle |= Bk
	new.Castle |= Bq
	return &new
}

// FEN String Parsing
func ParseFEN(fen string) *Game {
	var game Game
	game.SetupBitboards(false)
	game.SetupOccupancies(false)
	count := 0
	// parsing pieces of FEN String
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			square := rank*8 + file
			if (fen[count] >= 'a' && fen[count] <= 'z') || (fen[count] >= 'A' && fen[count] <= 'Z') {
				piece := ASCII_TO_CONSTANT[fen[count]]
				game.Boards[piece].SetBit(square)
				count++
			}
			if fen[count] >= '0' && fen[count] <= '9' {
				offset := int(fen[count] - '0')
				piece := -1
				for i, board := range game.Boards {
					if board.GetBit(square) {
						piece = i
					}
				}
				if piece == -1 {
					file--
				}
				file += offset
				count++
			}
			if fen[count] == '/' {
				count++
			}
		}
	}

	// parsing game state variables
	count++
	if fen[count] == 'w' {
		game.Side = WHITE
	} else {
		game.Side = BLACK
	}
	// skips next whitespace to castling rights
	count += 2

	// setup castling rights
	for fen[count] != ' ' {
		switch fen[count] {
		case 'K':
			game.Castle |= Wk
		case 'Q':
			game.Castle |= Wq
		case 'k':
			game.Castle |= Bk
		case 'q':
			game.Castle |= Bq
		}
		count++
	}
	count++

	// setup enpassant square
	if fen[count] != '-' {
		file := int(fen[count] - 'a')
		rank := 8 - int(fen[count+1]-'0')
		game.Enpassant = rank*8 + file
	} else {
		game.Enpassant = no_square
	}

	// setup occupancy bitboards
	for piece := WP; piece <= WK; piece++ {
		game.Occupancies[WHITE].Board |= game.Boards[piece].Board
	}
	for piece := BP; piece <= BK; piece++ {
		game.Occupancies[BLACK].Board |= game.Boards[piece].Board
	}
	game.Occupancies[BOTH].Board |= (game.Occupancies[WHITE].Board | game.Occupancies[BLACK].Board)
	return &game
}
