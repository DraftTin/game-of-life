package main

import (
	"fmt"
	"math/rand"
	"time"
)

type CellState int

const (
	Dead CellState = iota
	Alive
)

func update(universe [][]string) {
	n := len(universe)
	newUniverse := make([][]string, len(universe))
	for i := 0; i < len(newUniverse); i++ {
		newUniverse[i] = make([]string, len(universe))
		for j := 0; j < len(universe); j++ {
			newUniverse[i][j] = " "
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if helper(universe, i, j) == Dead {
				newUniverse[i][j] = " "
			} else {
				newUniverse[i][j] = "O"
			}
		}
	}
	for i := 0; i < len(universe); i++ {
		copy(universe[i], newUniverse[i])
	}
}

func helper(universe [][]string, row, col int) CellState {
	neighbors := 0
	if row > 0 && universe[row-1][col] == "O" {
		neighbors++
	}
	if col > 0 && universe[row][col-1] == "O" {
		neighbors++
	}
	if row < len(universe)-1 && universe[row+1][col] == "O" {
		neighbors++
	}
	if col < len(universe)-1 && universe[row][col+1] == "O" {
		neighbors++
	}
	if col > 0 && row > 0 && universe[row-1][col-1] == "O" {
		neighbors++
	}
	if row < len(universe)-1 && col > 0 && universe[row+1][col-1] == "O" {
		neighbors++
	}
	if row < len(universe)-1 && col < len(universe)-1 && universe[row+1][col+1] == "O" {
		neighbors++
	}
	if row > 0 && col < len(universe)-1 && universe[row-1][col+1] == "O" {
		neighbors++
	}
	if neighbors < 2 || neighbors > 3 || (universe[row][col] == " " && neighbors == 2) {
		return Dead
	}

	return Alive

}

func show(universe [][]string) {
	size := len(universe)
	for i := 0; i < size; i++ {
		for j := 0; j < size-1; j++ {
			fmt.Print(universe[i][j])
		}
		fmt.Println(universe[i][size-1])
	}
}

func main() {
	var size int
	var seed int64
	fmt.Scan(&size, &seed)
	rand.Seed(seed)
	universe := make([][]string, size)

	for i := 0; i < size; i++ {
		universe[i] = make([]string, size)
		for j := 0; j < size; j++ {
			sign := rand.Intn(2)
			if sign == 1 {
				universe[i][j] = "O"
			} else {
				universe[i][j] = " "
			}
		}
	}
	show(universe)
	fmt.Print("\n\n")
	for {
		time.Sleep(time.Second)
		update(universe)
		show(universe)
		fmt.Print("\n\n")
	}
}
