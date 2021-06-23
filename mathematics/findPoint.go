package mathematics

/**
 * Problem : https://www.hackerrank.com/challenges/find-point/problem
 */

// FindPoint find the point reflection of P(Px,Py) as reflected to Q(Qx,Qy)
func FindPoint(px int32, py int32, qx int32, qy int32) []int32 {
	rx := (2 * qx) - px
	ry := (2 * qy) - py
	return []int32{rx, ry}
}
