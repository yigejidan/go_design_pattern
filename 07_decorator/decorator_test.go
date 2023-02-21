package _7_decorator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColorSquare_Draw(t *testing.T) {
	sq := Square{}
	csq := NewColorSquare(&sq, "red")
	got := csq.Draw()
	assert.Equal(t, "这是一个正方形, color is red", got)
}
