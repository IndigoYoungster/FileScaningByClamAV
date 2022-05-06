package models

import (
	"fmt"
)

type Params struct {
	RemoteAddr string `json:"remote_addr"`
	App        string `json:"app"`
	Usr        string `json:"usr"`
	Comp       string `json:"comp"`
}

func (p *Params) String() string {
	return (fmt.Sprintf("Remote addr: %s\nApp: %s\nUsr: %s\nComp: %s\n", p.RemoteAddr, p.App, p.Usr, p.Comp))
}
