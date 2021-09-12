package custom

import (
	"database/sql"
	"errors"
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalNullFloat64(t sql.NullFloat64) graphql.Marshaler {
	if !t.Valid {
		return graphql.Null
	}

	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, fmt.Sprintf("%g", t.Float64))
	})
}

func UnmarshalNullFloat64(v interface{}) (sql.NullFloat64, error) {
	if tmpStr, ok := v.(float64); ok {
		return sql.NullFloat64{Float64: tmpStr}, nil
	}
	return sql.NullFloat64{}, errors.New("time should be RFC3339Nano formatted string")
}
