package create

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/token"
	"os"
	"weicai.zhao.io/script/internal"
	"weicai.zhao.io/tools"
)

func NewTmplCmd(v ast.Visitor, cfg TmplConfig) internal.Cmd {
	_, f, _ := tools.ParseAstFile(cfg.TmplFile)

	return &createTmplCmd{
		tmplFile:   f,
		v:          v,
		targetFile: cfg.TargetFile,
	}
}

type TmplConfig struct {
	TmplFile   string
	TargetFile string
}

const (
	TmplCmdName = CmdName + "_tmpl"
)

type createTmplCmd struct {
	tmplFile   *ast.File
	v          ast.Visitor
	targetFile string
}

func (cmd createTmplCmd) GetName() string {
	return TmplCmdName
}

func (cmd createTmplCmd) Do() error {
	// walk
	ast.Walk(cmd.v, cmd.tmplFile)

	// ast.file to byte
	var body []byte
	buf := bytes.NewBuffer(body)
	err := format.Node(buf, token.NewFileSet(), cmd.tmplFile)
	if err != nil {
		return err
	}

	err = createFile(cmd.targetFile)
	if err != nil {
		return err
	}

	// write to file
	err = os.WriteFile(cmd.targetFile, buf.Bytes(), 0777)
	if err != nil {
		return err
	}
	return nil
}
