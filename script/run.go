package main

import (
	"flag"
	"weicai.zhao.io/gormx"
	"weicai.zhao.io/parse"
	"weicai.zhao.io/script/script"
)

var (
	scriptName = ""
)

func init() {
	flag.StringVar(&scriptName, "scriptName", "", "script")
	flag.Parse()
}

func main() {
	for _, err := range []error{
		script.Register(&script.Script{
			Name: "create po",
			Key:  "create_po",
			Handle: []script.Handle{func() error {
				parse.GenerateModule("./", "system", gormx.Config{
					Usage:       "default",
					RunMode:     "debug",
					DSN:         "root:root@tcp(127.0.0.1:3307)",
					Database:    "worker_app",
					MaxIdleConn: 10,
					MaxOpenConn: 10,
					MaxLifeTime: 10,
				})
				return nil
			}},
		}),
	} {
		if err != nil {
			panic(err)
		}
	}

	err := script.Run(scriptName)
	if err != nil {
		panic(err)
	}
}
