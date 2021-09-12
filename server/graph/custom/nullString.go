package custom

import (
	"database/sql"
	"errors"
	"io"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalNullString(t sql.NullString) graphql.Marshaler {
	if !t.Valid {
		return graphql.Null
	}

	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, t.String)
	})
}

func UnmarshalNullString(v interface{}) (sql.NullString, error) {
	if tmpStr, ok := v.(string); ok {
		return sql.NullString{String: tmpStr}, nil
	}
	return sql.NullString{}, errors.New("time should be RFC3339Nano formatted string")
}
