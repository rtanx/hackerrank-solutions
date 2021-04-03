package mathematics

/**
 * Problem : https://www.hackerrank.com/challenges/lowest-triangle/problem
 *
 * Domain : Mathematics
 * Sub Domain : Fundamentals
 * Difficulty : Easy
 * Author : Ridwan Maulana Tanjung
 */

// LowestTriangle calculate the height of the triangle based on its base and area
func LowestTriangle(base, area int) int {
	return (area / (base / 2))
}
