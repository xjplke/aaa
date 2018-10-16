package main

import (
	"github.com/xjplke/aaa/portal/cmcc"
	version "github.com/xjplke/aaa/portal/cmcc/v1"
)

func main() {
	opts := &cmcc.CMCCOpts{
		Address: "0.0.0.0",
		Port:    "2000",
	}
	ver := version.Version{}
	portal := cmcc.NewCMCCPortal(opts, nil, nil, nil, ver)
	portal.Start()
}
