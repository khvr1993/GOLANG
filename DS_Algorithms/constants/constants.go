package constants

import "os"

const (
	//MERGESORTCUTOFF determines when to use insertion sort and when mergesort
	MERGESORTCUTOFF = 7
)

//HOME gets the HOME path of the system
var HOME = os.Getenv("HOME")

//JSONOPPATH gets the path where the output will be written
var JSONOPPATH = HOME + "/gocode/src/github.com/khvr1993/GOLANG/DS_Algorithms/JSON/output.txt"
