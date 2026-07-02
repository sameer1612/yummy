package nullable

import "database/sql"

func ToFloat64(n sql.NullFloat64) *float64 {
	if n.Valid {
		return &n.Float64
	}
	return nil
}

func ToNullFloat64(f *float64) sql.NullFloat64 {
	if f == nil {
		return sql.NullFloat64{Valid: false}
	}
	return sql.NullFloat64{Float64: *f, Valid: true}
}
