package countInvMerge

//CountInvMerge counts the inversion operations happening in merge
func CountInvMerge(a *[]int, aux *[]int, lo int, hi int, mid int) int {
	var i, j, k, invCount int
	//Copy the main Array to the temp Array
	for i = lo; i <= hi; i++ {
		(*aux)[i] = (*a)[i]
	}
	i = lo
	j = mid + 1

	for k = lo; k <= hi; k++ {
		if j <= hi && i <= mid {
			if (*aux)[i] > (*aux)[j] {
				(*a)[k] = (*aux)[j]
				j++
				invCount++ //Here inversion has happening => we moved the element position
			} else if (*aux)[i] <= (*aux)[j] {
				(*a)[k] = (*aux)[i]
				i++
			}
		} else if j < hi {
			//meaning all the elements indices > i have been exhausted.
			//There are some elements remaining for indices j
			(*a)[k] = (*aux)[j]
		} else if i <= mid {
			//meaning all the elements indices > j have been exhausted.
			//There are some elements remaining for indices i
			(*a)[k] = (*aux)[i]
			i++
		}
	}
	return invCount
}
