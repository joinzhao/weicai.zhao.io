package consts

type Option uint8

// DBOption database option number
type DBOption Option

const (
	_              = iota
	DBCreateOption // DBCreateOption create data to database
	DBUpdateOption // DBUpdateOption update data from database
	DBDeleteOption // DBDeleteOption delete data from database
	DBSelectOption // DBSelectOption select data from database
)
