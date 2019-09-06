package advmath

import (
	"fmt"
	"math"
	"testing"
)

/*
Testing of the derivative functions
*/

func tolerance(a, b, e float64) bool {
	// Multiplying by e here can underflow denormal values to zero.
	// Check a==b so that at least if a and b are small and identical
	// we say they match.
	if a == b {
		return true
	}
	d := a - b
	if d < 0 {
		d = -d
	}

	// note: b is correct (expected) value, a is actual value.
	// make error tolerance a fraction of b, not a.
	if b != 0 {
		e = e * b
		if e < 0 {
			e = -e
		}
	}
	return d < e
}

func close(a, b float64) bool      { return tolerance(a, b, 1e-14) }
func veryclose(a, b float64) bool  { return tolerance(a, b, 4e-16) }
func soclose(a, b, e float64) bool { return tolerance(a, b, e) }
func alike(a, b float64) bool {
	switch {
	case math.IsNaN(a) && math.IsNaN(b):
		return true
	case a == b:
		return math.Signbit(a) == math.Signbit(b)
	}
	return false
}
func alikeslices(a, b []float64) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestStandard(t *testing.T) {
	//New function
	x := func(w float64) float64 {
		return math.Log(w) / w
	}
	result := -(math.Log(2.0) - 1) / (2.0 * 2.0)
	z := Standard(2.0, x, 0.000000001)
	fmt.Printf("Standard(%g) = %g, want %g\n", 2.0, z, result)
	if !soclose(z, result, 0.000000001) {
		t.Errorf("Standard(%g) = %g, want %g", 2.0, z, result)
	}
}

func TestRidders(t *testing.T) {
	//New function
	x := func(w float64) float64 {
		return math.Log(w) / w
	}
	result := -(math.Log(2.0) - 1) / (2.0 * 2.0)
	z := Ridders(2.0, x, 0.000000001)
	fmt.Printf("Ridders(%g) = %g, want %g\n", 2.0, z, result)
	if !soclose(z, result, 0.000000001) {
		t.Errorf("Ridders(%g) = %g, want %g", 2.0, z, result)
	}
}

func TestSimpson(t *testing.T) {
	sup := 4.59
	inf := 2.87
	//New function
	x := func(w float64) float64 {
		return math.Log(w) / w
	}
	prim := func(j float64) float64 {
		return math.Log(j) * math.Log(j) / 2
	}
	result := prim(sup) - prim(inf)
	z, err := Simpson(inf, sup, x, 100000000)
	if err != nil {
		t.Errorf("Error while running Simpson method %v", err)
	}

	fmt.Printf("Simpson(%g, %g) = %g, want %g\n", inf, sup, z, result)

	if !soclose(z, result, 0.000001) {
		t.Errorf("Simpson(%g, %g) = %g, want %g", inf, sup, z, result)
	}
}

func TestSimpsonError(t *testing.T) {
	sup := 4.59
	inf := 2.87
	//New function
	x := func(w float64) float64 {
		return math.Log(w) / w
	}
	prim := func(j float64) float64 {
		return math.Log(j) * math.Log(j) / 2
	}
	result := prim(sup) - prim(inf)
	_, err := Simpson(inf, sup, x, 200001)
	if err != nil {
		fmt.Printf("Error while running Simpson method, expected result: %g, got error instead: %v\n", result, err)
	}
}

func TestTrapezoidal(t *testing.T) {
	sup := 4.59
	inf := 2.87
	//New function
	x := func(w float64) float64 {
		return math.Log(w) / w
	}
	prim := func(j float64) float64 {
		return math.Log(j) * math.Log(j) / 2
	}
	result := prim(sup) - prim(inf)
	z := Trapezoidal(inf, sup, x, 0, 0.000000000001)
	fmt.Printf("Trapezoidal(%g, %g) = %g, want %g\n", inf, sup, z, result)
	if !soclose(z, result, 0.000000000001) {
		t.Errorf("Trapezoidal(%g, %g) = %g, want %g", inf, sup, z, result)
	}
}

func TestRomberg(t *testing.T) {
	sup := 4.59
	inf := 2.87
	//New function
	x := func(w float64) float64 {
		return math.Log(w) / w
	}
	prim := func(j float64) float64 {
		return math.Log(j) * math.Log(j) / 2
	}
	result := prim(sup) - prim(inf)
	z := Romberg(inf, sup, x, 0, 0.0000000000001)
	fmt.Printf("Romberg(%g, %g) = %g, want %g\n", inf, sup, z, result)
	if !soclose(z, result, 0.0000000000001) {
		t.Errorf("Romberg(%g, %g) = %g, want %g", inf, sup, z, result)
	}
}

func TestSolve(t *testing.T) {
	y := func(x float64) float64 {
		return 7*math.Pow(x, 3.0) - 7*math.Pow(x, 5.0) + 3 - 3*math.Pow(x, 2.0)
	}
	result := 1.0
	z, err := Newton(0.6, y, 0, 0.000000001)
	fmt.Printf("Solve(%g) = %g, want %g\n", 0.6, z, result)
	if err != 0 {
		t.Errorf("Solve() = %g, want %g, returned error=%d", z, result, err)
	}
	if !soclose(z, result, 0.000000001) {
		t.Errorf("Solve() = %g, want %g", z, result)
	}
}

func TestSteffensen(t *testing.T) {
	y := func(x float64) float64 {
		return 7*math.Pow(x, 3.0) - 7*math.Pow(x, 5.0) + 3 - 3*math.Pow(x, 2.0)
	}
	result := 1.0
	z, err := Steffensen(0.3, y, 0, 0.0000000001)
	fmt.Printf("Steffensen(%g) = %g, want %g\n", 0.3, z, result)
	if err != 0 {
		t.Errorf("Solve() = %g, want %g, returned error =%d", z, result, err)
	}
	if !soclose(z, result, 0.0000000001) {
		t.Errorf("Solve() = %g, want %g", z, result)
	}
}

func TestGetRow(t *testing.T) {
	testMatrix := NewMatrix(3, 3)
	row1 := []float64{1, 2, 3}
	row2 := []float64{4, 5, 6}
	row3 := []float64{7, 8, 9}
	testMatrix.SetRow(0, row1)
	testMatrix.SetRow(1, row2)
	testMatrix.SetRow(2, row3)

	if !alikeslices(testMatrix.GetRow(1), row2) {
		t.Errorf("GetRow() Not the row expected %v, want %v", testMatrix.GetRow(1), row2)
	}
}

func TestTrace(t *testing.T) {
	testMatrix := NewMatrix(3, 3)
	row1 := []float64{1, 2, 3}
	row2 := []float64{4, 5, 6}
	row3 := []float64{7, 8, 9}
	testMatrix.SetRow(0, row1)
	testMatrix.SetRow(1, row2)
	testMatrix.SetRow(2, row3)

	result := 1.0 + 5.0 + 9.0
	calc, ok := testMatrix.Trace()
	fmt.Printf("Trace = %g, want %g\n", calc, result)

	if ok == nil && calc != result {
		t.Errorf("Trace() = %f, wanted %f", calc, result)
	}
}

func TestMultiply(t *testing.T) {
	testMatrixA := NewMatrix(2, 3)
	rowA1 := []float64{3, -2, 5}
	rowA2 := []float64{3, 0, 4}
	testMatrixA.SetRow(0, rowA1)
	testMatrixA.SetRow(1, rowA2)

	testMatrixB := NewMatrix(3, 2)
	rowB1 := []float64{2, 3}
	rowB2 := []float64{-9, 0}
	rowB3 := []float64{0, 4}
	testMatrixB.SetRow(0, rowB1)
	testMatrixB.SetRow(1, rowB2)
	testMatrixB.SetRow(2, rowB3)

	res, _ := testMatrixA.Multiply(testMatrixB)

	fmt.Printf("%v", res.M)
}

func TestDeterminant(t *testing.T) {
	testMatrix := NewMatrix(4, 4)
	row1 := []float64{3, 2, 1, -5}
	row2 := []float64{1, 5, -6, 3}
	row3 := []float64{-8, -6, 6, 3}
	row4 := []float64{1, 1, 8, -12}
	testMatrix.SetRow(0, row1)
	testMatrix.SetRow(1, row2)
	testMatrix.SetRow(2, row3)
	testMatrix.SetRow(3, row4)

	id := NewIdentity(4)
	fmt.Println(id)

	result := -50.0
	calc, ok := testMatrix.Determinant()
	fmt.Printf("Determinant = %g, want %g\n", calc, result)

	if ok == nil && !soclose(calc, result, 0.000000001) {
		t.Errorf("Determinant() = %f, wanted %f", calc, result)
	}
}

func TestInverse(t *testing.T) {
	testMatrix := NewMatrix(4, 4)
	row1 := []float64{1, 2, 3, 4}
	row2 := []float64{0, 2, 3, 4}
	row3 := []float64{1, 2, 0, 4}
	row4 := []float64{1, 0, 3, 4}
	testMatrix.SetRow(0, row1)
	testMatrix.SetRow(1, row2)
	testMatrix.SetRow(2, row3)
	testMatrix.SetRow(3, row4)

	det, _ := testMatrix.Determinant()
	fmt.Println(det)

	m, error := testMatrix.Inverse()
	fmt.Println(m)
	r, _ := testMatrix.Multiply(m)
	fmt.Println(r)

	fmt.Println(error)
}

func TestTranspose(t *testing.T) {
	/*testMatrix := NewMatrix(4, 5)
	row1 := []float64{1, 2, 3, 4, 5}
	row2 := []float64{0, 2, 3, 4, 5}
	row3 := []float64{1, 2, 0, 4, 5}
	row4 := []float64{1, 0, 3, 4, 5}
	testMatrix.SetRow(0, row1)
	testMatrix.SetRow(1, row2)
	testMatrix.SetRow(2, row3)
	testMatrix.SetRow(3, row4)*/
	testMatrix := NewMatrix(9, 9)
	row1 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	row2 := []float64{3, 2, 1, 4, 5, 6, 7, 8, 9}
	row3 := []float64{2, 1, 2, 4, 5, 6, 7, 8, 9}
	row4 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	row5 := []float64{3, 2, 1, 4, 5, 6, 7, 8, 9}
	row6 := []float64{2, 1, 2, 4, 5, 6, 7, 8, 9}
	row7 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	row8 := []float64{3, 2, 1, 4, 5, 6, 7, 8, 9}
	row9 := []float64{2, 1, 2, 4, 5, 6, 7, 8, 9}
	testMatrix.SetRow(0, row1)
	testMatrix.SetRow(1, row2)
	testMatrix.SetRow(2, row3)
	testMatrix.SetRow(3, row4)
	testMatrix.SetRow(4, row5)
	testMatrix.SetRow(5, row6)
	testMatrix.SetRow(6, row7)
	testMatrix.SetRow(7, row8)
	testMatrix.SetRow(8, row9)

	tr, _ := testMatrix.Transpose()

	fmt.Println(tr)
}
