package main

import (
	"bufio"
	"fmt"
	"os"

	JSON "github.com/khvr1993/GOLANG/DS_Algorithms/JSON/Json"
	"github.com/khvr1993/GOLANG/DS_Algorithms/constants"
)

func main() {
	e1 := JSON.Employee{
		Name:        "Harsha",
		JoiningDate: "14-JAN-209",
		EmployeeID:  1234,
		IsActive:    true,
		Role:        "Systems Engineer II",
		TeamMem:     []string{"Madhava", "Dhanu", "Kiran"},
	}

	e2 := JSON.Employee{
		Name:        "madhava",
		JoiningDate: "14-JAN-209",
		EmployeeID:  1234,
		IsActive:    true,
		Role:        "Systems Engineer II",
		TeamMem:     []string{"Harsha", "Dhanu", "Kiran"},
	}

	e3 := JSON.Employee{
		Name:        "ABC",
		JoiningDate: "14-JAN-209",
		EmployeeID:  1234,
		IsActive:    true,
		Role:        "Systems Engineer II",
		TeamMem:     []string{""},
	}

	emps := []JSON.Employee{e1, e2, e3}

	data := JSON.Marshal(&emps)

	/*   Writing output to file   */
	stdout, err := os.Create(constants.JSONOPPATH)
	checkError(err)
	defer stdout.Close()
	writer := bufio.NewWriterSize(stdout, 1024*1024)
	fmt.Fprintf(writer, "%s\n", data)
	writer.Flush()
	/* ************************ */
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
