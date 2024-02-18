package main

import (
	"flag"
	"strings"
	"weicai.zhao.io/script/cmd/create"
	"weicai.zhao.io/script/internal"
	"weicai.zhao.io/tools"
)

// parse flag
var (
	cmdName    string
	configFile string

	cfg Config
)

const (
	defaultConfigFile = "./etc/script.yaml"
	defaultPoPath     = "weicai.zhao.io/script/tmpl/po"
)

func init() {
	flag.StringVar(&cmdName, "cmd", "", "cmd")
	flag.StringVar(&configFile, "config", defaultConfigFile, "config")
	flag.Parse()

	err := tools.ReadFromYaml(configFile, &cfg)
	if err != nil {
		panic(err)
	}

	// bind cmd
	cfg.register()

}

func main() {
	switch strings.ToLower(cmdName) {
	case "po":
		cmdName = create.PoTmplCmdName
	case "repository":
		cmdName = create.RepositoryTmplCmdName
	case "dependency":
		cmdName = create.DependencyTmplCmdName
	case "all":
		cmdName = create.CmdName
	default:
		panic("don't has this cmd")
	}
	err := internal.Do(cmdName)
	if err != nil {
		panic(err)
	}
}

type Config struct {
	Po         *create.Config
	Repository *create.Config
	Dependency *create.Config
}

func (c *Config) register() {
	if c.Po != nil {
		var cmd = create.NewPoCmd(c.Po)
		internal.Bind(create.PoTmplCmdName, cmd)
		internal.Bind(create.CmdName, cmd)
	}

	if c.Repository != nil {
		if c.Repository.OldPoPath != "" {
			c.Repository.OldPoPath = defaultPoPath
		}
		var cmd = create.NewRepoCmd(c.Repository)
		internal.Bind(create.RepositoryTmplCmdName, cmd)
		internal.Bind(create.CmdName, cmd)
	}

	if c.Dependency != nil {
		if c.Repository.OldPoPath != "" {
			c.Repository.OldPoPath = defaultPoPath
		}
		var cmd = create.NewDependencyCmd(c.Dependency)
		internal.Bind(create.DependencyTmplCmdName, cmd)
		internal.Bind(create.CmdName, cmd)
	}
}
