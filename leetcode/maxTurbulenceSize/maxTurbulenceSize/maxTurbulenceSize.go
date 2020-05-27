package maxTurbulenceSize

import "log"

//MaxTurbulenceSize returns the longest Turbulent Subarray
//the subarray is turbulent if the comparison sign flips between each adjacent pair of elements in the subarray.
func MaxTurbulenceSize(A []int) int {
	var size, start, end, currSize int
	var pastSign string
	log.Println("pastSign ", pastSign)
	if len(A) == 1 {
		return 1
	}

	for end < len(A)-1 {
		log.Println("currVal of end and start ", end, start, " the value at this position is ", A[end], " The next val is ", A[end+1], " and the pastSign is ", pastSign)
		if A[end] > A[end+1] && (pastSign == "S" || pastSign == "") {
			log.Println("value is Greater")
			pastSign = "G"
			end++
			currSize = end - start + 1
			log.Println("currSize", currSize)
			size = max(size, currSize)
		} else if A[end] < A[end+1] && (pastSign == "G" || pastSign == "") {
			log.Println("value is Smaller")
			pastSign = "S"
			end++
			currSize = end - start + 1
			log.Println("currSize", currSize)
			size = max(size, currSize)
		} else {
			start = end
			if A[end] > A[end+1] {
				pastSign = "G"
			} else if A[end] < A[end+1] {
				pastSign = "S"
			} else {
				pastSign = ""
				start++
			}
			end++
			currSize = end - start + 1
			size = max(size, currSize)
			log.Println("restart the turbulence the size is and current size ", size, currSize)
		}
	}
	return size
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
