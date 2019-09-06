package advmath

import (
	"math"
)

/*
Ridders computes derivative based on the ridders algorithm from 'Numerical Recipes Fortran 77'

First parameter (t) is the value to use for the computation
Second parameter (f) is the function for which we want a derivative
The method returns the derivative value computed for the function, note that it doesn't verify that
the function can be derived or not.
*/
func Ridders(t float64, f F, err float64) float64 {
	var calculatedError float64
	var fac float64

	h := math.Sqrt(err)
	cn := 1.2
	cn2 := cn * cn
	const n = 20
	//a := make([][]float64, n*n)
	var a [n][n]float64
	d := 0.0
	a[0][0] = (f(t+h) - f(t-h)) / (2.0 * h)

	for i := 1; i < n; i++ {
		h = h / cn
		a[0][i] = (f(t+h) - f(t-h)) / (2.0 * h)
		fac = cn2
		for j := 1; j < i; j++ {
			a[j][i] = (a[j-1][i]*fac - a[j-1][i-1]) / (fac - 1.0)
			fac = cn2 * fac
			calculatedError = math.Max(math.Abs(a[j][i]-a[j-1][i]), math.Abs(a[j][i]-a[j-1][i-1]))
			if calculatedError <= err {
				err = calculatedError
				d = a[j][i]
			}
			if math.Abs(a[i][i]-a[i-1][i-1]) >= 2*err {
				return d
			}
		}
	}
	return d
}

/*
Standard is a function to compute the derivative using the good old Newton's difference quotient.
Ridders usually gives results probably faster but precision might be better with this one ...

First parameter t is the value to use for the computation
Second parameter f is the function for which we want a derivative
It returns the derivative value
*/
func Standard(t float64, f F, err float64) float64 {
	h := math.Sqrt(err)
	return (f(t+h) - f(t-h)) / (2.0 * h)
}
