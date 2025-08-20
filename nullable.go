package directus

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
)

// Nullable represents a value that can be null
type Nullable[T any] struct {
	Data  T
	Valid bool
	Set   bool
}

// NewNullable creates a new Nullable with the given value
func NewNullable[T any](value T) Nullable[T] {
	return Nullable[T]{
		Data:  value,
		Valid: true,
		Set:   true,
	}
}

// Null creates a new Nullable with null value
func Null[T any]() Nullable[T] {
	return Nullable[T]{
		Valid: false,
		Set:   true,
	}
}

// IsNull returns true if the value is null
func (n Nullable[T]) IsNull() bool {
	return !n.Valid
}

// IsZero returns true if the value is zero
func (n Nullable[T]) IsZero() bool {
	return !n.Valid || reflect.ValueOf(n.Data).IsZero()
}

// MarshalJSON implements json.Marshaler
func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.Data)
}

// UnmarshalJSON implements json.Unmarshaler
func (n *Nullable[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		n.Valid = false
		n.Set = true
		return nil
	}

	var value T
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	n.Data = value
	n.Valid = true
	n.Set = true
	return nil
}

// Scan implements sql.Scanner
func (n *Nullable[T]) Scan(value interface{}) error {
	if value == nil {
		n.Data = *new(T)
		n.Valid = false
		n.Set = true
		return nil
	}

	switch v := value.(type) {
	case sql.NullString:
		if v.Valid {
			var target T
			if err := json.Unmarshal([]byte(v.String), &target); err != nil {
				return err
			}
			n.Data = target
			n.Valid = true
			n.Set = true
		} else {
			n.Data = *new(T)
			n.Valid = false
			n.Set = true
		}
	case sql.NullInt64:
		if v.Valid {
			switch any(*new(T)).(type) {
			case int, int8, int16, int32, int64:
				n.Data = any(int(v.Int64)).(T)
				n.Valid = true
				n.Set = true
			case uint, uint8, uint16, uint32, uint64:
				n.Data = any(uint(v.Int64)).(T)
				n.Valid = true
				n.Set = true
			default:
				return fmt.Errorf("unsupported type for NullInt64")
			}
		} else {
			n.Data = *new(T)
			n.Valid = false
			n.Set = true
		}
	case sql.NullFloat64:
		if v.Valid {
			switch any(*new(T)).(type) {
			case float32, float64:
				n.Data = any(v.Float64).(T)
				n.Valid = true
				n.Set = true
			default:
				return fmt.Errorf("unsupported type for NullFloat64")
			}
		} else {
			n.Data = *new(T)
			n.Valid = false
			n.Set = true
		}
	case sql.NullBool:
		if v.Valid {
			switch any(*new(T)).(type) {
			case bool:
				n.Data = any(v.Bool).(T)
				n.Valid = true
				n.Set = true
			default:
				return fmt.Errorf("unsupported type for NullBool")
			}
		} else {
			n.Data = *new(T)
			n.Valid = false
			n.Set = true
		}
	default:
		if val, ok := value.(T); ok {
			n.Data = val
			n.Valid = true
			n.Set = true
		} else {
			return fmt.Errorf("unsupported type: %T", value)
		}
	}
	return nil
}

// Value implements driver.Valuer
func (n Nullable[T]) Value() (interface{}, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Data, nil
}
