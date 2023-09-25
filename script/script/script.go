package script

import (
	"fmt"
	"log"
)

type Handle func() error

var _s = make(map[string]*Script)

func Register(script *Script) error {
	if script == nil {
		panic(fmt.Errorf("runtime error: invalid memory address or nil pointer dereference"))
	}
	_, ok := _s[script.Key]
	if ok {
		return fmt.Errorf("key [%s] is exist", script.Key)
	}
	_s[script.Key] = script
	return nil
}

func Run(key string) error {
	script, ok := _s[key]
	if !ok {
		return fmt.Errorf("key [%s] is not exist", key)
	}

	return do(script)
}

type Script struct {
	Name   string `json:"name"`
	Key    string `json:"key"`
	Handle []Handle
}

func do(script *Script) error {
	log.Println("do script ->")
	log.Println("script name: ", script.Name)
	log.Println("script key: ", script.Key)
	if script.Handle == nil || len(script.Handle) == 0 {
		return nil
	}
	for i := 0; i < len(script.Handle); i++ {
		if err := script.Handle[i](); err != nil {
			return err
		}
	}
	return nil
}
