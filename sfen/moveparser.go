package sfen

import (
	"fmt"
	"io"
	"strings"
)

type move struct {
	putted   PieceType
	from, to *Pos
	promoted bool
}

type moveParser struct {
	r io.RuneScanner
}

func newMoveParser(s string) *moveParser {
	return &moveParser{
		r: strings.NewReader(s),
	}
}

func (m *moveParser) read() (rune, error) {
	r, _, err := m.r.ReadRune()
	return r, err
}

func (m *moveParser) unread() error {
	return m.r.UnreadRune()
}

func (m *moveParser) parsePuttedPiece() (PieceType, error) {
	p, err := m.read()
	if err != nil {
		return Piece_NULL, err
	}

	i := strings.IndexRune(sfenPiece, p)
	if i == -1 {
		return Piece_NULL, m.unread()
	}

	c, err := m.read()
	if err != nil {
		return Piece_NULL, err
	}
	if c != '*' {
		return Piece_NULL, fmt.Errorf("expected '*' but actual=`%c`", c)
	}

	return PieceType(i), nil
}

func (m *moveParser) parsePos() (*Pos, error) {
	xr, err := m.read()
	if err != nil {
		return nil, err
	}

	yr, err := m.read()
	if err != nil {
		return nil, err
	}

	return &Pos{X: RX(xr), Y: RY(yr)}, nil
}

func (m *moveParser) parseMove() (*move, error) {
	ptype, err := m.parsePuttedPiece()
	if err != nil {
		return nil, err
	}

	var from *Pos
	if ptype == Piece_NULL {
		from, err = m.parsePos()
		if err != nil {
			return nil, fmt.Errorf("Parse src: %v", err)
		}
	}

	to, err := m.parsePos()
	if err != nil {
		return nil, fmt.Errorf("Parse dst: %v", err)
	}

	var promoted bool
	if pr, err := m.read(); err == nil {
		if pr == '+' {
			promoted = true
		} else {
			return nil, fmt.Errorf("unexpected rune='%c'", pr)
		}
	}

	return &move{putted: ptype, from: from, to: to, promoted: promoted}, nil
}
