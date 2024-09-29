package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Board struct {
	WhitePawns   uint64
	WhiteKnights uint64
	WhiteBishops uint64
	WhiteRooks   uint64
	WhiteQueens  uint64
	WhiteKing    uint64

	BlackPawns   uint64
	BlackKnights uint64
	BlackBishops uint64
	BlackRooks   uint64
	BlackQueens  uint64
	BlackKing    uint64
}

func (b *Board) LoadFromFen(fen string) {
	rows := strings.Split(fen, "/")
	if len(rows) != 8 {
		fmt.Printf("Invalid FEN format\n")
	}

	for i, row := range rows {
		file := 0
		for _, piece := range row {
			if unicode.IsDigit(piece) {
				file += int(piece - '0')
			} else {
				b.InsertPiece(piece, 7-i, file)
				file += 1
			}
		}
	}
}

// InsertPiece sets a piece on the board in the corresponding rank and file.
// piece: int32 representing the type of piece (e.g., 'r' for black rook).
// rank: 0-7, where 0 is the 1st rank (bottom) and 7 is the 8th rank (top).
// file: 0-7, where 0 is 'a' and 7 is 'h'.
func (b *Board) InsertPiece(piece int32, rank int, file int) {
	index := (rank * 8) + file
	switch piece {
	case 'r':
		b.BlackRooks |= 1 << index
	case 'n':
		b.BlackKnights |= 1 << index
	case 'b':
		b.BlackBishops |= 1 << index
	case 'q':
		b.BlackQueens |= 1 << index
	case 'k':
		b.BlackKing |= 1 << index
	case 'p':
		b.BlackPawns |= 1 << index
	case 'R':
		b.WhiteRooks |= 1 << index
	case 'N':
		b.WhiteKnights |= 1 << index
	case 'B':
		b.WhiteBishops |= 1 << index
	case 'Q':
		b.WhiteQueens |= 1 << index
	case 'K':
		b.WhiteKing |= 1 << index
	case 'P':
		b.WhitePawns |= 1 << index
	}

}

func PrintBitboard(bitboard uint64) {
	for rank := 7; rank >= 0; rank-- { // Start from rank 7 (8th rank) down to rank 0 (1st rank)
		for file := 0; file < 8; file++ {
			// Calculate the bit index (rank*8 + file) and check if it's set
			index := rank*8 + file
			if (bitboard & (1 << index)) != 0 {
				fmt.Print("1 ") // There is a piece
			} else {
				fmt.Print("0 ") // Empty square
			}
		}
		fmt.Println() // Newline after each rank
	}
}

// GetPieceAt returns the character representing the piece at the given index.
func (b *Board) GetPieceAt(index int) string {
	if b.WhitePawns&(1<<index) != 0 {
		return "P"
	}
	if b.WhiteKnights&(1<<index) != 0 {
		return "N"
	}
	if b.WhiteBishops&(1<<index) != 0 {
		return "B"
	}
	if b.WhiteRooks&(1<<index) != 0 {
		return "R"
	}
	if b.WhiteQueens&(1<<index) != 0 {
		return "Q"
	}
	if b.WhiteKing&(1<<index) != 0 {
		return "K"
	}
	if b.BlackPawns&(1<<index) != 0 {
		return "p"
	}
	if b.BlackKnights&(1<<index) != 0 {
		return "n"
	}
	if b.BlackBishops&(1<<index) != 0 {
		return "b"
	}
	if b.BlackRooks&(1<<index) != 0 {
		return "r"
	}
	if b.BlackQueens&(1<<index) != 0 {
		return "q"
	}
	if b.BlackKing&(1<<index) != 0 {
		return "k"
	}
	return "." // Empty square
}

func (b *Board) PrintBoard() {
	for rank := 7; rank >= 0; rank-- { // From rank 7 (8th rank) to rank 0 (1st rank)
		for file := 0; file < 8; file++ {
			index := rank*8 + file // Calculate the bit index
			fmt.Print(b.GetPieceAt(index), " ")
		}
		fmt.Println() // Newline after each rank
	}
}
