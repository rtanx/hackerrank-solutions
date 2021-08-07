package algorithms

func Birthday(s []int32, d int32, m int32) int32 {
	var w int32
	for i, v := range s {
		itv := int32(1)
		for j := (i + 1); j < len(s); j++ {
			v += s[j]
			itv++
			if itv == m {
				break
			}
		}
		if v == d {
			w++
		}
	}
	return w
}
