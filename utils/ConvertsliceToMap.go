package utils

//ConvertsliceToMap Converts the provided slice to map
func ConvertsliceToMap(a []int) map[int]int {
	newMap := make(map[int]int)
	var i int
	for i < len(a) {
		newMap[a[i]] = i
		i++
	}
	return newMap
}
