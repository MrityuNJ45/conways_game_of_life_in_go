package utility

import (
	"conwaysgameoflife/cell"
	"testing"
)

func TestIsCellAliveAt(t *testing.T){

	matrix := [][]cell.Cell{
		{cell.NewCell(false), cell.NewCell(true)},
		{cell.NewCell(true), cell.NewCell(false)},
	}
	
	result := IsCellAliveAt(0,1,matrix);

	if result != true {
		t.Errorf("Expected to get true as cell is alive there, but got false")
	}
	
	result = IsCellAliveAt(0,0,matrix);

	if result == true {
		t.Errorf("Expected to get false as cell is dead there, but got true")
	}


}


func TestNumberOfLiveNeighboursAt(t *testing.T){

	matrix := [][]cell.Cell{
		{cell.NewCell(false), cell.NewCell(true)},
		{cell.NewCell(true), cell.NewCell(true)},
	}

	result := NumberOfLiveNeighboursAt(0,0, matrix);

	if result != 3 {
		t.Errorf("Expected to get 3 live neighbouring cells, but got %v" , result)
	}

	result = NumberOfLiveNeighboursAt(0,1,matrix);

	if result == 3 {
		t.Errorf("Expected to get 2 live neigbouring cells, but got 3")
	}

}
