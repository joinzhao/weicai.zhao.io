package create

import (
	"weicai.zhao.io/gormx"
	"weicai.zhao.io/script/internal"
	"weicai.zhao.io/script/internal/create"
)

const (
	CmdName               = create.TmplCmdName
	PoTmplCmdName         = CmdName + "_po"
	RepositoryTmplCmdName = CmdName + "_repository"
	DependencyTmplCmdName = CmdName + "_dependency"

	databaseName = "information_schema"
)

type Config struct {
	Mysql      gormx.Config
	Prefix     string
	Table      string
	SourceFile string
	TargetPath string
	NewPoPath  string
	OldPoPath  string
}

func NewDependencyCmd(cfg *Config) internal.Cmd {
	return &dependencyTmplCmd{cfg: cfg}
}

func NewPoCmd(cfg *Config) internal.Cmd {
	return &poTmplCmd{cfg: cfg}
}

func NewRepoCmd(cfg *Config) internal.Cmd {
	return &repoTmplCmd{cfg: cfg}
}
