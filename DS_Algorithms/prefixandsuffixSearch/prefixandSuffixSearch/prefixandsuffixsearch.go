package prefixandSuffixSearch

import "strings"

type WordFilter struct {
	word []string
}

func Constructor(words []string) WordFilter {
	wordFilter := new(WordFilter)
	// for _,val := range words{
	// fmt.Println(val)
	wordFilter.word = words
	// }
	return *wordFilter
}

func (this *WordFilter) F(prefix string, suffix string) int {
	maxIndex := -1
	searchVal := suffix + string('#') + prefix
	index := len(this.word) - 1
	for index >= 0 {
		if strings.Contains(this.word[index]+string('#')+this.word[index], searchVal) {
			maxIndex = index
			return maxIndex
		}
		index -= 1
	}
	return maxIndex
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(prefix,suffix);
 */
