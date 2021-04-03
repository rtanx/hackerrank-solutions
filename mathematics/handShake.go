package mathematics

/**
 * Problem : https://www.hackerrank.com/challenges/handshake/problem
 *
 * Domain : Mathematics
 * Sub Domain : Fundamentals
 * Difficulty : Easy
 * Author : Ridwan Maulana Tanjung
 */

// HandShake calculate how many handshakes are performed in a group of n people
func HandShake(n int32) int32 {
	return (n * (n - 1)) / 2
}
