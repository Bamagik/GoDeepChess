package board

import (
	"fmt"
)

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
	str := "  "
	for i := 0; i < len(board.board); i++ {
		str += string(i+97) + " " // get a -> h
	}

	str += "\n"

	for i := 7; i >= 0; i-- { // i is y axis. i.e 8 -> 1
		str += string(i+1) + " "
		for j := 0; j < len(board.board); j++ { // j is the x axis. i.e. a->h
			str += board.board[i][j].String() + " "
		}
	}

	return str
}
