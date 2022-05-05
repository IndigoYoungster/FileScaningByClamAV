package filesender

import (
	"fmt"
	"io"
	"strings"
)

func (rm *responseModel) Read(p []byte) (n int, err error) {
	result := rm.String()
	n = copy(p, []byte(result))
	return n, io.EOF
}

type responseModel struct {
	Success bool `json:"success"`
	Data    data `json:"data"`
}

type data struct {
	Result []result `json:"result"`
}

type result struct {
	Name       string   `json:"name"`
	IsInfected bool     `json:"is_infected"`
	Viruses    []string `json:"viruses"`
}

func (rm *responseModel) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Success: %t\n", rm.Success))
	if rm.Success {
		sb.WriteString("Data:\n\tResult:\n")
		for _, val := range rm.Data.Result {
			sb.WriteString(fmt.Sprintf("\t\tName: %s\n\t\tIs infected: %t\n\t\tViruses: %v\n\n", val.Name, val.IsInfected, val.Viruses))
		}
	}
	return sb.String()
}
