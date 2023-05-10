package cell

import (
	"testing"
)

func TestCellIsAlive(t *testing.T) {
    
    c1 := NewCell(true)
    if !c1.IsAlive() {
        t.Errorf("Expected cell to be alive, but it is not")
    }

    c2 := NewCell(false)
    if c2.IsAlive() {
        t.Errorf("Expected cell to be dead, but it is alive")
    }
}

func TestNextGenerationCell(t *testing.T) {
   
    c1 := NewCell(true)
    nextGenCell := c1.NextGenerationCell(2)
    if !nextGenCell.IsAlive() {
        t.Errorf("Expected cell to survive with 2 live neighbors, but it died")
    }

    c2 := NewCell(true)
    nextGenCell = c2.NextGenerationCell(4)
    if nextGenCell.IsAlive() {
        t.Errorf("Expected cell to die with 4 live neighbors, but it survived")
    }

    c3 := NewCell(false)
    nextGenCell = c3.NextGenerationCell(3)
    if !nextGenCell.IsAlive() {
        t.Errorf("Expected cell to come to life with 3 live neighbors, but it remained dead")
    }

    c4 := NewCell(false)
    nextGenCell = c4.NextGenerationCell(2)
    if nextGenCell.IsAlive() {
        t.Errorf("Expected cell to remain dead with 2 live neighbors, but it came to life")
    }
}

