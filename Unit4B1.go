// Author: William Provost
// Version: 1.0.0
// Date: 2025-11-25
// Fileoverview: This program prints a 10x10 grid.
// Odd rows contain X's and even rows contain O's.

package main

import "fmt"

func main() {

	// loop through 10 rows
	for row := 1; row <= 10; row++ {
		line := ""

		// loop through 10 columns
		for col := 1; col <= 10; col++ {
			if row%2 == 1 {
				line += "X "
			} else {
				line += "O "
			}
		}

		fmt.Println(line)
	}

	fmt.Println("\nDone.")
}
