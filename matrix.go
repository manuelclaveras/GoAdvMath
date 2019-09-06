package advmath

/*
Matrix is a standard mathematical array of numbers
*/
type Matrix struct {
	NumberOfRows    uint
	NumberOfColumns uint
	M               []float64
}

/*
NewMatrix is a method to create a new matrix type. By default
when created the matrix is filled with float64 default value
(which is 0.0)
First parameter is the number of rows
Second parameter is the number of columns
*/
func NewMatrix(rows, cols uint) *Matrix {
	m := new(Matrix)
	m.NumberOfRows = rows
	m.NumberOfColumns = cols
	m.M = make([]float64, rows*cols)
	return m
}

/*
NewIdentity is a method to create an identity square matrix, hence only one parameter
the number of rows.
First parameter is the number of rows and columns
*/
func NewIdentity(rows uint) *Matrix {
	i := new(Matrix)
	i.NumberOfRows = rows
	i.NumberOfColumns = rows
	i.M = make([]float64, rows*rows)

	var j, k uint
	for j = 0; j < i.NumberOfRows; j++ {
		for k = 0; k < i.NumberOfRows; k++ {
			if j == k {
				i.M[j*i.NumberOfColumns+k] = 1.0
			}
		}
	}

	return i
}

/*
IsSquare is a method to find if a matrix is a square matrix or not.
This is mainly used because some methods cannot work with a non square
matrix.
*/
func (m Matrix) IsSquare() bool {
	return m.NumberOfColumns == m.NumberOfRows
}

/*
Get is a method to retrieve the content of a matrix at the given
row and column.
It returns the value found.
*/
func (m Matrix) Get(row uint, column uint) float64 {
	return m.M[row*m.NumberOfColumns+column]
}

/*
GetRow is method used to return the specified row of a matrix. It takes the
row number as an input. Note that rowNumber should start at 0.
*/
func (m Matrix) GetRow(rowNumber uint) []float64 {
	row := make([]float64, m.NumberOfColumns)

	var cols uint
	for cols = 0; cols < m.NumberOfColumns; cols++ {
		row[cols] = m.M[rowNumber*m.NumberOfColumns+cols]
	}

	return row
}

/*
GetColumn is a method used to retrieve a specific column of the matrix.
Note that colNumber should start at 0 as always.
First parameter is the column number
*/
func (m Matrix) GetColumn(colNumber uint) []float64 {
	col := make([]float64, m.NumberOfRows)

	var rows uint
	for rows = 0; rows < m.NumberOfRows; rows++ {
		col[rows] = m.M[rows*m.NumberOfColumns+colNumber]
	}

	return col
}

/*
Set is a method to set the value at the given row and column
it doesn't return anything but changes the underlying matrix.
*/
func (m *Matrix) Set(row uint, column uint, value float64) {
	m.M[row*m.NumberOfColumns+column] = value
}

/*
SetRow is a method to set the value at the given row
it doesn't return anything but changes the underlying matrix.
*/
func (m *Matrix) SetRow(rowNumber uint, row []float64) *Matrix {
	var cols uint
	for cols = 0; cols < m.NumberOfColumns; cols++ {
		m.M[rowNumber*m.NumberOfColumns+cols] = row[cols]
	}
	return m
}

/*
SubMatrix is a method that returns a sub matrix of the original
matrix starting from row and col taking the number of rows and
columns specified.
For instance, if we have a matrix:
	[1 2 3]
	[4 5 6]
	[7 8 9]
and SubMatrix is called with the following parameters:
- 1
- 1
- 2
- 2
it will return:
	[5 6]
	[8 9]
*/
func (m *Matrix) SubMatrix(row, col, numberRows, numberCols uint) *Matrix {
	sub := NewMatrix(numberRows, numberCols)
	sub.M = m.M[row*m.NumberOfColumns+col : row*m.NumberOfColumns+col+(numberRows)*m.NumberOfColumns+numberCols]
	sub.NumberOfColumns = numberCols
	sub.NumberOfRows = numberRows
	return sub
}

/*
Multiply is a method to multiply the matrix by the given matrix.
Since multiplication is not commutative it means that:

a.Multiply(b) will result in A*B

First parameter is the matrix used for the multiplication
*/
func (m Matrix) Multiply(in *Matrix) (*Matrix, error) {
	//Columns and rows must match
	if m.NumberOfColumns != in.NumberOfRows {
		return nil, &MathError{
			code: errorCannotMultiply,
		}
	}

	result := NewMatrix(m.NumberOfRows, in.NumberOfColumns)

	var i, j, k uint
	for i = 0; i < m.NumberOfRows; i++ {
		for j = 0; j < in.NumberOfColumns; j++ {
			for k = 0; k < m.NumberOfColumns; k++ {
				result.M[i*result.NumberOfColumns+j] += m.M[i*m.NumberOfColumns+k] * in.M[k*in.NumberOfColumns+j]
			}
		}
	}
	return result, nil
}

/*
ScalarMultiply is a method to multiply a matrix by a scalar.
First parameter is a scalar used to multiply
*/
func (m Matrix) ScalarMultiply(scal float64) *Matrix {
	result := NewMatrix(m.NumberOfColumns, m.NumberOfRows)

	var row, col uint
	for row = 0; row < m.NumberOfRows; row++ {
		for col = 0; col < m.NumberOfColumns; col++ {
			result.M[row*result.NumberOfColumns+col] *= scal
		}
	}

	return result
}

/*
Add is a method to add a matrix to another matrix
First parameter is a matrix to add
*/
func (m Matrix) Add(in *Matrix) (*Matrix, error) {
	if in.NumberOfColumns != m.NumberOfColumns || in.NumberOfRows != m.NumberOfRows {
		return nil, &MathError{
			code: errorCannotAdd,
		}
	}

	result := NewMatrix(m.NumberOfColumns, m.NumberOfRows)

	var row, col uint
	for row = 0; row < m.NumberOfRows; row++ {
		for col = 0; col < m.NumberOfColumns; col++ {
			result.M[row*result.NumberOfColumns+col] = m.M[row*m.NumberOfColumns+col] + in.M[row*in.NumberOfColumns+col]
		}
	}

	return result, nil
}

/*
Subtract is a method to subtract a matrix with another one.
First parameter is the matrix to subtract
*/
func (m Matrix) Subtract(in *Matrix) (*Matrix, error) {
	if in.NumberOfColumns != m.NumberOfColumns || in.NumberOfRows != m.NumberOfRows {
		return nil, &MathError{
			code: errorCannotAdd,
		}
	}
	return m.Add(in.Neg())
}

/*
Neg is a method to return the negative version of a matrix. i.e multiply the underlying matrix by -1
*/
func (m Matrix) Neg() *Matrix {
	return m.ScalarMultiply(-1.0)
}

/*
Trace is a method to compute the trace of a square matrix, i.e. adding the elements
on the diagonal of the matrix. If it is not a square matrix, it just returns 0.0 and an
error indicating that trace cannot be computed on a non-square matrix.
It takes no parameters and returns the sum.
*/
func (m Matrix) Trace() (float64, error) {
	//Check if it is possible to find one
	if !m.IsSquare() {
		return 0.0, &MathError{
			code: errorNonSquareMatrix,
		}
	}
	var trace float64
	var column uint
	var row uint
	for row = 0; row < m.NumberOfRows; row++ {
		trace += m.Get(row, column)
		column++
	}
	return trace, nil
}

/*
LUDecomposition is a method to create the LU decomposition of a square matrix. It provides
a lower triangular matrix with ones on the diagonal and an upper triangular matrix.
First return value is the lower triangular matrix
Second return value is the upper triangular matrix
Third return value is the error that can occur in the process (if non square matrix)
*/
func (m Matrix) LUDecomposition() (*Matrix, *Matrix, error) {
	if !m.IsSquare() {
		return nil, nil, &MathError{
			code: errorNonSquareMatrix,
		}
	}

	l := NewMatrix(m.NumberOfRows, m.NumberOfColumns)
	u := NewMatrix(m.NumberOfRows, m.NumberOfColumns)

	// Decomposing matrix into Upper and Lower
	// triangular matrix
	n := m.NumberOfColumns
	var i, j, k uint
	for i = 0; i < n; i++ {
		// Upper Triangular
		for k = i; k < n; k++ {
			// Summation of L(i, j) * U(j, k)
			sum := 0.0
			for j = 0; j < i; j++ {
				sum += (l.M[i*l.NumberOfColumns+j] * u.M[j*u.NumberOfColumns+k])
			}
			// Evaluating U(i, k)
			u.M[i*u.NumberOfColumns+k] = m.M[i*m.NumberOfColumns+k] - sum
		}
		// Lower Triangular
		for k = i; k < n; k++ {
			if i == k {
				//Set the diagonal to ones
				l.M[i*l.NumberOfColumns+i] = 1.0
			} else {
				// Summation of L(k, j) * U(j, i)
				sum := 0.0
				for j = 0; j < i; j++ {
					sum += (l.M[k*l.NumberOfColumns+j] * u.M[j*u.NumberOfColumns+i])
				}
				// Evaluating L(k, i)
				l.M[k*l.NumberOfColumns+i] = (m.M[k*m.NumberOfColumns+i] - sum) / u.M[i*u.NumberOfColumns+i]
			}
		}
	}

	return l, u, nil
}

/*
Determinant is a method to compute the determinant of a square matrix. It uses the
LU decomposition to compute the value
*/
func (m Matrix) Determinant() (float64, error) {
	if !m.IsSquare() {
		return 0.0, &MathError{
			code: errorNonSquareMatrix,
		}
	}

	_, u, err := m.LUDecomposition()
	if err != nil {
		return 0.0, err
	}

	//We just need to compute the determinant of the upper matrix and since it's a triangular matrix that's just
	//mulitplying the elements on the diagonal
	det := 1.0
	var column uint
	var row uint
	for row = 0; row < m.NumberOfRows; row++ {
		det *= u.Get(row, column)
		column++
	}

	return det, nil
}

func (m Matrix) determinantLU() (float64, *Matrix, *Matrix, error) {
	if !m.IsSquare() {
		return 0.0, nil, nil, &MathError{
			code: errorNonSquareMatrix,
		}
	}

	l, u, err := m.LUDecomposition()
	if err != nil {
		return 0.0, nil, nil, err
	}

	//We just need to compute the determinant of the upper matrix
	//and since it's a triangular matrix that's just
	//mulitplying the elements on the diagonal
	det := 1.0
	var column uint
	var row uint
	for row = 0; row < m.NumberOfRows; row++ {
		det *= u.Get(row, column)
		column++
	}

	return det, l, u, nil
}

/*
Inverse is a method to compute the inverse of a square matrix. If this method is called on a
non square matrix then an error will be returned.
This method uses the LU decomposition to compute the inverse:

A*A^-1 = I <=> (L*U)*[a1 a2 ... aN] = [e1 e2 ... eN]

This is like solving sets of equations for :

L*y = en
U*an = y

That should be easy since we have triangular matrices. Once we've done that, all the an are simply
the inverse of our A matrix.
*/
func (m Matrix) Inverse() (*Matrix, error) {
	//First get the LU decomposition and the determinant
	det, l, u, error := m.determinantLU()
	if error != nil {
		return nil, error
	}

	if det == 0.0 {
		//Ok cannot find inverse
		return nil, &MathError{
			code: errorNotInversible,
		}
	}

	id := NewIdentity(m.NumberOfRows)
	y := NewMatrix(m.NumberOfRows, m.NumberOfColumns)

	//Let solve L*Y = I
	var i, j, k int
	var sum float64
	for k = 0; k < int(y.NumberOfColumns); k++ {
		y.M[k] = id.GetColumn(uint(k))[0] / l.Get(0, 0)
		for i = 1; i < int(l.NumberOfRows); i++ {
			for j = 0; j < i; j++ {
				sum += l.Get(uint(i), uint(j)) * y.M[uint(j)*y.NumberOfColumns+uint(k)]
			}
			y.M[uint(i)*y.NumberOfColumns+uint(k)] = (id.Get(uint(i), uint(k)) - sum) / l.Get(uint(i), uint(i))
			sum = 0.0
		}
	}

	x := NewMatrix(m.NumberOfRows, m.NumberOfColumns)
	var sum2 float64
	//Now let solve U*X = Y
	for n := 0; n < int(x.NumberOfColumns); n++ {
		x.Set(x.NumberOfRows-1, x.NumberOfColumns-1-uint(n), y.GetColumn(x.NumberOfColumns - 1 - uint(n))[int(y.NumberOfRows)-1]/u.Get(x.NumberOfRows-1, x.NumberOfColumns-1))
		for o := int(x.NumberOfColumns) - 2; o >= 0; o-- {
			for p := o + 1; p < int(x.NumberOfRows); p++ {
				sum2 += u.Get(uint(o), uint(p)) * x.Get(uint(p), x.NumberOfColumns-1-uint(n))
			}

			x.Set(uint(o), x.NumberOfColumns-1-uint(n), (y.Get(uint(o), x.NumberOfColumns-1-uint(n))-sum2)/u.Get(uint(o), uint(o)))
			sum2 = 0.0
		}
	}

	return x, nil
}

/*
QRDecomposition is a method to compute a QR decomposition of the matrix. The goal is to create
a matrix Q and a matrix R so that:
- A = Q*R
- Q is an orthogonal matrix
- R is a upper diagonal matrix
*/
func (m Matrix) QRDecomposition() (*Matrix, error) {
	return nil, nil
}

/*
Transpose is a method to compute the transposition of a matrix. This method uses 2 different
methods to compute it depending on whether the matrix is a square or not:

- Square matrices: first copy diagonal and then iterate to swap the values
- Non-square matrices: pseudo in place transpose, an algorithm with O(1) space
*/
func (m Matrix) Transpose() (*Matrix, error) {
	if !m.IsSquare() {
		return m.nonSquareTranspose(), nil
	}

	ret := NewMatrix(m.NumberOfColumns, m.NumberOfRows)
	//copy diagonal
	for k := 0; k < int(m.NumberOfRows); k++ {
		ret.Set(uint(k), uint(k), m.Get(uint(k), uint(k)))
	}

	for i := 0; i <= int(m.NumberOfRows)-2; i++ {
		for j := i + 1; j <= int(m.NumberOfRows)-1; j++ {
			ret.Set(uint(j), uint(i), m.Get(uint(i), uint(j)))
			ret.Set(uint(i), uint(j), m.Get(uint(j), uint(i)))
		}
	}
	return ret, nil
}

func (m Matrix) nonSquareTranspose() *Matrix {
	ret := NewMatrix(m.NumberOfColumns, m.NumberOfRows)
	var start, j, i int64
	var tmp float64

	for start = 0; start <= int64(m.NumberOfRows*m.NumberOfColumns-1); start++ {
		j = start
		i = 0
		for ok := true; ok; {
			i++
			j = (j%int64(m.NumberOfRows))*int64(m.NumberOfColumns) + j/int64(m.NumberOfRows)
			ok = (j > start)
		}

		j = start
		tmp = m.M[j]
		for ok := true; ok; {
			i = (j%int64(m.NumberOfRows))*int64(m.NumberOfColumns) + j/int64(m.NumberOfRows)
			if ret.M[j] = m.M[i]; i == start {
				ret.M[j] = tmp
			}
			j = i
			ok = (j > start)
		}
	}

	return ret
}

/*
Cofactor is a method to compute the cofactors
*/
func (m Matrix) Cofactor() (*Matrix, error) {
	if !m.IsSquare() {
		return nil, &MathError{
			code: errorNonSquareMatrix,
		}
	}

	//c := NewMatrix(m.NumberOfRows, m.NumberOfColumns)
	//n := m.NumberOfColumns

	return nil, nil
}
