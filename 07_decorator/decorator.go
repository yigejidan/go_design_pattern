package _7_decorator

/*
装饰器模式
*/

type IDraw interface {
	Draw() string
}

type Square struct {
}

func (s *Square) Draw() string {
	return "这是一个正方形"
}

type ColorSquare struct {
	Square IDraw
	Color  string
}

func NewColorSquare(square IDraw, color string) ColorSquare {
	return ColorSquare{Color: color, Square: square}
}

func (c ColorSquare) Draw() string {
	return c.Square.Draw() + ", color is " + c.Color
}
