package sfen

import (
	"context"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Surface struct {
	phase    Player
	board    [9][9]*Piece
	captured []*Piece
	nextStep int
}

func NewSurfaceEmpty() *Surface {
	return &Surface{
		phase:    Player_BLACK,
		nextStep: 1,
	}
}

func NewSurface(pos string) (*Surface, error) {
	var sf, pl, cp string
	var step int
	if _, err := fmt.Sscanf(pos, "%s %s %s %d", &sf, &pl, &cp, &step); err != nil {
		return nil, err
	}

	parser := newPosParser(sf)
	b, err := parser.parseBoard()
	if err != nil {
		return nil, err
	}

	var Player Player
	switch pl {
	case "b":
		Player = Player_BLACK
	case "w":
		Player = Player_WHITE
	}

	captured := parseCaptured(cp)

	sort.Sort(pieceSlice(captured))

	return &Surface{
		phase:    Player,
		board:    b,
		captured: captured,
		nextStep: step,
	}, nil
}

func NewSurfaceStartpos() *Surface {
	s := &Surface{}

	s.InitStartpos()

	return s
}

func (s *Surface) InitStartpos() {
	var board [9][9]*Piece

	board[Y_a][X_9] = &Piece{
		Player: Player_WHITE,
		Type:   Piece_KYOU,
	}
	board[Y_a][X_8] = &Piece{
		Player: Player_WHITE,
		Type:   Piece_KEI,
	}
	board[Y_a][X_7] = &Piece{
		Player: Player_WHITE,
		Type:   Piece_GIN,
	}
	board[Y_a][X_6] = &Piece{
		Player: Player_WHITE,
		Type:   Piece_KIN,
	}
	board[Y_a][X_5] = &Piece{
		Player: Player_WHITE,
		Type:   Piece_GYOKU,
	}
	board[Y_a][X_4] = &Piece{
		Player: Player_WHITE,
		Type:   Piece_KIN,
	}
	board[Y_a][X_3] = &Piece{
		Player: Player_WHITE,
		Type:   Piece_GIN,
	}
	board[Y_a][X_2] = &Piece{
		Player: Player_WHITE,
		Type:   Piece_KEI,
	}
	board[Y_a][X_1] = &Piece{
		Player: Player_WHITE,
		Type:   Piece_KYOU,
	}
	board[Y_b][X_2] = &Piece{
		Player: Player_WHITE,
		Type:   Piece_KAKU,
	}
	board[Y_b][X_8] = &Piece{
		Player: Player_WHITE,
		Type:   Piece_HISHA,
	}

	board[Y_i][X_9] = &Piece{
		Player: Player_BLACK,
		Type:   Piece_KYOU,
	}
	board[Y_i][X_8] = &Piece{
		Player: Player_BLACK,
		Type:   Piece_KEI,
	}
	board[Y_i][X_7] = &Piece{
		Player: Player_BLACK,
		Type:   Piece_GIN,
	}
	board[Y_i][X_6] = &Piece{
		Player: Player_BLACK,
		Type:   Piece_KIN,
	}
	board[Y_i][X_5] = &Piece{
		Player: Player_BLACK,
		Type:   Piece_GYOKU,
	}
	board[Y_i][X_4] = &Piece{
		Player: Player_BLACK,
		Type:   Piece_KIN,
	}
	board[Y_i][X_3] = &Piece{
		Player: Player_BLACK,
		Type:   Piece_GIN,
	}
	board[Y_i][X_2] = &Piece{
		Player: Player_BLACK,
		Type:   Piece_KEI,
	}
	board[Y_i][X_1] = &Piece{
		Player: Player_BLACK,
		Type:   Piece_KYOU,
	}
	board[Y_h][X_8] = &Piece{
		Player: Player_BLACK,
		Type:   Piece_KAKU,
	}
	board[Y_h][X_2] = &Piece{
		Player: Player_BLACK,
		Type:   Piece_HISHA,
	}

	for _, x := range PosXs {
		board[Y_c][x] = &Piece{
			Player: Player_WHITE,
			Type:   Piece_FU,
		}
		board[Y_g][x] = &Piece{
			Player: Player_BLACK,
			Type:   Piece_FU,
		}
	}

	s.board = board
	s.phase = Player_BLACK
	s.nextStep = 1
}

func (s *Surface) PrintAA(out io.Writer) error {
	for _, x := range posX {
		fmt.Fprintf(out, " %c", x)
	}
	fmt.Fprintln(out)
	for _, _ = range posX {
		fmt.Fprint(out, "--")
	}
	fmt.Fprintln(out, "--")

	for y, xs := range s.board {
		for _, p := range xs {
			fmt.Fprint(out, p.aa())
		}
		fmt.Fprintf(out, " | %c", posY[y])
		fmt.Fprintln(out)
	}

	fmt.Fprint(out, "captured: ")
	for _, p := range s.captured {
		fmt.Fprintf(out, "%c", p.a())
	}
	fmt.Fprintln(out)

	return nil
}

func printCaptured(captured []*Piece) string {
	if len(captured) == 0 {
		return "-"
	}

	var cps []string
	count := 1
	var prev *Piece
	for _, cp := range captured {
		if prev == nil {
			prev = cp
			continue
		} else if prev.equal(cp) {
			count++
			continue
		}

		p := sfenPiece[prev.Type : prev.Type+1]
		if prev.Player == Player_WHITE {
			p = strings.ToLower(p)
		}
		if count != 1 {
			p = fmt.Sprintf("%d%s", count, p)
		}
		cps = append(cps, p)

		count = 1
		prev = cp
	}

	p := sfenPiece[prev.Type : prev.Type+1]
	if prev.Player == Player_WHITE {
		p = strings.ToLower(p)
	}
	if count != 1 {
		p = fmt.Sprintf("%d%s", count, p)
	}
	cps = append(cps, p)

	return strings.Join(cps, "")
}

func (s *Surface) PrintSFEN(w io.Writer) error {
	var lines []string
	for _, xs := range s.board {
		var sp int
		var line string
		for _, p := range xs {
			if p.isEmpty() {
				sp++
				continue
			}

			if sp != 0 {
				line += fmt.Sprintf("%d", sp)
				sp = 0
			}
			line += p.sfen()
		}

		if sp != 0 {
			line += fmt.Sprintf("%d", sp)
		}

		lines = append(lines, line)
	}

	var phase string
	switch s.phase {
	case Player_WHITE:
		phase = "w"
	case Player_BLACK:
		phase = "b"
	}

	_, err := fmt.Fprintf(w,
		"%s %s %s %d",
		strings.Join(lines, "/"),
		phase,
		printCaptured(s.captured),
		s.nextStep)
	return err
}

func (s *Surface) move(m *move) error {
	var p, pp *Piece
	if m.putted != Piece_NULL {
		for i, cp := range s.captured {
			if cp.Player == s.phase && cp.Type == m.putted {
				p = s.captured[i]
				s.captured = append(s.captured[:i], s.captured[i+1:]...)
				break
			}
		}
		if p == nil {
			return fmt.Errorf("captured Piece not found: Piece=%d, Player=%d", m.putted, s.phase)
		}
	} else {
		p, s.board[m.from.Y][m.from.X] = s.board[m.from.Y][m.from.X], nil
	}

	p.Promoted = p.Promoted || m.promoted

	pp, s.board[m.to.Y][m.to.X] = s.board[m.to.Y][m.to.X], p

	if !pp.isEmpty() {
		pp.Player = pp.Player.flip()
		pp.Promoted = false
		s.captured = append(s.captured, pp)
	}

	s.nextStep++
	s.phase = s.phase.flip()

	sort.Sort(pieceSlice(s.captured))

	return nil
}

func (s *Surface) Move(move string) error {
	parser := newMoveParser(move)

	m, err := parser.parseMove()
	if err != nil {
		return err
	}

	return s.move(m)
}

func (s *Surface) SetStep(step int) {
	s.nextStep = step
}

func (s *Surface) SetPlayer(player Player) {
	s.phase = player
}

func (s *Surface) SetPiece(pos *Pos, piece *Piece) {
	if pos.IsEmpty() {
		s.captured = append(s.captured, piece)
		sort.Sort(pieceSlice(s.captured))
		return
	}

	s.board[pos.Y][pos.X] = piece
}

func (s *Surface) GetPiece(x PosX, y PosY) *Piece {
	return s.board[y][x]
}

func (s *Surface) GetCaptured() []*Piece {
	return s.captured
}

func (s *Surface) Scan(ctx context.Context, f func(*Pos, *Piece)) error {
	for _, y := range PosYs {
		for _, x := range PosXs {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				f(&Pos{X: x, Y: y}, s.board[y][x])
			}
		}
	}

	for _, p := range s.captured {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			f(nil, p)
		}
	}

	return nil
}
