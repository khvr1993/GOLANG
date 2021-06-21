package decodeString

import (
	stack2 "github.com/khvr1993/GOLANG/DS_Algorithms/Stack/Stack"
	"log"
)

var stack = stack2.NewStack()
var decodedString string

func DecodeMessage(s string) string {
	log.Println("Begin Decode Message")

	//Assuming that the first value in the string is always an integer and the value before [ is always an int Type
	var decodedString string
	var decodedStringArr []string
	stack.Push(string(s[0]))

	for i := 1; i < len(s); i++ {
		stack.Push(string(s[i]))
	}
	//log.Println("No of elements Pushed => ", stack.Size())

	for !stack.IsEmpty() {
		val := stack.Pop()
		if val == "]" {
			log.Println(decodedString)
			/*
				since we encountered ] in the middle of the decoded string
				we will add this to an array so that we keep track of it and concantenate it to
				the larger part of the code
				This will store the first <count>[<substr>] in the array
			*/
			decodedStringArr = append(decodedStringArr, decodedString)
			decodedString = ""
		} else if val == "[" {
			repeatTimes := stack.Pop()
			log.Println("No of Repeat times ", repeatTimes)
			decodedString = repeatString(decodedString, repeatTimes.(string))
			/*
				After repeating the current data we have to check if there is any older data present
				(old repetitive data 2[a2[b]1[z]2[h]] ) and
				add it to the current one and delete from the stored data
			*/
			for _, d := range decodedStringArr {
				decodedString = decodedString + d
			}
			decodedStringArr = decodedStringArr[:len(decodedStringArr)-1]
		} else {
			decodedString = val.(string) + decodedString
			log.Println("decodedString => ", decodedString)
		}

	}
	return decodedString
}
