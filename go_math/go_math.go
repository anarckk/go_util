package go_math

import "math"

// Ceil 四舍五入，向上取整
//
//	@param f
//	@return int
func Ceil(f float64) int64 {
	return int64(math.Ceil(f))
}
