package board

// contains commands for moving

type coordinate struct {
	file, rank int
}

type move interface {
	PerformMove()
}

type normalMove struct {
	start, dest coordinate
}

func (m normalMove) PerformMove(board *Board) *Board {
	b := *board

	b.board[m.dest.rank][m.dest.file] = b.board[m.start.rank][m.start.file]
	b.board[m.start.rank][m.start.file] = square{}

	return &b
}
