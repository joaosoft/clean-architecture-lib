package table

import (
	"fmt"

	dbr "github.com/gocraft/dbr/v2"
)

type Table struct {
	Schema string
	Table  string
}

func New(schema, table string) *Table {
	return &Table{
		Schema: schema,
		Table:  table,
	}
}

func (t *Table) As(alias string) dbr.Builder {
	return dbr.BuildFunc(func(d dbr.Dialect, buf dbr.Buffer) error {
		_ = t.Build(d, buf)
		_, _ = buf.WriteString(" AS ")
		_, _ = buf.WriteString(d.QuoteIdent(alias))
		return nil
	})
}

func (t *Table) Build(d dbr.Dialect, buf dbr.Buffer) error {
	_, err := buf.WriteString(d.QuoteIdent(t.String()))
	return err
}

// String Will assume the dialect as PostgreSQL
func (t *Table) String() string {
	return fmt.Sprintf("%s.%s", t.Schema, t.Table)
}
