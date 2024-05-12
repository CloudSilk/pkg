package utils

import (
	"testing"
)

func TestConvertGCJ02ToBD09(t *testing.T) {
	lon, lat := 39.908722, 116.397496
	t.Log("TX coordinates:", lon, lat)

	bdLon, bdLat := ConvertGCJ02ToBD09(lat, lon)
	t.Log("BD coordinates:", bdLon, bdLat)

	txLon, txLat := ConvertBD09ToGCJ02(bdLat, bdLon)
	t.Log("TX coordinates:", txLon, txLat)
}
