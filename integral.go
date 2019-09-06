package advmath

import (
	"math"
)

/*
Simpson uses the simpson method to compute the integral of a given function between a and b.
Number of intervals computed are by default 10^5, this is the best precision that you can have with go.

First parameter inf is the lower boundary
Second parameter sup is the upper boundary
Third parameter is the number of iteration
Fourth parameter f is the function to integrate
The method returns the value of the integral
*/
func Simpson(inf float64, sup float64, f F, n int) (float64, error) {
	if n%2 != 0 {
		return 0, &MathError{
			s: "Invalid number of iterations, for simpson, iterations number has to be even",
		}
	}
	h := (sup - inf) / float64(n)
	s := f(inf) + f(sup)
	var i, j int
	for i = 1; i < n; i += 2 {
		s += 4 * f(inf+float64(i)*h)
	}
	for j = 2; j < n-1; j += 2 {
		s += 2 * f(inf+float64(j)*h)
	}
	return s * h / 3, nil
}

/*
Trapezoidal uses the Trapezoidal rule to compute the integral of a function. It provides
a quite good approximation but it should probably be used for very simple computations

First parameter is the inferior boundary
Second parameter is the first boundary
Third parameter is the number of iterations
Fourth parameter is the precision
*/
func Trapezoidal(inf float64, sup float64, f F, n int, precision float64) float64 {
	//Finding optimal n is cumbersome and would cost too much, so we define it
	//to 100000 and compute the error to see if we are close.
	if n == 0 {
		n = 100000
	}

	h := (sup - inf) / float64(n)
	result := 0.5*f(inf) + 0.5*f(sup)
	var previous float64
	for i := 1; i < n; i++ {
		previous = result
		result += f(inf + float64(i)*h)
		if math.Abs(result-previous) <= precision {
			break
		}
	}
	result *= h
	return result
}

/*
Romberg uses the romberg method to compute the integral of a function. It provides a better
approximation than the Trapezoidal method.

First parameter is the inferior boundary
Second parameter is the first boundary
Third is the function
Fourth parameter is the precision
*/
func Romberg(inf float64, sup float64, f F, maxSteps int, precision float64) float64 {
	if maxSteps == 0 {
		//This should be enough for most precisions but it will be a bit slower!
		maxSteps = 20
	}
	previousNew := 0.0
	currentNew := 0.0

	for i := 1; i <= maxSteps; i++ {
		previous := previousNew
		previousNew = trapezoidalr(inf, sup, f, i, previous)

		if i == 1 {
			currentNew = previousNew
		} else {
			current := currentNew
			currentNew = (4.0*previousNew - previous) / 3.0
			if i > 1 && math.Abs(currentNew-current) < precision {
				break
			}
		}
	}
	//Might not be the best result if we've been through the maxSteps iterations
	return currentNew
}

/*
trapezoidalr is a helper function used to compute the trapezoidal rule of a function based
on the iteration and the previous value. This is used by the Romberg method to aproximate the values
at each steps.
*/
func trapezoidalr(inf float64, sup float64, f F, m int, previous float64) float64 {
	if m > 1 {
		ep := int(math.Pow(2, float64(m-2)))
		c := 0.0

		for j := 1; j <= ep; j++ {
			y := (float64(2*ep-2*j+1)*inf + float64(2*j-1)*sup) / (2 * float64(ep))
			c += f(y)
		}
		c = .5*previous + (sup-inf)*c/(2*float64(ep))
		return c
	}
	return (sup - inf) / 2.0 * (f(sup) + f(inf))
}
