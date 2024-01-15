package geometry

import (
	"reflect"
	"testing"
)

func TestTilePoint_ToLatLng(t *testing.T) {
	type fields struct {
		tilePoint *TilePoint
	}
	tests := []struct {
		name   string
		fields fields
		want   LatLng
	}{
		{
			name: "test",
			fields: fields{
				tilePoint: &TilePoint{
					TileSize: 256,
					TileX:    1158148,
					TileY:    667418,
					Zoom:     21,
					PxX:      127,
					PxY:      127,
				},
			},
			want: LatLng{
				Lat: 54.59489282994036,
				Lng: 18.809365555644035,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tilePoint := tt.fields.tilePoint
			if got := tilePoint.ToLatLng(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TilePoint.ToLatLng() = %v, want %v", got, tt.want)
			}
		})
	}
}
