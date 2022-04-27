package main

import (
	"fmt"
	"strings"
)

type responseModel struct {
	Success bool
	Data    data
}

type data struct {
	Result []result
}

type result struct {
	Name       string
	IsInfected bool
	Viruses    []string
}

func (rm *responseModel) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Success: %t\n", rm.Success))
	if rm.Success {
		sb.WriteString("Data:\n\tResult:\n")
		for _, val := range rm.Data.Result {
			sb.WriteString(fmt.Sprintf("\t\tName: %s\n\t\tIs infected: %t\n\t\tViruses: %v\n", val.Name, val.IsInfected, val.Viruses))
		}
	}
	return sb.String()
}
