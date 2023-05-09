package board

import (
	"conwaysgameoflife/cell"
	"conwaysgameoflife/utility"
	"errors"
	"fmt"
	"math/rand"
)

type Board struct {
	rows    int
	columns int
	matrix  [][]cell.Cell
}

func NewBoard(rows, columns int) (Board, error) {
	if rows <= 0 || columns <= 0 {
		return Board{}, errors.New("Rows and Columns must be above 0")
	}

	matrix := make([][]cell.Cell, rows)
	for i := range matrix {
		matrix[i] = make([]cell.Cell, columns)
		for j := range matrix[i] {
			matrix[i][j] = cell.NewCell(false)
		}
	}

	return Board{rows: rows, columns: columns, matrix: matrix}, nil
}

func (b Board) Equals(otherBoard Board) bool {

	if b.rows != otherBoard.rows || b.columns != otherBoard.columns {
		return false
	}

	for row := 0; row < b.rows; row++ {
		for column := 0; column < b.columns; column++ {
			if b.matrix[row][column].IsAlive() != otherBoard.matrix[row][column].IsAlive() {
				return false
			}
		}
	}
	return true

}

func (b Board) Populate() Board {

	matrix := make([][]cell.Cell, b.rows)
	for i := range matrix {
		matrix[i] = make([]cell.Cell, b.columns)
		for j := range matrix[i] {
			random := rand.Intn(2)

			if random == 1 {
				matrix[i][j] = cell.NewCell(true)
			} else {
				matrix[i][j] = cell.NewCell(false)
			}

		}
	}
	return Board{rows: b.rows, columns: b.columns, matrix: matrix}

}

func (b Board) NextGenerationBoard() Board {

	matrix := make([][]cell.Cell, b.rows)
	for row := range matrix {
		matrix[row] = make([]cell.Cell, b.columns)
		for column := range matrix[row] {
			numOfLiveNeighbours := utility.NumberOfLiveNeighboursAt(row, column, b.matrix)
			matrix[row][column] = b.matrix[row][column].NextGenerationCell(numOfLiveNeighbours)
		}
	}
	return Board{rows: b.rows, columns: b.columns, matrix: matrix}
}


func (b Board) Print() {

	matrix := b.matrix;

	for row := range matrix {
		for column := range matrix[row] {
			cellAt := matrix[row][column];
			if(cellAt.IsAlive()){
				fmt.Print("* ");
				continue
			}
			fmt.Print("_ ")
		}
		fmt.Println()
	}
	
}