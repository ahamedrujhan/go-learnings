package ioManager

type IOManger interface {
	ReadLines() ([]string, error)
	WriteData(interface{}) error
}
