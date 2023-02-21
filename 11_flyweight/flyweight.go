package _1_flyweight

// ChessPieceUnit 棋子享元
type ChessPieceUnit struct {
	ID    uint
	Name  string
	Color string
}

// 享元池
var units = map[int]*ChessPieceUnit{
	1: {
		ID:    1,
		Name:  "将",
		Color: "red",
	},
	2: {
		ID:    1,
		Name:  "帅",
		Color: "black",
	},
	// ....
}

func NewChessPieceUnit(id int) *ChessPieceUnit {
	return units[id]
}

type ChessPiece struct {
	Unit *ChessPieceUnit
	X    int
	Y    int
}

// ChessBoard 棋局
type ChessBoard struct {
	chessPieces map[int]*ChessPiece
}

// NewChessBoard 初始化棋盘
func NewChessBoard() *ChessBoard {
	board := &ChessBoard{chessPieces: map[int]*ChessPiece{}}
	for id := range units {
		board.chessPieces[id] = &ChessPiece{
			Unit: NewChessPieceUnit(id),
			X:    0, // 这里应该是对应棋子的初始位置
			Y:    0,
		}
	}
	return board
}

func (c *ChessBoard) Move(id, x, y int) {
	c.chessPieces[id].X = x
	c.chessPieces[id].Y = y
}
