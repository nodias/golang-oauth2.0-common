package model

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type TomlConfig struct {
	Service   string
	Logpaths  logpaths
	Databases map[string]databases
	Servers   map[string]servers
}

type servers struct {
	IP   string
	PORT string
}

type logpaths struct {
	Logpath string
}

type databases struct {
	Server string
	Port   string
	Enable bool
}

func (t *TomlConfig) Load(cp string) {
	cmdargs := GetCmdargs()
	fpath := fmt.Sprintf(cp, cmdargs.Phase)
	if _, err := toml.DecodeFile(fpath, &t); err != nil {
		fmt.Println(err)
	}
}

func (t *TomlConfig) ApmServerUrl() string {
	return fmt.Sprintf("%s%s", t.Servers["APM"].IP, t.Servers["APM"].PORT)
}
