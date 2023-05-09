package utility

import (
	"conwaysgameoflife/cell"
)

func IsCellAliveAt(rowNo, columnNo int, matrix [][]cell.Cell) bool {

	if rowNo < 0 || columnNo < 0 || rowNo >= len(matrix) || columnNo >= len(matrix[0]) {
		return false
	}

	return matrix[rowNo][columnNo].IsAlive()

}

func NumberOfLiveNeighboursAt(rowNo, columnNo int, matrix [][]cell.Cell) int {

	count := 0
	if IsCellAliveAt(rowNo+1, columnNo, matrix) {
		count += 1
	}
	if IsCellAliveAt(rowNo-1, columnNo, matrix) {
		count += 1
	}
	if IsCellAliveAt(rowNo, columnNo+1, matrix) {
		count += 1
	}
	if IsCellAliveAt(rowNo, columnNo-1, matrix) {
		count += 1
	}
	if IsCellAliveAt(rowNo+1, columnNo+1, matrix) {
		count += 1
	}
	if IsCellAliveAt(rowNo+1, columnNo-1, matrix) {
		count += 1

	}
	if IsCellAliveAt(rowNo-1, columnNo+1, matrix) {
		count += 1
	}
	if IsCellAliveAt(rowNo-1, columnNo-1, matrix) {
		count += 1
	}
	return count

}
