package advmath

const (
	//When we try to divide by zero
	errorDivisionByZero = 1
	//Error when the matrix is not a square matrix, some operations
	//only support square matrix
	errorNonSquareMatrix = 2
	//Error when the matrix is not initialized, i.e. did not use the
	errorMatrixIsNil = 3
	//Error when multiplication is required but matrices do not match
	errorCannotMultiply = 4
	//Error when multiplication is required but matrices do not match
	errorCannotAdd = 5
	//Error when we cannot find an inverse for the matrix
	errorNotInversible = 6
)

/*
MathError is the error type used throughout the library
*/
type MathError struct {
	code int
	s    string
}

/*
Error returns the description of the error
*/
func (e *MathError) Error() string {
	if e.code != 0 {
		switch e.code {
		case errorDivisionByZero:
			return "Tried to divide by zero"
		case errorNonSquareMatrix:
			return "Tried to do an operation on a non square matrix"
		case errorMatrixIsNil:
			return "Matrix is empty and has not been initialized with the rigth method"
		case errorCannotMultiply:
			return "Number of columns of first do not match the number of rows of second matrix"
		case errorCannotAdd:
			return "Can only add matrices of same size"
		case errorNotInversible:
			return "Matrix is not inversible"
		}
	}
	return e.s
}
