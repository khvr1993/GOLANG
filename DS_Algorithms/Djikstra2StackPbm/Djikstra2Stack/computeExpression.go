package Djikstra2Stack

import (
	"errors"
	"strconv"

	stack "github.com/khvr1993/GOLANG/DS_Algorithms/linkedStackOfElements/stack"
)

//Compute calculates the expression passed. Assumption is correct expression will be passed
func Compute(a string) (int, error) {
	var i, result int
	var operatorStack = new(stack.Stack)
	var numberStack = new(stack.Stack)
	var ignoreOp = map[string]bool{
		"(": false,
		")": true,
	}
	// var closedPar = map[string]bool{
	// 	")": true,
	// }
	var operators = map[string]bool{
		"*": true,
		"/": true,
		"+": true,
		"-": true,
	}

	for i < len(a) {
		val, _ := ignoreOp[string(a[i])]
		if val {
			//write POP
			val2, _ := strconv.Atoi(numberStack.Pop().(string))
			val1, _ := strconv.Atoi(numberStack.Pop().(string))
			operation := operatorStack.Pop()
			if operation.(string) == "*" {
				result = val1 * val2
			}
			if operation.(string) == "-" {
				result = val1 - val2
			}
			if operation.(string) == "+" {
				result = val1 + val2
			}
			if operation.(string) == "/" {
				result = val1 / val2
			}
			pushVal := strconv.Itoa(result)
			numberStack.Push(pushVal)
		}
		op, _ := operators[string(a[i])]
		if op {
			//write push to operator stack
			operatorStack.Push(string(a[i]))
		}

		if _, err := strconv.Atoi(string(a[i])); err == nil {
			//push into number stack
			numberStack.Push(string(a[i]))
		}
		i++
	}
	//operatorStack.ShowStack()
	numberStack.ShowStack()
	if numberStack.Size() > 1 || operatorStack.Size() > 0 {
		return 0, errors.New("The expression is of wrong type")
	}
	returnVal, _ := strconv.Atoi(numberStack.Pop().(string))
	return returnVal, nil
}
