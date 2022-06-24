type segmentTree struct {
	n int
	t []int
}
 
func NewsegmentTree(n int) *segmentTree {
	st := new(segmentTree)
	st.n = n
	st.t = make([]int, 4 * n)
	return st
}
 
func (st segmentTree) build(a []int, tl int, tr int, v int) {
	if tl == tr {
		st.t[v] = a[tl]
		return
	}
	var mid = tl + tr >> 1
	st.build(a, tl, mid, v * 2)
	st.build(a, mid + 1, tr, v * 2 + 1)
	st.t[v] = st.t[v * 2] + st.t[v * 2 + 1]
	return
}
 
func (st segmentTree) get(l int, r int, tl int, tr int, v int) int {
	if tl > r || tr < l {
		return 0
	}
	if tl >= l && tr <= r {
		return st.t[v]
	}
	mid := (tl + tr) >> 1
	return st.get(l, r, tl, mid, v * 2) + st.get(l, r, mid + 1, tr, v * 2 + 1)
}
 
func (st segmentTree) update(pos int, val int, tl int, tr int, v int) {
	if tl == tr {
		st.t[v] += val
		return
	}
	var mid = (tl + tr) >> 1
	if pos <= mid {
		st.update(pos, val, tl, mid, v * 2)
	} else {
		st.update(pos, val, mid + 1, tr, v * 2 + 1)
	}
	st.t[v] = st.t[v * 2] + st.t[v * 2 + 1]
	return
}
 
func (st segmentTree) print(tl int, tr int, v int) {
	if tl == tr {
		printf("%v ", st.t[v])
		return
	}
	mid := (tl + tr) >> 1
	st.print(tl, mid, v * 2)
	st.print(mid + 1, tr, v * 2 + 1)
}