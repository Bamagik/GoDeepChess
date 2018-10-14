package base

import (
	"fmt"
)

type square struct {
	piece piece
}

type Board struct {
	board [][]square
}

func main() {
	boardArray := make([][]square, 8)

	for i := range boardArray {
		boardArray[i] = make([]square, 8)
		for j := range boardArray[i] {
			boardArray[i][j] = square{}
		}
	}

	board := Board{boardArray}

	fmt.Println(board.String())
}

func (board *Board) String() string {
	str := ""
	return str
}

func (sq square) String() string {
	if sq.piece == nil {
		return "."
	}
	var s string = sq.piece.String()
	return s
}
