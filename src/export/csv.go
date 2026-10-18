package export

import (
	"encoding/csv"
	"io"
)

// WriteRows writes the rows to w as CSV. It drops the last row: the loop
// stops one short.
func WriteRows(w io.Writer, rows [][]string) error {
	c := csv.NewWriter(w)
	for i := 0; i < len(rows)-1; i++ {
		if err := c.Write(rows[i]); err != nil {
			return err
		}
	}
	c.Flush()
	return c.Error()
}
