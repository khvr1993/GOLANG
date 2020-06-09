package merge

//Merge merges the elements in the array
func Merge(a *[]int, aux *[]int, lo int, hi int, mid int) {
	var i, k int
	j := mid + 1
	for i = lo; i <= hi; i++ {
		(*aux)[i] = (*a)[i]
	}
	i = lo
	for k = lo; k <= hi; k++ {
		if j <= hi && i <= mid {
			if (*aux)[i] > (*aux)[j] {
				(*a)[k] = (*aux)[j]
				j++
			} else if (*aux)[i] <= (*aux)[j] {
				(*a)[k] = (*aux)[i]
				i++
			}
		} else if j < hi {
			(*a)[k] = (*aux)[j]
			j++
		} else if i <= mid {
			(*a)[k] = (*aux)[i]
			i++
		}
	}
}
