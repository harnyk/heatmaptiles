package geometry

import "math"

type TilePoint struct {
	TileX    int
	TileY    int
	Zoom     int
	PxX      int
	PxY      int
	TileSize int
}

func (t *TilePoint) ToLatLng() LatLng {
	normX, normY := t.toNoramlizedXY()

	return LatLng{
		Lat: math.Atan(math.Sinh(math.Pi*(1-2*normY))) * 180 / math.Pi,
		Lng: normX*360 - 180,
	}
}

func (t *TilePoint) toAbsolutePx() (int, int) {
	return t.TileX*t.TileSize + t.PxX, t.TileY*t.TileSize + t.PxY
}

func (t *TilePoint) toNoramlizedXY() (float64, float64) {
	absPxX, absPxY := t.toAbsolutePx()
	linearZoom := 1 << t.Zoom

	return float64(absPxX) / float64(t.TileSize*linearZoom),
		float64(absPxY) / float64(t.TileSize*linearZoom)
}
