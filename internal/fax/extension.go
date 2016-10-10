package fax

import (
	"errors"
	"fmt"
)

var uncompressedMode = errors.New("fax: uncompressed mode not supported")

func extension(d *decoder) error {
	extension := d.head >> 22
	extension &= 0x7
	handler := extensionTable[extension]
	if handler == nil {
		return fmt.Errorf("fax: unknown extension 0b%03b", extension)
	}
	if e := d.pop(10); e != nil {
		return e
	}
	return handler(d)
}

var extensionTable = [8]func(d *decoder) error{
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	func(d *decoder) error {
		return uncompressedMode
	},
}
