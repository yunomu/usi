package sfen

import (
	"testing"

	"bytes"
	"fmt"
)

func TestParsePosition(t *testing.T) {
	s, err := ParsePosition("position startpos moves 7g7f 3c3d 6g6f 8c8d 2h6h 8d8e 8h7g 7a6b 5i4h 5c5d 3i3h 5a4b 4h3i 1c1d 1g1f 3a3b 7i7h 6b5c 3i2h 6a5b 6i5h 4b3a 4g4f 2c2d 3g3f 3b2c 5h4g 4a3b 2i3g 2b4d 5g5f 3a2b 6f6e 4d7g+ 7h7g 4c4d 9g9f 9c9d 2g2f 7c7d B*6f 8b6b 3g4e 5c4b 6f4d 4b3c 4e3c+ 2a3c 4d6b+ 5b6b R*7a N*4e 7a4a+ B*7c 6e6d 7c6d 4a8a B*7c N*3g 4e3g+ 3h3g N*4e 4i3h 4e3g+ 3h3g 3c4e 6h6d 7c6d B*4d S*3c 4d6b+ 4e3g+ 2h3g R*4c G*2a 2b1c S*2b 3c2b N*4d P*4e 3g4h 4e4f 4g3g S*4g 4h5g 4g5f+ 5g5f G*5e 5f6g 4f4g+ 2a1a 5e4f 4d3b+ 2c3b S*1b 4g5g 6g7h 6d5e P*4d 4f3g L*1h N*6e 1f1e 6e7g+ 8i7g 5g6g 7h8h 6g7g 8h9h 7g8h 9h9g S*8f 8g8f G*8g")
	if err != nil {
		t.Fatalf("ParsePosition: %v", err)
	}

	var buf bytes.Buffer
	if err := s.PrintSFEN(&buf); err != nil {
		t.Fatalf("PrintSFEN: %v", err)
	}

	expected := "l+R6G/3+B2ssS/3p1r2k/p1p1pPppp/1p2b3P/PPP3PP1/Kg4g2/1+p6L/L7L b GS3Nn2p 115"
	if str := buf.String(); str != expected {
		t.Logf("expected=%v", expected)
		t.Logf("actual  =%v", buf.String())
		t.Errorf("PrintSFEN")
	}
}

func ExampleParsePosition() {
	s, err := ParsePosition("position startpos moves 7g7f 3c3d 6g6f 8c8d 2h6h 8d8e 8h7g 7a6b 5i4h 5c5d 3i3h 5a4b 4h3i 1c1d 1g1f 3a3b 7i7h 6b5c 3i2h 6a5b 6i5h 4b3a 4g4f 2c2d 3g3f 3b2c 5h4g 4a3b 2i3g 2b4d 5g5f 3a2b 6f6e 4d7g+ 7h7g 4c4d 9g9f 9c9d 2g2f 7c7d B*6f 8b6b 3g4e 5c4b 6f4d 4b3c 4e3c+ 2a3c 4d6b+ 5b6b R*7a N*4e 7a4a+ B*7c 6e6d 7c6d 4a8a B*7c N*3g 4e3g+ 3h3g N*4e 4i3h 4e3g+ 3h3g 3c4e 6h6d 7c6d B*4d S*3c 4d6b+ 4e3g+ 2h3g R*4c G*2a 2b1c S*2b 3c2b N*4d P*4e 3g4h 4e4f 4g3g S*4g 4h5g 4g5f+ 5g5f G*5e 5f6g 4f4g+ 2a1a 5e4f 4d3b+ 2c3b S*1b 4g5g 6g7h 6d5e P*4d 4f3g L*1h N*6e 1f1e 6e7g+ 8i7g 5g6g 7h8h 6g7g 8h9h 7g8h 9h9g S*8f 8g8f G*8g")
	if err != nil {
		return
	}

	var buf bytes.Buffer
	if err := s.PrintSFEN(&buf); err != nil {
		return
	}

	expected := "l+R6G/3+B2ssS/3p1r2k/p1p1pPppp/1p2b3P/PPP3PP1/Kg4g2/1+p6L/L7L b GS3Nn2p 115"
	str := buf.String()
	fmt.Println(str == expected)
	// Output: true
}
