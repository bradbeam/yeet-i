package maps

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type MapTile struct {
	Tile

	X int
	Y int
}

type Level struct {
	Tiles []*MapTile
	Rooms []Rect
	Dimensions
}

type Dimensions struct {
	Height     int
	Width      int
	TileWidth  int
	TileHeight int
}

func NewLevel(dimensions Dimensions, mapTiles map[TileType]Tile) *Level {
	l := &Level{
		Dimensions: dimensions,
	}

	l.generateLevel(dimensions, mapTiles)

	return l
}

func (l *Level) Draw(screen *ebiten.Image) {
	for x := 0; x < l.Width; x++ {
		for y := 0; y < l.Height; y++ {
			tile := l.Tiles[l.getIndexFromXY(x, y)]
			op := &ebiten.DrawImageOptions{}

			op.GeoM.Translate(float64(tile.X), float64(tile.Y))
			screen.DrawImage(tile.Image, op)
		}
	}
}

const (
	minRoomSize = 8
	maxRoomSize = 12
	minRooms    = 6
	maxRooms    = 30
)

func (l *Level) generateLevel(dimensions Dimensions, mapTiles map[TileType]Tile) {
	l.Tiles = initializeMap(dimensions, mapTiles)

	// TODO levelHeight needs to be adjusted to account for UI room

	rooms := rand.Intn(maxRooms) + minRooms

	for idx := 0; idx < rooms; idx++ {
		w := rand.Intn(maxRoomSize) + minRoomSize
		h := rand.Intn(maxRoomSize) + minRoomSize
		x := rand.Intn(dimensions.Width - w - 1)
		y := rand.Intn(dimensions.Height - h - 1)

		newRoom := NewRect(x, y, w, h)

		if l.invalidRoom(newRoom) {
			continue
		}

		l.createRoom(newRoom, mapTiles[GroundTileType])

		l.Rooms = append(l.Rooms, newRoom)

		// No rooms to connect with yet
		if len(l.Rooms) <= 1 {
			continue
		}

		// Create tunnels between rooms
		newX, newY := newRoom.Center()

		// Since we've just added our new room above, we need to look at n-2
		prevX, prevY := l.Rooms[len(l.Rooms)-2].Center()

		// NOTE
		// slight change in source here; need to verify it still works
		horizontalXStart, horizontalXEnd, horizontalY := prevX, newX, newY
		verticalYStart, verticalYEnd, verticalX := prevY, newY, prevX

		if rand.Intn(100) > 50 {
			horizontalY = prevY
			verticalX = newX
		}

		l.createHorizontalTunnel(horizontalXStart, horizontalXEnd, horizontalY, mapTiles[GroundTileType])
		l.createVerticalTunnel(verticalYStart, verticalYEnd, verticalX, mapTiles[GroundTileType])
	}
}

func (l *Level) invalidRoom(room Rect) bool {
	for _, existingRoom := range l.Rooms {
		if room.Intersect(existingRoom) {
			return true
		}
	}

	return false
}

func (l *Level) createRoom(room Rect, groundTile Tile) {
	for y := room.Y1 + 1; y < room.Y2; y++ {
		for x := room.X1 + 1; x < room.X2; x++ {
			index := l.getIndexFromXY(x, y)
			// l.Tiles[index].Blocked = false
			l.Tiles[index].Tile = groundTile
		}
	}
}

// These two tunnel funcs are near identical and might be good candidates for refactor.
// The key difference is that we are either supplying the X or Y value that is constant.
// So most likely we could trigger this with another parameter, but that might make it
// a bit heafty.
func (l *Level) createHorizontalTunnel(x1 int, x2 int, y int, groundTile Tile) {
	minX := x1
	if x2 < minX {
		minX = x2
	}

	maxX := x1
	if x2 > maxX {
		maxX = x2
	}

	for x := minX; x <= maxX; x++ {
		index := l.getIndexFromXY(x, y)
		// TODO s/Height/levelHeight/
		// if index > 0 && index < gd.ScreenWidth*levelHeight {
		if index > 0 && index < l.Width*l.Height {
			l.Tiles[index].Tile = groundTile
		}
	}
}

func (l *Level) createVerticalTunnel(y1 int, y2 int, x int, groundTile Tile) {
	minY := y1
	if y2 < minY {
		minY = y2
	}

	maxY := y1
	if y2 > maxY {
		maxY = y2
	}

	for y := minY; y <= maxY; y++ {
		index := l.getIndexFromXY(x, y)
		// TODO s/Height/levelHeight/
		// if index > 0 && index < gd.ScreenWidth*levelHeight {
		if index > 0 && index < l.Width*l.Height {
			l.Tiles[index].Tile = groundTile
		}

	}
}

func initializeMap(dimensions Dimensions, mapTiles map[TileType]Tile) []*MapTile {
	tiles := make([]*MapTile, 0, dimensions.Height*dimensions.Width)

	// We'll start with Y first so we can assemble the slice
	// row by row
	for y := 0; y < dimensions.Height; y++ {
		for x := 0; x < dimensions.Width; x++ {
			tiles = append(tiles, &MapTile{
				X:    x * dimensions.TileWidth,
				Y:    y * dimensions.TileHeight,
				Tile: mapTiles[WallTileType],
			})

		}
	}
	return tiles
}

func (l *Level) getIndexFromXY(x int, y int) int {
	return (y * l.Dimensions.Width) + x
}
