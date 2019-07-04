package model

import (
	"flag"
	"fmt"
	"sync"
)

var ca cmdargs

type cmdargs struct {
	Phase string
}

func (ca cmdargs) String() string {
	return fmt.Sprintf("phase : %s", ca.Phase)
}

func init() {
	p := flag.String("phase", "local", "input phase e.g)local, dv")
	flag.Parse()
	ca = cmdargs{
		Phase: *p,
	}
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
