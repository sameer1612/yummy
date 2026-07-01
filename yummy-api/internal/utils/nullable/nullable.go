package nullable

import "database/sql"

func NullableFloat(n sql.NullFloat64) *float64 {
	if n.Valid {
		return &n.Float64
	}
	return nil
}
