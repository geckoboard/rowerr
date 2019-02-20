# rowerr

rowerr provides a mechanism for constructing a
[`*sql.Row`](https://golang.org/pkg/database/sql/#Row) rigged to return an
error when you call [`Scan()`](https://golang.org/pkg/database/sql/#Row.Scan)
on it.

This is useful for testing error-handling when you're using the
[`QueryRow()`](https://golang.org/pkg/database/sql/#DB.QueryRow) and
[`QueryRowContext()`](https://golang.org/pkg/database/sql/#DB.QueryRowContext)
methods on [`*sql.DB`](https://golang.org/pkg/database/sql/#DB) and 
[`*sql.Tx`](https://golang.org/pkg/database/sql/#Tx).

## Example

```Go
// foo.go
type driver interface {
    QueryRow(string, ...interface{}) *sql.Row
}

func foo(db driver) error {
    var v int
    err := db.QueryRow("SELECT ... FROM ...").Scan(&v)
    if err == sql.ErrNoRows {
        return nil
    }
    return err
}

// foo_test.go
type fakeDriver struct {
    err error
}

func (f fakeDriver) QueryRow(string, ...interface{}) *sql.Row {
    return rowerr.New(f.err)
}

func TestFoo_ErrNoRows(t *testing.T) {
    driver := fakeDriver{err: sql.ErrNoRows}
    if err := foo(driver); err != nil {
        t.Error("expected no error, got: %v", err)
    }
}

```

## License

rowerr is licensed under the [MIT License](https://opensource.org/licenses/MIT).
