package internal

import (
	"fmt"
	"sync"
)

var (
	_container *container
)

func init() {
	_container = &container{
		cmdMap: map[string][]Cmd{},
		lock:   sync.RWMutex{},
	}
}

func Bind(name string, cmd Cmd) {
	_container.Bind(name, cmd)
}

func Do(name string) error {
	return _container.Do(name)
}

type container struct {
	cmdMap map[string][]Cmd
	lock   sync.RWMutex
}

func (c *container) Bind(name string, cmd Cmd) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if v, ok := c.cmdMap[name]; ok {
		v = append(v, cmd)
		c.cmdMap[name] = v
	} else {
		c.cmdMap[name] = []Cmd{cmd}
	}
}

func (c *container) Do(name string) error {
	c.lock.RLock()
	defer c.lock.RUnlock()
	if v, ok := c.cmdMap[name]; ok {
		if v == nil || len(v) == 0 {
			return fmt.Errorf("empty cmd entity")
		}
		for _, cmd := range v {
			err := cmd.Do()
			if err != nil {
				return err
			}
		}
		return nil
	} else {
		return fmt.Errorf("don't have this name by cmd container")
	}
}
