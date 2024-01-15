package buffer

import "strings"

type Buffer struct {
	strings.Builder
	v []interface{}
}

// NewBuffer creates a new Buffer.
func NewBuffer() *Buffer {
	return &Buffer{}
}

func (b *Buffer) WriteValue(v ...interface{}) error {
	b.v = append(b.v, v...)
	return nil
}

func (b *Buffer) Value() []interface{} {
	return b.v
}
