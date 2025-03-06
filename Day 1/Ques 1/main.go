package main

import ( 
	"encoding/json" 
	"fmt"
)

type Matrix struct {
	Rows    int       `json:"rows"`
	Columns int       `json:"columns"`
	Elements [][]int  `json:"elements"`
}

func (m *Matrix) GetRows() int {
	return m.Rows
}

func (m *Matrix) GetColumns() int {
	return m.Columns
}

func (m *Matrix) SetElement(i, j, value int) {
	if i >= 0 && i < m.Rows && j >= 0 && j < m.Columns {
		m.Elements[i][j] = value
	}
}

func (m *Matrix) Add(other Matrix) (Matrix, error) {
	if m.Rows != other.Rows || m.Columns != other.Columns {
		return Matrix{}, fmt.Errorf("matrices dimensions do not match")
	}
	
	result := Matrix{
		Rows:    m.Rows,
		Columns: m.Columns,
		Elements: make([][]int, m.Rows),
	}
	
	for i := 0; i < m.Rows; i++ {
		result.Elements[i] = make([]int, m.Columns)
	}
	
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Columns; j++ {
			result.Elements[i][j] = m.Elements[i][j] + other.Elements[i][j]
		}
	}
	
	return result, nil
}

func (m *Matrix) PrintJSON() {
	matrixJSON, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling matrix to JSON:", err)
		return
	}
	fmt.Println(string(matrixJSON))
}

func main() {
	matrix1 := Matrix{
		Rows:    2,
		Columns: 3,
		Elements: [][]int{
			{1, 2, 3},
			{4, 5, 6},
		},
	}

	matrix2 := Matrix{
		Rows:    2,
		Columns: 3,
		Elements: [][]int{
			{7, 8, 9},
			{10, 11, 12},
		},
	}

	result, err := matrix1.Add(matrix2)
	if err != nil {
		fmt.Println("Error adding matrices:", err)
		return
	}

	fmt.Println("Matrix 1:")
	matrix1.PrintJSON()
	fmt.Println("Matrix 2:")
	matrix2.PrintJSON()
	fmt.Println("Result of Addition:")
	result.PrintJSON()
}
