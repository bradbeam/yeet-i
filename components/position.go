package components

import (
	"math"

	"github.com/yohamta/donburi"
)

type PositionComponent struct {
	X int
	Y int
}

var Position = donburi.NewComponentType[PositionComponent]()

func (p *PositionComponent) GetManhattanDistance(other *PositionComponent) int {
	xDist := math.Abs(float64(p.X - other.X))
	yDist := math.Abs(float64(p.Y - other.Y))
	return int(xDist) + int(yDist)
}

func (p *PositionComponent) IsEqual(other *PositionComponent) bool {
	return (p.X == other.X && p.Y == other.Y)
}
