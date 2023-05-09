package cell

import (
	"testing"
)

func TestCellIsAlive(t *testing.T) {
    // Create a new alive cell and verify it is alive
    c1 := NewCell(true)
    if !c1.IsAlive() {
        t.Errorf("Expected cell to be alive, but it is not")
    }

    // Create a new dead cell and verify it is not alive
    c2 := NewCell(false)
    if c2.IsAlive() {
        t.Errorf("Expected cell to be dead, but it is alive")
    }
}

func TestNextGenerationCell(t *testing.T) {
    // Create a new alive cell with 2 live neighbors, expected to survive
    c1 := NewCell(true)
    nextGenCell := c1.NextGenerationCell(2)
    if !nextGenCell.IsAlive() {
        t.Errorf("Expected cell to survive with 2 live neighbors, but it died")
    }

    // Create a new alive cell with 4 live neighbors, expected to die
    c2 := NewCell(true)
    nextGenCell = c2.NextGenerationCell(4)
    if nextGenCell.IsAlive() {
        t.Errorf("Expected cell to die with 4 live neighbors, but it survived")
    }

    // Create a new dead cell with 3 live neighbors, expected to come to life
    c3 := NewCell(false)
    nextGenCell = c3.NextGenerationCell(3)
    if !nextGenCell.IsAlive() {
        t.Errorf("Expected cell to come to life with 3 live neighbors, but it remained dead")
    }

    // Create a new dead cell with 2 live neighbors, expected to remain dead
    c4 := NewCell(false)
    nextGenCell = c4.NextGenerationCell(2)
    if nextGenCell.IsAlive() {
        t.Errorf("Expected cell to remain dead with 2 live neighbors, but it came to life")
    }
}

