package matrix

//一些列常用数据结构实现

import (
	"errors"
	"math"
)

//矩阵结构
//type Matrix interface {
//CountRows() int
//CountCols() int
//GetElm(i int, j int) float64
//trace() float64

//SetElm(i int, j int, v float64)
//add(*Matrix) error
//substract(*Matrix) error
//scale(float64)

//copy() []float64
//diagonalCopy() []float64
//}

type Matrix struct {
	// Number of rows
	rows int
	// Number of columns
	cols int
	// Matrix stored as a flat array: Aij = Elements[i*step + j]
	Elements []float64
	// Offset between rows
	step int
}

func MakeMatrix(Elements []float64, rows, cols int) *Matrix {
	A := new(Matrix)
	A.rows = rows
	A.cols = cols
	A.step = cols
	A.Elements = Elements

	return A
}
func (A *Matrix) CountRows() int {
	return A.rows
}

func (A *Matrix) CountCols() int {
	return A.cols
}

func (A *Matrix) GetElm(i int, j int) float64 {
	return A.Elements[i*A.step+j]
}

func (A *Matrix) SetElm(i int, j int, v float64) {
	A.Elements[i*A.step+j] = v
}

func (A *Matrix) diagonalCopy() []float64 {
	diag := make([]float64, A.cols)
	for i := 0; i < len(diag); i++ {
		diag[i] = A.GetElm(i, i)
	}
	return diag
}

func (A *Matrix) copy() *Matrix {
	B := new(Matrix)
	B.rows = A.rows
	B.cols = A.cols
	B.step = A.step

	B.Elements = make([]float64, A.cols*A.rows)

	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			B.Elements[i*A.step+j] = A.GetElm(i, j)
		}
	}
	return B
}

func (A *Matrix) trace() float64 {
	var tr float64 = 0
	for i := 0; i < A.cols; i++ {
		tr += A.GetElm(i, i)
	}
	return tr
}

func (A *Matrix) add(B *Matrix) error {
	if A.cols != B.cols && A.rows != B.rows {
		return errors.New("Wrong input sizes")
	}
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			A.SetElm(i, j, A.GetElm(i, j)+B.GetElm(i, j))
		}
	}

	return nil
}

func (A *Matrix) substract(B *Matrix) error {
	if A.cols != B.cols && A.rows != B.rows {
		return errors.New("Wrong input sizes")
	}
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			A.SetElm(i, j, A.GetElm(i, j)-B.GetElm(i, j))
		}
	}

	return nil
}

func (A *Matrix) scale(a float64) {
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			A.SetElm(i, j, a*A.GetElm(i, j))
		}
	}
}

func Add(A *Matrix, B *Matrix) *Matrix {
	result := MakeMatrix(make([]float64, A.cols*A.rows), A.cols, A.rows)

	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			result.SetElm(i, j, A.GetElm(i, j)+B.GetElm(i, j))
		}
	}

	return result
}

func Substract(A *Matrix, B *Matrix) *Matrix {
	result := MakeMatrix(make([]float64, A.cols*A.rows), A.cols, A.rows)

	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			result.SetElm(i, j, A.GetElm(i, j)-B.GetElm(i, j))
		}
	}

	return result
}

//题目描述 请编程实现矩阵乘法，并考虑当矩阵规模较大时的优化方法。
//根据wikipedia上的介绍：两个矩阵的乘法仅当第一个矩阵B的列数和另一个矩阵A的行数相等时才能定义。
//如A是m×n矩阵和B是n×p矩阵，它们的乘积AB是一个m×p矩阵，它的一个元素其中 1 ≤ i ≤ m, 1 ≤ j ≤ p。

//解法一、常规解法
//其实，通过前面的分析，我们已经很明显的看出，两个具有相同维数的矩阵相乘，其复杂度为O（n^3），参考代码如下：
func MultiplyCommon(A *Matrix, B *Matrix) *Matrix {
	result := MakeMatrix(make([]float64, A.cols*A.rows), A.cols, A.rows)

	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			sum := float64(0)
			for k := 0; k < A.cols; k++ {
				sum += A.GetElm(i, k) * B.GetElm(k, j)
			}
			result.SetElm(i, j, sum)
		}
	}

	return result
}

//解法二、Strassen算法
//在解法一中，我们用了3个for循环搞定矩阵乘法，但当两个矩阵的维度变得很大时，
//O（n^3）的时间复杂度将会变得很大，于是，我们需要找到一种更优的解法。
//一般说来，当数据量一大时，我们往往会把大的数据分割成小的数据，各个分别处理。
//遵此思路，如果丢给我们一个很大的两个矩阵呢，是否可以考虑分治的方法循序渐进处理各个小矩阵的相乘，
//因为我们知道一个矩阵是可以分成更多小的矩阵的。
func MultiplyStrassen(A, B *Matrix) *Matrix {
	n := A.CountRows()
	bigN := scaleSize(n)

	bigA := MakeMatrix(make([]float64, bigN*bigN), bigN, bigN)
	bigB := MakeMatrix(make([]float64, bigN*bigN), bigN, bigN)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			bigA.SetElm(i, j, A.GetElm(i, j))
			bigB.SetElm(i, j, B.GetElm(i, j))
		}
	}

	bigC := recurse(bigA, bigB)

	C := MakeMatrix(make([]float64, n*n), n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			C.SetElm(i, j, bigC.GetElm(i, j))
		}
	}

	return C
}

func recurse(A, B *Matrix) *Matrix {
	n := A.CountRows()

	newSize := n / 2

	if n < 2 {
		return MakeMatrix([]float64{A.GetElm(0, 0) * B.GetElm(0, 0)}, 1, 1)
	}

	A11 := MakeMatrix(make([]float64, newSize*newSize), newSize, newSize)
	A12 := MakeMatrix(make([]float64, newSize*newSize), newSize, newSize)
	A21 := MakeMatrix(make([]float64, newSize*newSize), newSize, newSize)
	A22 := MakeMatrix(make([]float64, newSize*newSize), newSize, newSize)

	B11 := MakeMatrix(make([]float64, newSize*newSize), newSize, newSize)
	B12 := MakeMatrix(make([]float64, newSize*newSize), newSize, newSize)
	B21 := MakeMatrix(make([]float64, newSize*newSize), newSize, newSize)
	B22 := MakeMatrix(make([]float64, newSize*newSize), newSize, newSize)

	for i := 0; i < newSize; i++ {
		for j := 0; j < newSize; j++ {
			A11.SetElm(i, j, A.GetElm(i, j))
			A12.SetElm(i, j, A.GetElm(i, j+newSize))
			A21.SetElm(i, j, A.GetElm(i+newSize, j))
			A22.SetElm(i, j, A.GetElm(i+newSize, j+newSize))

			B11.SetElm(i, j, B.GetElm(i, j))
			B12.SetElm(i, j, B.GetElm(i, j+newSize))
			B21.SetElm(i, j, B.GetElm(i+newSize, j))
			B22.SetElm(i, j, B.GetElm(i+newSize, j+newSize))
		}
	}

	a := MakeMatrix(make([]float64, newSize*newSize), newSize, newSize)
	b := MakeMatrix(make([]float64, newSize*newSize), newSize, newSize)

	// P1 = (A11 + A22) * (B11 + B22)
	a = Add(A11, A22)
	b = Add(B11, B22)
	P1 := recurse(a, b)

	// P2 = (A21 + A22) * (B11)
	a = Add(A21, A22)
	P2 := recurse(a, B11)

	// P3 = (A11) * (B12 - B11)
	b = Substract(B12, B22)
	P3 := recurse(A11, b)

	// P4 = A22 * (B21 - B22)
	b = Substract(B21, B11)
	P4 := recurse(A22, b)

	// P5 = (A11 + A12) * B22
	a = Add(A11, A12)
	P5 := recurse(a, B22)

	// P6 = (A21 - A11) * (B11 + B12)
	a = Substract(A21, A11)
	b = Add(B11, B12)
	P6 := recurse(a, b)

	// P7 = (A12 - A22) * (B21 + B22)
	a = Substract(A12, A22)
	b = Add(B21, B22)
	P7 := recurse(a, b)

	// Aggregates the result into C
	C12 := Add(P3, P5)
	C21 := Add(P2, P4)

	a = Add(P1, P4)
	b = Add(a, P7)
	C11 := Substract(b, P5)

	a = Add(P1, P3)
	b = Add(a, P6)
	C22 := Substract(b, P2)

	C := MakeMatrix(make([]float64, n*n), n, n)

	for i := 0; i < newSize; i++ {
		for j := 0; j < newSize; j++ {
			C.SetElm(i, j, C11.GetElm(i, j))
			C.SetElm(i, j+newSize, C12.GetElm(i, j))
			C.SetElm(i+newSize, j, C21.GetElm(i, j))
			C.SetElm(i+newSize, j+newSize, C22.GetElm(i, j))
		}
	}

	return C
}

func scaleSize(n int) int {
	log2 := math.Ceil(math.Log(float64(n)) / math.Log(float64(2)))
	return int(math.Pow(2, log2))
}
