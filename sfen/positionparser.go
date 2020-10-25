package sfen

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func parseString(s *bufio.Scanner) (string, error) {
	if !s.Scan() {
		if err := s.Err(); err != nil {
			return "", err
		} else {
			return "", io.ErrUnexpectedEOF
		}
	}

	return s.Text(), nil
}

func parsePhrase(s *bufio.Scanner, str string) error {
	t, err := parseString(s)
	if err == io.ErrUnexpectedEOF {
		return fmt.Errorf("Unexpected EOF: parse `%v`", str)
	} else if err != nil {
		return err
	}

	if t != str {
		return fmt.Errorf("unexpected token: expected=`%s` but actual=`%s`", str, t)
	}

	return nil
}

// Parse position command
func ParsePosition(cmd string) (*Surface, error) {
	s := bufio.NewScanner(strings.NewReader(cmd))
	s.Split(bufio.ScanWords)

	if err := parsePhrase(s, "position"); err != nil {
		return nil, err
	}

	t, err := parseString(s)
	if err == io.ErrUnexpectedEOF {
		return nil, fmt.Errorf("Unexpected EOF: parse position header")
	} else if err != nil {
		return nil, err
	}

	var surface *Surface
	switch t {
	case "startpos":
		surface = NewSurfaceStartpos()
	case "sfen":
		t, err := parseString(s)
		if err == io.ErrUnexpectedEOF {
			return nil, fmt.Errorf("Unexpected EOF: parse sfen body")
		} else if err != nil {
			return nil, err
		}

		suf, err := NewSurface(t)
		if err != nil {
			return nil, err
		}
		surface = suf
	default:
		return nil, fmt.Errorf("unexpected token: `%s`", t)
	}

	if err := parsePhrase(s, "moves"); err != nil {
		return nil, err
	}

	for s.Scan() {
		if err := surface.Move(s.Text()); err != nil {
			return nil, err
		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}

	return surface, nil
}
