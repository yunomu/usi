package sfen

import (
	"unicode"
)

var sfenPiece = " KRBGSNLP"

type Player int

const (
	Player_NULL Player = iota
	Player_BLACK
	Player_WHITE
)

func (p Player) flip() Player {
	switch p {
	case Player_BLACK:
		return Player_WHITE
	case Player_WHITE:
		return Player_BLACK
	default:
		return Player_NULL
	}
}

type PieceType int

const (
	Piece_NULL PieceType = iota
	Piece_GYOKU
	Piece_HISHA
	Piece_KAKU
	Piece_KIN
	Piece_GIN
	Piece_KEI
	Piece_KYOU
	Piece_FU
)

type Piece struct {
	Player   Player
	Type     PieceType
	Promoted bool
}

type pieceSlice []*Piece

func (p pieceSlice) Len() int { return len(p) }

func (p pieceSlice) Less(i int, j int) bool {
	switch {
	case p[i] == nil || p[i].Type == Piece_NULL || p[i].Player == Player_NULL:
		return true
	case p[j] == nil || p[j].Type == Piece_NULL || p[j].Player == Player_NULL:
		return false
	case p[i].Player != p[j].Player:
		return p[i].Player == Player_BLACK
	case p[i].Type != p[j].Type:
		return p[i].Type < p[j].Type
	default:
		return true
	}
}

func (p pieceSlice) Swap(i int, j int) { p[i], p[j] = p[j], p[i] }

func (p *Piece) isEmpty() bool {
	return p == nil || p.Player == Player_NULL || p.Type == Piece_NULL
}

func (p *Piece) equal(o *Piece) bool {
	if p == nil {
		return o == nil
	}
	return p.Player == o.Player && p.Type == o.Type
}

func (p *Piece) a() rune {
	r := []rune(sfenPiece)[p.Type]
	if p.Player == Player_WHITE {
		r = unicode.ToLower(r)
	}
	return r
}

func (p *Piece) aa() string {
	if p.isEmpty() {
		return " ."
	}

	d := ' '
	if p.Promoted {
		d = '+'
	}

	return string([]rune{d, p.a()})
}

func (p *Piece) sfen() string {
	if p.isEmpty() {
		return ""
	}

	var ret []rune
	if p.Promoted {
		ret = append(ret, '+')
	}

	return string(append(ret, p.a()))
}
