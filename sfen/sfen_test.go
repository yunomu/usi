package sfen

import (
	"testing"

	"bytes"
	"fmt"
)

func TestNewSurface(t *testing.T) {
	input := []string{
		"lnsgkgsnl/1r5b1/ppppppppp/9/9/9/PPPPPPPPP/1B5R1/LNSGKGSNL b - 1",
		"l+R6G/3+B2ssS/3p1r2k/p1p1pPppp/1p2b3P/PPP3PP1/Kg4g2/1+p6L/L7L b GS3Nn2p 115",
	}

	for _, sfen := range input {
		s, err := NewSurface(sfen)
		if err != nil {
			t.Fatalf("NewSurface: %v", err)
		}

		var buf bytes.Buffer
		if err := s.PrintSFEN(&buf); err != nil {
			t.Fatalf("Surface.PrintSFEN: %v", err)
		}

		if sfen != buf.String() {
			t.Logf("expected=%s", sfen)
			t.Logf("actual  =%s", buf.String())
			t.Errorf("mismatch")
		}
	}
}

func ExampleNewSurface() {
	sfen := "l+R6G/3+B2ssS/3p1r2k/p1p1pPppp/1p2b3P/PPP3PP1/Kg4g2/1+p6L/L7L b GS3Nn2p 115"
	s, err := NewSurface(sfen)
	var _ = err

	p := s.GetPiece(X_9, Y_a)
	fmt.Printf("%c", p.a())
	p = s.GetPiece(X_6, Y_b)
	fmt.Printf("%c", p.a())
	// Output: lB
}

func TestCaptured(t *testing.T) {
	input := []string{
		"GS3Nn2p",
		"G12p",
	}

	for _, s := range input {
		if a := printCaptured(parseCaptured(s)); a != s {
			t.Errorf("mismatch: expected=%s actual=%s", s, a)
		}
	}
}
