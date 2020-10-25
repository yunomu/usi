package sfen

import (
	"strings"
)

var (
	posX = "987654321"
	posY = "abcdefghi"
)

type PosX int

const (
	X_9 PosX = iota
	X_8
	X_7
	X_6
	X_5
	X_4
	X_3
	X_2
	X_1
)

var PosXs = []PosX{
	X_9,
	X_8,
	X_7,
	X_6,
	X_5,
	X_4,
	X_3,
	X_2,
	X_1,
}

func RX(r rune) PosX {
	return PosX(strings.IndexRune(posX, r))
}

type PosY int

const (
	Y_a PosY = iota
	Y_b
	Y_c
	Y_d
	Y_e
	Y_f
	Y_g
	Y_h
	Y_i
)

var PosYs = []PosY{
	Y_a,
	Y_b,
	Y_c,
	Y_d,
	Y_e,
	Y_f,
	Y_g,
	Y_h,
	Y_i,
}

func RY(r rune) PosY {
	return PosY(strings.IndexRune(posY, r))
}

type Pos struct {
	X PosX
	Y PosY
}

func (p *Pos) IsEmpty() bool {
	return p == nil || p.X < 0 || p.Y < 0 || int(p.X) >= len(PosXs) || int(p.Y) >= len(PosYs)
}
