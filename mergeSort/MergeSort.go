func MergeSort(s []int) []int {
	if len(s) <= 1 {
		return s
	}

	n := len(s) / 2

	var l []int
	var r []int

	l = MergeSort(s[:n])
	r = MergeSort(s[n:])

	return merge(l, r)
}
