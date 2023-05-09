package main

import (
	"conwaysgameoflife/board"
	"fmt"
)

func main() {

	var rows, columns int

	fmt.Println("Enter the number of rows and columns : ")
	fmt.Scanf("%v %v", &rows, &columns)

	board, _ := board.NewBoard(rows, columns)
	board.Print()
	fmt.Println("============================")

	for true {
		fmt.Println("Enter 1 to populate.")
		fmt.Println("Enter 2 to generate next life.")
		fmt.Println("Enter 3 to exit.")

		var input int
		fmt.Scanf("%v", &input)

		if input == 1 {
			board = board.Populate()
			board.Print()
		}
		if input == 2 {
			board = board.NextGenerationBoard()
			board.Print()
		}
		fmt.Println("============================")
		if input == 3 {
			fmt.Println("Thank you for playing thr game...")
			break
		}

	}

}
