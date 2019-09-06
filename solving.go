package advmath

import (
	"fmt"
	"math"
)

/*
Newton finds a zero near the initial value using the Newton
algorithm. Note that depending on the sign of the init value the
algorithm will return a positive solution or a negative solution.
It doesn't find all the solutions but only the closest to init.

If there is a division by zero, the process will stop and return -1 in the err return
value. Note that it doesn't mean that there is no solution but that the algorithm didn't
manage to find one.

First param init is an initial estimated value of the zero
Second param f is the function to solve
Third param is the number of iteration, it is optional and set to 1000 by default
Fourth param precision is the precision required, used to have an end condition
return the zero and zero in the error field or a -1 in the error field if it failed
*/
func Newton(init float64, f F, n int, precision float64) (float64, int) {
	//This is in case of a zero division
	defer func() {
		if err := recover(); err != nil {
			fmt.Print("Error in solving.go: %T", err)
		}
	}()

	if n == 0 {
		//This should be enough for pretty much every precision
		n = 1000
	}

	var previous float64
	x := init
	var i int
	for i = 0; i < n; i++ {
		previous = x
		x = x - f(x)/Standard(x, f, precision)

		if math.Abs(x-previous) <= precision {
			break
		}
	}

	if i == (n - 1) {
		//Very likely we didn't find what we were looking for
		return 0.0, -1
	}

	return x, 0
}

/*
Steffensen is a method used to find the solution of an equation in the neighborhood
of a value. This method uses the Steffensen to find the solution. Note that choosing
the right init value is key to find the solution, it doesn't find all the solutions of
the equation.

If there is a division by zero, the process will stop and return -1 in the err return
value. Note that it doesn't mean that there is no solution but that the algorithm didn't
manage to find one.

First param init is an initial estimated value of the zero
Second param f is the function to solve
Third param is the number of iteration, it is optional and set to 1000 by default
Fourth param precision is the precision required, used to have an end condition
return the zero and zero in the error field or a -1 in the error field if it failed
*/
func Steffensen(init float64, f F, n int, precision float64) (float64, int) {
	if n == 0 {
		//ok let's try 1000
		n = 1000
	}
	p0 := init
	var p1, p2, p float64
	for i := 1; i < n; i++ {
		p1 = p0 + f(p0)
		p2 = p1 + f(p1)
		p = p2 - math.Pow(p2-p1, 2.0)/(p2-2*p1+p0)

		if math.Abs(p-p0) < precision {
			return p, 0
		}
		if math.IsNaN(p) {
			//Ok so we have the exact value
			return p0, 0
		}
		p0 = p
	}
	return p, -1
}
