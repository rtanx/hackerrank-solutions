package algorithms

func migratoryBirds(arr []int32) int32 {
	var t [5]int32
	for _, v := range arr {
		t[v-1] += 1
	}
	var r int32
	for i := int32(0); i < int32(len(t)-1); i++ {
		if t[r] < t[i+1] {
			r = i + 1
		}
	}
	return (r + 1)
}
