package JSON

import (
	"encoding/json"
	"log"
)

//Employee implements the  type employee struct
type Employee struct {
	Name        string
	JoiningDate string `json:"joining_date"`
	EmployeeID  int    `json:"employee_id"`
	IsActive    bool   `json:"is_active"`
	Role        string
	TeamMem     []string `json:"team_members"`
}

//Marshal converts the GO struct to JSON Type
func Marshal(E *[]Employee) []byte {
	data, err := json.MarshalIndent(E, "", " ") //You can use JSON Marshal to get JSON output without structure
	if err != nil {
		log.Fatalf("JSON marshalling failed %s", err)
	}
	return data
}
