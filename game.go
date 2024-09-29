package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Game struct {
	Board           *Board
	WhoseMove       int    // 0 for white, 1 for black
	CastlingRights  string // "KQkq" format for castling rights
	EnPassantSquare string // Square where en passant is possible, e.g., "e3", or "-" if none
	HalfmoveClock   int    // Number of halfmoves since the last capture or pawn advance (for the fifty-move rule)
	FullmoveCounter int    // Number of full moves (increments after Black's move)
}

//func (g *Game) canCastle() bool {
//	if g.CastlingRights == "-" {
//		return false
//	} else {
//		if g.CastlingRights {
//			return true
//		}
//		return false
//	}
//}

func (g *Game) LoadFromFen(fen string) {
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
				g.Board.InsertPiece(piece, 7-i, file)
				file += 1
			}
		}
	}

	g.ReadMetadata(rows[7])
}

func (g *Game) ReadMetadata(metadata string) {
	data := strings.Split(metadata, " ")

	if len(data) != 6 {
		fmt.Printf("Invalid FEN format\n")
	}

	if data[1] == "w" {
		g.WhoseMove = 0
	} else {
		g.WhoseMove = 1
	}

	g.CastlingRights = data[2]
	g.EnPassantSquare = data[3]
	g.HalfmoveClock, _ = strconv.Atoi(data[4])
	g.FullmoveCounter, _ = strconv.Atoi(data[5])
}

func NewBoard() *Board {
	return &Board{
		WhitePawns:   0,
		WhiteKnights: 0,
		WhiteBishops: 0,
		WhiteRooks:   0,
		WhiteQueens:  0,
		WhiteKing:    0,
		BlackPawns:   0,
		BlackKnights: 0,
		BlackBishops: 0,
		BlackRooks:   0,
		BlackQueens:  0,
		BlackKing:    0,
	}
}

func SetUpGame() *Game {
	board := NewBoard()
	return &Game{
		Board:           board,
		WhoseMove:       0,      // 0 for white's turn
		CastlingRights:  "KQkq", // Both sides can castle both ways
		EnPassantSquare: "-",    // No en passant square initially
		HalfmoveClock:   0,      // Halfmove clock starts at 0
		FullmoveCounter: 1,      // Fullmove counter starts at 1
	}
}

func (g *Game) PrintState() {
	g.Board.PrintBoard()
	println()
	fmt.Printf("Whose Move: %d\nCastling Rights: %s\nEn Passant Square: %s\nHalfmove Clock:%d\nFullmove Clock:%d\n", g.WhoseMove, g.CastlingRights, g.EnPassantSquare, g.HalfmoveClock, g.FullmoveCounter)
}

func (g *Game) makeMove() {
	isLegalMove()
}
