package lib

import (
	"database/sql"
	"encoding/binary"
	"encoding/json"
	"time"
)

type NullString struct {
	sql.NullString
}

type NullTime struct {
	sql.NullTime
}

type NullInt64 struct {
	sql.NullInt64
}

func (s NullInt64) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.Int64)
}

func (s *NullInt64) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		s.Int64, s.Valid = int64(0), false
		return nil
	}

	s.Int64, s.Valid = int64(binary.BigEndian.Uint64(data)), true
	return nil
}

func (s NullTime) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.Time)
}

func (s *NullTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		s.Time, s.Valid = time.Now(), false
		return nil
	}
	s.Time, s.Valid = time.Now(), true
	return nil
}

func (s NullString) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}

func (s *NullString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		s.String, s.Valid = "", false
		return nil
	}
	s.String, s.Valid = string(data), true
	return nil
}
