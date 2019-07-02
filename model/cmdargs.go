package model

import (
	"flag"
	"sync"
)

type cmdargs struct {
	Phase string
}

//flag
var p = flag.String("phase", "local", "input phase e.g)local, dv")
var ca = cmdargs{
	Phase: *p,
}

//singleton
var instance *cmdargs
var once sync.Once

func GetCmdargs() *cmdargs {
	once.Do(func() {
		instance = &ca
	})
	return instance
}
