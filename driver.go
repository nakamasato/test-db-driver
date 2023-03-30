package driver

import (
	"database/sql"
	. "database/sql/driver"
)

func Register(name, drv, dsn string, options ...func(*testDBConn) error) {
	sql.Register(name, &testDBDriver{})
}

type testDBDriver struct{}

func (t *testDBDriver) Open(name string) (Conn, error) {
	return testDBConn{}, nil
}

type testDBConn struct{}

func (t testDBConn) Prepare(query string) (Stmt, error) {
	return testDBStmt{}, nil
}
func (t testDBConn) Close() error {
	return nil
}
func (t testDBConn) Begin() (Tx, error) {
	return testDBTx{}, nil
}

type testDBStmt struct{}

func (t testDBStmt) Close() error {
	return nil
}
func (t testDBStmt) NumInput() int {
	return 0
}
func (t testDBStmt) Exec(args []Value) (Result, error) {
	return testDBResult{}, nil
}

func (t testDBStmt) Query(args []Value) (Rows, error) {
	return testDBRows{}, nil
}

type testDBResult struct{}

func (t testDBResult) LastInsertId() (int64, error) {
	return 1, nil
}
func (t testDBResult) RowsAffected() (int64, error) {
	return 1, nil
}

type testDBRows struct{}

func (t testDBRows) Columns() []string {
	return []string{}
}

func (t testDBRows) Close() error {
	return nil
}

func (t testDBRows) Next(dest []Value) error {
	return nil
}

type testDBTx struct{}

func (t testDBTx) Commit() error {
	return nil
}

func (t testDBTx) Rollback() error {
	return nil
}
