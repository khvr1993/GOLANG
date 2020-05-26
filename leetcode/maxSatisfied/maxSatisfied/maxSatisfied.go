package maxSatisfied

import "log"

//MaxSatisfied returns the Maximum customers the shopkeeper can keep
func MaxSatisfied(customers []int, grumpy []int, X int) int {
	length := len(customers)
	var start, end int
	var custCount int
	var maxCust int
	log.Println("The number of Customers => ", length)
	satisfiedCust := sumSatisfiedCust(&customers, &grumpy, &length)
	for end < length {
		log.Println("Current value of endWindow ", end, " startWindow ", start, " value of X ", X)
		if X > 0 {
			X--
			if grumpy[end] == 1 {
				custCount += customers[end]
			}

			maxCust = max(maxCust, custCount)
			log.Println("The max customers at this point in time => ", maxCust, " custCount is => ", custCount)
			end++
		} else {
			X++
			log.Println("deleting the customer count ", customers[start], " at ", start)
			if grumpy[start] == 1 {
				custCount -= customers[start]
			}
			start++
		}
	}

	return maxCust + satisfiedCust

}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func sumSatisfiedCust(cust *[]int, grump *[]int, len *int) int {
	var i, sum int
	for i < *len {
		if (*grump)[i] == 0 {
			sum += (*cust)[i]
		}
		i++
	}
	return sum
}
