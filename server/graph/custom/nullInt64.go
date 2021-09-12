package custom

import (
	"database/sql"
	"errors"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalNullInt64(t sql.NullInt64) graphql.Marshaler {
	if !t.Valid {
		return graphql.Null
	}

	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatInt(t.Int64, 10))
	})
}

func UnmarshalNullInt64(v interface{}) (sql.NullInt64, error) {
	if tmpStr, ok := v.(int64); ok {
		return sql.NullInt64{Int64: tmpStr}, nil
	}
	return sql.NullInt64{}, errors.New("time should be RFC3339Nano formatted string")
}
