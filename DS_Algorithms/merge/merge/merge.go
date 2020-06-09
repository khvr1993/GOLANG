package merge

//Merge merges the elements in the array
func Merge(a *[]int, lo int, hi int, mid int) {
	var Copy []int
	var i, k int
	j := mid + 1
	for i <= hi {
		Copy = append(Copy, (*a)[i])
		i++
	}
	i = lo
	for k = lo; k <= hi; k++ {
		if j <= hi && i <= mid {
			if Copy[i] > Copy[j] {
				(*a)[k] = Copy[j]
				j++
			} else if Copy[i] <= Copy[j] {
				(*a)[k] = Copy[i]
				i++
			}
		} else if j < hi {
			(*a)[k] = Copy[j]
			j++
		} else if i <= mid {
			(*a)[k] = Copy[i]
			i++
		}
	}
}
