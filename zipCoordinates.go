func makeUniq(a []int) []int {
	sort.Ints(a)
	var unique []int
	m := map[int]bool{}
	for _, v := range a {
		if !m[v] {
			m[v] = true
			unique = append(unique, v)
		}
	}
	return unique
}
 
func lowerBound(array []int, element int) int {
	left := 0
	right := len(array)
	for ; right - left > 1; {
		mid := (left + right) >> 1
		if array[mid] <= element {
			left = mid
		} else {
			right = mid
		}
	}
	return left
}
 
func zip(a []int) (int, []int) {
	b := make([]int, len(a))
	copy(b, a)
	b = makeUniq(b)
 
	for index, element := range a {
		a[index] = lowerBound(b, element)
	}
	return len(b), a
}