package geometry

import "math"

const TILE_SIZE = 256

type TilePoint struct {
	TileX int
	TileY int
	Zoom  int
	PxX   int
	PxY   int

	linearZoom float64
}

func NewTilePoint(tileX, tileY, zoom, pxX, pxY int) *TilePoint {
	return &TilePoint{
		TileX: tileX,
		TileY: tileY,
		Zoom:  zoom,
		PxX:   pxX,
		PxY:   pxY,

		linearZoom: float64(int(1) << zoom),
	}
}

func (t *TilePoint) ToLatLng() LatLng {
	normX, normY := t.toNoramlizedXY()

	return LatLng{
		Lat: math.Atan(math.Sinh(math.Pi*(1-2*normY))) * 180 / math.Pi,
		Lng: normX*360 - 180,
	}
}

func (t *TilePoint) toAbsolutePx() (int, int) {
	return t.TileX*TILE_SIZE + t.PxX, t.TileY*TILE_SIZE + t.PxY
}

func (t *TilePoint) toNoramlizedXY() (float64, float64) {
	absPxX, absPxY := t.toAbsolutePx()

	return float64(absPxX) / (TILE_SIZE * t.linearZoom),
		float64(absPxY) / (TILE_SIZE * t.linearZoom)
}
