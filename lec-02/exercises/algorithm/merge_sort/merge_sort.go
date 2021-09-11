package main

func MergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	m := int(len(a) / 2)

	l := make([]int, m)
	r := make([]int, len(a)-m)

	for i := 0; i < len(a); i++ {
		if i < m {
			l[i] = a[i]
		} else {
			r[i-m] = a[i]
		}
	}

	return mergeUnsortedArrs(MergeSort(l), MergeSort(r))
}

func mergeUnsortedArrs(l, r []int) (rel []int) {
	rel = make([]int, len(l)+len(r))

	i := 0
	for len(l) > 0 && len(r) > 0 {
		if l[0] < r[0] {
			rel[i] = l[0]
			l = l[1:]
		} else {
			rel[i] = r[0]
			r = r[1:]
		}
		i++
	}

	for j := 0; j < len(l); j++ {
		rel[i] = l[j]
		i++
	}

	for j := 0; j < len(r); j++ {
		rel[i] = r[j]
		i++
	}
	return
}
