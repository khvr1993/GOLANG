package sentenceSimilarity

import (
	"fmt"
	"github.com/khvr1993/GOLANG/DS_Algorithms/PrintType"
	"strings"
)

/*
	arr1 will have the smaller sentence words
	arr2 will have the larger sentence words
*/
func checkPrefixAndSuffix(arr1 []string, arr2 []string) bool {
	PrintType.PrintString(&arr1)
	PrintType.PrintString(&arr2)

	itr := 0

	for itr < len(arr1) {
		//check for prefix
		if arr1[itr] == arr2[itr] {
			itr++
		} else {
			// As and when the match has not occurred break from the loop
			break
		}
	}
	fmt.Println("No of Prefixes matches ", itr)
	if itr >= len(arr1) {
		//We found that the complete sentence is a prefix of second sentence
		return true
	}

	itrRevArr2 := len(arr2) - 1
	itrRevArr1 := len(arr1) - 1
	revMatchCount := 0

	fmt.Println("itrRevArr1 => ", itrRevArr1)
	fmt.Println("itrRevArr2 => ", itrRevArr2)

	// Count the second iterator backwards and to the position where the prefix has ended
	for itrRevArr1 >= 0 && itrRevArr1 >= itr {
		fmt.Println("Array1 ", arr1[itrRevArr1])
		fmt.Println("Array2 ", arr2[itrRevArr2])
		if arr1[itrRevArr1] == arr2[itrRevArr2] {
			itrRevArr1--
			itrRevArr2--
			revMatchCount++
		} else {
			//As soon as match has not occurred break from the loop
			break
		}
	}
	fmt.Println("Reverse match Count ", revMatchCount)
	// Here we check if the second string is both a prefix as well as a suffix and then get the result.
	// We check the length because there might be some words in arr1 which may not be prefix as well as suffix and
	// we should think of those sentences as not similar
	if revMatchCount >= len(arr1) {
		return true
	} else if (revMatchCount + itr) == len(arr1) {
		return true
	} else {
		return false
	}

}

func SentencesSimilar(sentence1 string, sentence2 string) bool {
	/*
		Different cases we have in this problem are
		sentence1 is completely equal to sentence2
		sentence2/sentence1 is completely present in sentence1/sentence2 as either prefix or suffix
		both the sentences are completely distinct

		Approach can be
		check which is the shorter sentences
		split sentence one and sentence 2 and make 2 arrays.
		taking the smaller sentence
			1. check from the beginning of array till you get the match and store the index
			2. from the that index check for the suffix (may be in reverse order.
			3. If complete array is satisfied with true then we can return true else false

	*/
	var sen1Arr []string
	var sen2Arr []string
	var op bool

	if sentence1 == sentence2 {
		return true
	}

	sen1Arr = strings.Split(sentence1, " ")
	sen2Arr = strings.Split(sentence2, " ")

	if len(sen1Arr) > len(sen2Arr) {
		op = checkPrefixAndSuffix(sen2Arr, sen1Arr)
	} else {
		op = checkPrefixAndSuffix(sen1Arr, sen2Arr)
	}
	fmt.Println("Returning ", op)
	return op
}
