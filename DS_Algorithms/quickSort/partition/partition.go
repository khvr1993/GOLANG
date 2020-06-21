package partition

import (
	"github.com/khvr1993/GOLANG/DS_Algorithms/PrintType"
)

//Partition places all the eleements which are less than key to the ledt side and
// all the elements which are greater to the right side
func Partition(a *[]int, lo int, high int) int {
	i := lo + 1
	j := high
	//(*a)[lo] is the key
	for {
		//log.Println("the key value is ", (*a)[lo], "i is ", i, " j is ", j, " lo is ", lo, "a[i]", (*a)[i], "high", high)
		for (*a)[i] <= (*a)[lo] {
			i++
			if i == j || i > high {
				break
			}
		}

		for (*a)[lo] < (*a)[j] {
			j--
			if i == j || j < lo {
				break
			}
		}
		if j < i {
			break
		}

		exchange(a, i, j)
		PrintType.PrintArray(a)
	}
	//log.Println("i => ", i, "j => ", j)
	exchange(a, lo, j)
	//log.Println("j ", (*a)[j], " lo ", (*a)[lo])
	PrintType.PrintArray(a)

	return j
}

func exchange(a *[]int, i int, j int) {
	tempVal := (*a)[i]
	(*a)[i] = (*a)[j]
	(*a)[j] = tempVal
}
