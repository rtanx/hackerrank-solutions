package algorithms

func Kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	var Y, N string = "YES", "NO"
	if v1 == v2 {
		return N
	}
	if (v1 < v2 && x1 < x2) || (v1 > v2 && x1 > x2) {
		return N
	}
	if (x2-x1)%(v1-v2) == 0 {
		return Y
	}
	return N
}
