package mysql

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql/driver"
	"fmt"
)

// BookType is the 'book_type' enum type from schema 'booktest'.
type BookType uint16

const (
	// BookTypeFiction is the 'FICTION' book_type.
	BookTypeFiction BookType = 1
	// BookTypeNonfiction is the 'NONFICTION' book_type.
	BookTypeNonfiction BookType = 2
)

// String satisfies the fmt.Stringer interface.
func (bt BookType) String() string {
	switch bt {
	case BookTypeFiction:
		return "FICTION"
	case BookTypeNonfiction:
		return "NONFICTION"
	}
	return fmt.Sprintf("BookType(%d)", bt)
}

// MarshalText marshals BookType into text.
func (bt BookType) MarshalText() ([]byte, error) {
	return []byte(bt.String()), nil
}

// UnmarshalText unmarshals BookType from text.
func (bt *BookType) UnmarshalText(buf []byte) error {
	switch s := string(buf); s {
	case "FICTION":
		*bt = BookTypeFiction
	case "NONFICTION":
		*bt = BookTypeNonfiction
	default:
		return ErrInvalidBookType(s)
	}
	return nil
}

// Value satisfies the driver.Valuer interface.
func (bt BookType) Value() (driver.Value, error) {
	return bt.String(), nil
}

// Scan satisfies the sql.Scanner interface.
func (bt *BookType) Scan(v interface{}) error {
	if buf, ok := v.([]byte); ok {
		return bt.UnmarshalText(buf)
	}
	return ErrInvalidBookType(fmt.Sprintf("%T", v))
}

// ErrInvalidBookType is the invalid BookType error.
type ErrInvalidBookType string

// Error satisfies the error interface.
func (err ErrInvalidBookType) Error() string {
	return fmt.Sprintf("invalid BookType(%s)", string(err))
}
