package main

import "fmt"

type Matrix [][]int

func NewMatrix(rows, cols int) Matrix {
	matrix := make(Matrix, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}
	return matrix
}

// MultiplyInPlace modifies the receiver matrix (pointer receiver)
func (m *Matrix) MultiplyInPlace(other Matrix) error {
	if len((*m)[0]) != len(other) {
		return fmt.Errorf("incompatible matrix dimensions")
	}

	rows, cols := len(*m), len(other[0])
	result := NewMatrix(rows, cols)

	for i := range rows {
		for j := range cols {
			for k := range other {
				result[i][j] += (*m)[i][k] * other[k][j]
			}
		}
	}

	*m = result
	return nil
}

// Multiply returns a new matrix (value receiver)
func (m Matrix) Multiply(other Matrix) (Matrix, error) {
	if len(m[0]) != len(other) {
		return nil, fmt.Errorf("incompatible matrix dimensions")
	}

	rows, cols := len(m), len(other[0])
	result := NewMatrix(rows, cols)

	for i := range rows {
		for j := range cols {
			for k := range other {
				result[i][j] += m[i][k] * other[k][j]
			}
		}
	}

	return result, nil
}
