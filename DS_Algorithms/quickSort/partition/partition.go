package partition

import (
	"log"

	"github.com/khvr1993/GOLANG/DS_Algorithms/PrintType"
)

//Partition places all the elements which are less than key to the left side and
// all the elements which are greater to the right side
func Partition(a *[]int, lo int, high int) int {
	log.Println("Partition")
	i := lo + 1
	j := high
	//(*a)[lo] is the key
	for {
		log.Println("the key value is ", (*a)[lo], "i is ", i, " j is ", j, " lo is ", lo, "a[i]", (*a)[i], "high", high)
		//How many elements are less than the key
		for (*a)[i] <= (*a)[lo] {
			i++
			if i == j || i > high {
				break
			}
		}
		// How many elements are greater than the key
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
	log.Println("i => ", i, "j => ", j)
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
