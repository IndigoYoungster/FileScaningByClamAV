package main

import (
	"fmt"
)

type params struct {
	RemoteAddr string `json:"remote_addr"`
	App        string `json:"app"`
	Usr        string `json:"user"`
	Comp       string `json:"comp"`
}

func (p *params) String() string {
	return (fmt.Sprintf("Remote addr: %s\nApp: %s\nUser: %s\nComp: %s\n", p.RemoteAddr, p.App, p.Usr, p.Comp))
}
