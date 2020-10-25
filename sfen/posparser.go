package sfen

import (
	"fmt"
	"io"
	"strings"
	"unicode"
)

type posParser struct {
	r io.RuneScanner
}

func newPosParser(s string) *posParser {
	return &posParser{
		r: strings.NewReader(s),
	}
}

func (p *posParser) read() (rune, error) {
	r, _, err := p.r.ReadRune()
	return r, err
}

var sfenPieceBW = sfenPiece + strings.ToLower(sfenPiece)

func runeToPiece(r rune) (Player, PieceType) {
	idx := strings.IndexRune(sfenPieceBW, r)
	if idx == -1 {
		return Player_NULL, Piece_NULL
	}

	var pl Player
	switch idx / len(sfenPiece) {
	case 0:
		pl = Player_BLACK
	case 1:
		pl = Player_WHITE
	}

	return pl, PieceType(idx % len(sfenPiece))
}

func (p *posParser) parseBoard() ([9][9]*Piece, error) {
	var board [9][9]*Piece

	for _, y := range PosYs {
		var promoted bool
		var skip int
		for _, x := range PosXs {
			if skip > 0 {
				skip--
				continue
			}

		redo:
			r, err := p.read()
			if err != nil {
				return board, err
			}

			p := &Piece{}
			if r == '/' {
				goto redo
			} else if r == '+' {
				promoted = true
				goto redo
			} else if unicode.IsDigit(r) {
				skip = int(r-'0') - 1
				continue
			} else {
				pl, pi := runeToPiece(r)
				if pl == Player_NULL || pi == Piece_NULL {
					return board, fmt.Errorf("unknown Piece: %c", r)
				}

				p.Player = pl
				p.Type = pi
				p.Promoted = promoted
			}
			board[y][x] = p

			promoted = false
		}
	}

	if _, err := p.read(); err == nil {
		return board, fmt.Errorf("remain")
	}

	return board, nil
}

func parseCaptured(s string) []*Piece {
	var ret []*Piece
	if s == "-" {
		return ret
	}

	count := 0
	for _, r := range []rune(s) {
		if unicode.IsDigit(r) {
			if count != 0 {
				count *= 10
			}
			count += int(r - '0')
			continue
		}

		pl, pi := runeToPiece(r)
		if pl == Player_NULL || pi == Piece_NULL {
			return nil
		}

		if count == 0 {
			count = 1
		}
		for i := 0; i < count; i++ {
			ret = append(ret, &Piece{Player: pl, Type: pi})
		}
		count = 0
	}

	return ret
}
