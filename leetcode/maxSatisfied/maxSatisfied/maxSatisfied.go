package maxSatisfied

import "log"

//MaxSatisfied returns the Maximum customers the shopkeeper can keep
func MaxSatisfied(customers []int, grumpy []int, X int) int {
	length := len(customers)
	var start, end int
	var custCount int
	var maxCust int
	log.Println("The number of Customers => ", length)
	for end < length {
		log.Println("Current value of endWindow ", end, " startWindow ", start, " value of X ", X)
		if X > 0 {
			if grumpy[end] != 0 {
				X--
			}
			custCount += customers[end]
			maxCust = max(maxCust, custCount)
			log.Println("The max customers at this point in time => ", maxCust, " custCount is => ", custCount)
			end++
		} else {
			if grumpy[start] != 0 {
				X++
			}
			log.Println("deleting the customer count ", customers[start], " at ", start)
			custCount -= customers[start]
			start++
		}
	}
	return maxCust

}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
