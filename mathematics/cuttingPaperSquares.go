package mathematics

/**
 * Problem : https://www.hackerrank.com/challenges/p1-paper-cutting/problem
 *
 * Domain : Mathematics
 * Sub Domain : Fundamentals
 * Difficulty : Easy
 * Author : Ridwan Maulana Tanjung
 */

// CuttingPaperSquares calculate NxM papers into 1x1 pieces
func CuttingPaperSquares(n int32, m int32) int64 {
        return int64(n)*int64(m)-1
}

