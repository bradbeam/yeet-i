package maps

import (
	"fmt"
	"image"
	_ "image/png"
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
)

type TileType int

const (
	UnknownTileType TileType = iota
	GroundTileType
	WallTileType
)

func TileTypeToAsset(tileType TileType) string {
	switch tileType {
	case GroundTileType:
		return "assets/floor.png"
	case WallTileType:
		return "assets/wall.png"
	default:
		panic(fmt.Sprintf("unknown tile type: %d", tileType))
	}
}

type Tile struct {
	TileType TileType
	Image    *ebiten.Image
}

// Do we want to do anything here?
func NewTile() Tile {
	// TODO how to map up tile type
	return Tile{}
}

// Would like to find a way that we can load up the image/file with each tiletype
// so we can make use of them by tile type and load the images only once
func LoadTileImages(fs fs.FS) (map[TileType]Tile, error) {
	tileImages := make(map[TileType]Tile)

	// Is it dangerous to ignore UnknownTileType
	// Or should we make it some obvious looking image
	tileTypes := []TileType{GroundTileType, WallTileType}

	for _, tileType := range tileTypes {
		assetPath := TileTypeToAsset(tileType)

		assetFile, err := fs.Open(assetPath)
		if err != nil {
			return nil, fmt.Errorf("failed to find tile asset: %w", err)
		}

		asset, _, err := image.Decode(assetFile)
		if err != nil {
			_ = assetFile.Close()
			return nil, fmt.Errorf("failed to read tile image: %w", err)
		}

		_ = assetFile.Close()

		tileImages[tileType] = Tile{
			TileType: tileType,
			Image:    ebiten.NewImageFromImage(asset),
		}

	}

	return tileImages, nil
}
