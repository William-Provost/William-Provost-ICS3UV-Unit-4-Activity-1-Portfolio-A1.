/**
 * @author William Provost
 * @version 1.0.0
 * @date 2025-11-25
 * @fileoverview This program prints a 10x10 grid.
 * Odd rows contain X's and even rows contain O's.
 */

// loop through 10 rows
for (let row = 1; row <= 10; row++) {
  let line: string = "";

  // loop through 10 columns
  for (let col = 1; col <= 10; col++) {
    if (row % 2 === 1) {
      line += "X ";
    } else {
      line += "O ";
    }
  }

  console.log(line);
}

console.log("\nDone.");
