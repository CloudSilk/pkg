/*
* 百度地图使用BD09坐标
* 中国正常坐标是GCJ02
* 腾讯地图用的也是GCJ02坐标
 */

package utils

import (
	"math"

	"github.com/shopspring/decimal"
)

var (
	param1 = decimal.NewFromFloat(0.00002)
	param2 = decimal.NewFromFloat(0.000003)
	param3 = decimal.NewFromFloat(0.0065)
	param4 = decimal.NewFromFloat(0.006)
)

/**
* 坐标转换，腾讯地图转换成百度地图坐标
* @param lat 腾讯纬度
* @param lon 腾讯经度
* @return 返回结果：经度,纬度
 */
func ConvertGCJ02ToBD09(lat, lon float64) (float64, float64) {
	latDecimal := decimal.NewFromFloat(lat)
	lonDecimal := decimal.NewFromFloat(lon)

	xPi := decimal.NewFromFloat(math.Pi)
	x, y := lonDecimal.Copy(), latDecimal.Copy()

	z := decimal.NewFromFloat(math.Sqrt(x.Mul(x).Add(y.Mul(y)).InexactFloat64())).Add(param1.Mul(y.Mul(xPi).Sin()))
	theta := decimal.NewFromFloat(math.Atan2(y.InexactFloat64(), x.InexactFloat64())).Add(param2.Mul(x.Mul(xPi).Cos()))
	bdLon, _ := z.Mul(theta.Cos()).Add(param3).Round(6).Float64()
	bdLat, _ := z.Mul(theta.Sin()).Add(param4).Round(6).Float64()

	return bdLon, bdLat
}

/**
* 坐标转换，百度地图坐标转换成腾讯地图坐标
* @param lat  百度坐标纬度
* @param lon  百度坐标经度
* @return 返回结果：纬度,经度
 */
func ConvertBD09ToGCJ02(lat, lon float64) (float64, float64) {
	xPi := decimal.NewFromFloat(math.Pi)
	x := decimal.NewFromFloat(lon).Sub(param3)
	y := decimal.NewFromFloat(lat).Sub(param4)

	z := decimal.NewFromFloat(math.Sqrt(x.Mul(x).Add(y.Mul(y)).InexactFloat64())).Sub(param1.Mul(decimal.NewFromFloat(math.Sin(y.Mul(xPi).InexactFloat64()))))

	theta := decimal.NewFromFloat(math.Atan2(y.InexactFloat64(), x.InexactFloat64())).Sub(param2.Mul(decimal.NewFromFloat(math.Cos(x.Mul(xPi).InexactFloat64()))))
	txLon, _ := z.Mul(theta.Cos()).Round(6).Float64()
	txLat, _ := z.Mul(theta.Sin()).Round(6).Float64()

	return txLon, txLat
}
