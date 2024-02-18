package create

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"weicai.zhao.io/gormx"
	"weicai.zhao.io/script/internal/create"
	"weicai.zhao.io/script/internal/source"
	"weicai.zhao.io/script/internal/visitor"
	"weicai.zhao.io/tools"
)

type poTmplCmd struct {
	cfg *Config
}

func (cmd *poTmplCmd) GetName() string {
	return PoTmplCmdName
}

func (cmd *poTmplCmd) Do() error {
	var cfg = cmd.cfg.Mysql
	var dbName = cfg.Database
	cfg.Database = databaseName

	var manager = gormx.New([]*gormx.Config{&cfg})
	db, err := manager.Use(context.TODO(), cfg.Usage)
	if err != nil {
		return err
	}

	var errS error
	for _, table := range cmd.cfg.Table {
		columns, err := source.New(db).Find(dbName, table)

		if err != nil {
			errS = errors.Join(errS, err)
			continue
		}

		// transfer
		fields := visitor.ColumnToField(columns)

		// rename
		structName := tools.UnderlineToUpperCamelCase(strings.TrimPrefix(table, cmd.cfg.Prefix))

		// visitor
		v := visitor.NewPoVisitor(fields, structName, table)

		var targetFile = filepath.Join(cmd.cfg.TargetPath, fmt.Sprintf("%s.go", strings.TrimPrefix(table, cmd.cfg.Prefix)))

		c := create.NewTmplCmd(v, create.TmplConfig{
			TmplFile:   cmd.cfg.SourceFile,
			TargetFile: targetFile,
		})

		err = c.Do()
		if err != nil {
			errS = errors.Join(errS, err)
		}
	}

	return errS
}
