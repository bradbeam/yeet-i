package layers

import "github.com/yohamta/donburi/ecs"

const (
	Default ecs.LayerID = iota
	Floor
	Wall
	RealWorld
)
