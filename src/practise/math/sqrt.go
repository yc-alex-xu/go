package math

/*
Sqrt (x float64)  return the square root of x
*/
func Sqrt(x float64) (z float64) {
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * x)
	}
	return
}
