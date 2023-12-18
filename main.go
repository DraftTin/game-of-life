package main

import (
	"fmt"
	"math/rand"
)

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
	for i := 0; i < size; i++ {
		for j := 0; j < size-1; j++ {
			fmt.Print(universe[i][j])
		}
		fmt.Println(universe[i][size-1])
	}
}
