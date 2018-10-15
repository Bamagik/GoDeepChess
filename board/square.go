package board

type color string // enum type

const (
	White color = "W"
	Black color = "B"
)

type square struct {
	piece piece
	color color
}

func (sq square) String() string {
	if sq.piece == nil {
		return "."
	}
	s := sq.piece.String()
	return s
}
