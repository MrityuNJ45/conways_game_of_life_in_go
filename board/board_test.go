package board

import (
	"conwaysgameoflife/cell"
	
	"testing"
)

func TestNewBoard(t *testing.T) {
	
	rows := 3
	columns := 4
	board, err := NewBoard(rows, columns)
	if err != nil {
		t.Errorf("Error creating board: %v", err)
	}
	if board.rows != rows {
		t.Errorf("Expected rows to be %v but got %v", rows, board.rows)
	}
	if board.columns != columns {
		t.Errorf("Expected columns to be %v but got %v", columns, board.columns)
	}
	if len(board.matrix) != rows {
		t.Errorf("Expected matrix to have %v rows but got %v", rows, len(board.matrix))
	}
	if len(board.matrix[0]) != columns {
		t.Errorf("Expected matrix to have %v columns but got %v", columns, len(board.matrix[0]))
	}

	
	_, err = NewBoard(-1, 5)
	if err == nil {
		t.Errorf("Expected error for negative rows")
	}
	_, err = NewBoard(5, 0)
	if err == nil {
		t.Errorf("Expected error for zero columns")
	}
}

func TestEquals(t *testing.T) {

	board, _ := NewBoard(3, 3)
	anotherBoard, _ := NewBoard(3, 3)

	if !(board.Equals(anotherBoard)) {
		t.Errorf("Expect boards to equal, but they were not equal")
	}

	anotherBoard, _ = NewBoard(4, 4)
	if board.Equals(anotherBoard) {
		t.Errorf("Expect boards to be unequal, but they were equal")
	}

	board3, _ := NewBoard(3, 3)
	board4, _ := NewBoard(3, 3)

	matrix := make([][]cell.Cell, board3.rows)
	for row := range matrix {
		matrix[row] = make([]cell.Cell, board3.columns)
		for column := range matrix[row] {
			if row%2 == 0 {
				matrix[row][column] = cell.NewCell(true)
			} else {
				matrix[row][column] = cell.NewCell(false)
			}
		}
	}
	board3.matrix = matrix

	if board3.Equals(board4) {
		t.Errorf("Expect boards to be unequal, but they were equal")
	}

}

func TestPoplulate(t *testing.T) {

	board, _ := NewBoard(3, 3)

	populatedBoard := board.Populate()
	
	if board.Equals(populatedBoard) {
		t.Errorf("Expects boards to be unequal, but they are equal")
	}

}

func TestNextGenerationBoard(t *testing.T) {

	board, _ := NewBoard(3,3);

	populatedBoard := board.Populate();
	nextGenerationBoard := populatedBoard.NextGenerationBoard();

	if populatedBoard.Equals(nextGenerationBoard){
		t.Errorf("Expected boards to be different, but they are equal");
	}

}



