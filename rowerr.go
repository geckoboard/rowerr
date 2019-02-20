// Package rowerr provides a mechanism for constructing a *sql.Row rigged to
// return an error when you call Scan() on it, something the standard library
// doesn't make possible. This is useful for testing error-handling when you're
// using the QueryRow() and QueryRowContext() methods on *sql.DB and *sql.Tx.
package rowerr

import (
	"database/sql"
	"reflect"
	"unsafe"
)

// New constructs an *sql.Row that will return err when you call Scan() on it.
func New(err error) *sql.Row {
	row := &sql.Row{}

	// Fetch the unexported err field using reflection.
	errField := reflect.ValueOf(row).Elem().FieldByName("err")

	// Construct a new error field at the same memory address, so that we're able
	// to control its value.
	newErrField := reflect.NewAt(
		errField.Type(),
		unsafe.Pointer(errField.UnsafeAddr()),
	).Elem()

	// Set the value of the error field.
	newErrField.Set(reflect.ValueOf(err))

	return row
}
