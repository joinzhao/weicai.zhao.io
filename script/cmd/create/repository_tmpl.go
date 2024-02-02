package create

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"weicai.zhao.io/gormx"
	"weicai.zhao.io/script/internal/create"
	"weicai.zhao.io/script/internal/source"
	"weicai.zhao.io/script/internal/visitor"
	"weicai.zhao.io/tools"
)

type repoTmplCmd struct {
	cfg *Config
}

func (cmd *repoTmplCmd) GetName() string {
	return RepositoryTmplCmdName
}

func (cmd *repoTmplCmd) Do() error {
	var cfg = cmd.cfg.Mysql
	var dbName = cfg.Database
	cfg.Database = databaseName

	var manager = gormx.New([]*gormx.Config{&cfg})
	db, err := manager.Use(context.TODO(), cfg.Usage)
	if err != nil {
		return err
	}

	columns, err := source.NewUnq(db).Find(dbName, cmd.cfg.Table)
	if err != nil {
		return err
	}
	if len(columns) == 0 {
		return nil
	}

	// rename
	structName := tools.UnderlineToUpperCamelCase(strings.TrimPrefix(cmd.cfg.Table, cmd.cfg.Prefix))

	// visitor
	v := visitor.NewRepoVisitor(structName, cmd.cfg.NewPoPath, cmd.cfg.OldPoPath, columns[0])

	var targetFile = filepath.Join(cmd.cfg.TargetPath, fmt.Sprintf("%s.go", strings.TrimPrefix(cmd.cfg.Table, cmd.cfg.Prefix)))

	c := create.NewTmplCmd(v, create.TmplConfig{
		TmplFile:   cmd.cfg.SourceFile,
		TargetFile: targetFile,
	})

	return c.Do()
}
