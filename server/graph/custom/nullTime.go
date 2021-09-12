package custom

import (
	"database/sql"
	"errors"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

var NullTime sql.NullTime

func MarshalNullTime(t sql.NullTime) graphql.Marshaler {
	if t.Time.IsZero() {
		return graphql.Null
	}

	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(t.Time.Format(time.RFC3339Nano)))
	})
}

func UnmarshalNullTime(v interface{}) (sql.NullTime, error) {
	if tmpStr, ok := v.(string); ok {
		data, err := time.Parse(time.RFC3339Nano, tmpStr)
		if err != nil {
			return sql.NullTime{}, err
		}
		return sql.NullTime{Time: data}, nil
	}
	return sql.NullTime{}, errors.New("time should be RFC3339Nano formatted string")
}
