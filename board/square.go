package board

type square struct {
	piece piece
}

func (sq square) String() string {
	if sq.piece == nil {
		return "."
	}
	s := sq.piece.String()
	return s
}
