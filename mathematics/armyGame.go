package mathematics

/**
 * Problem : https://www.hackerrank.com/challenges/game-with-cells/problem
 *
 * Domain : Mathematics
 * Sub Domain : Fundamentals
 * Difficulty : Easy
 * Author : Ridwan Maulana Tanjung
 */

// GameWithCells calculate the minimum amount of supplies dropped for military bases based on the strategic point for each cells
func GameWithCells(n int32, m int32) int32 {
	return ((n + n%2) * (m + m%2)) / 4
}
