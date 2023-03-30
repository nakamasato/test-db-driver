package driver_test

import (
	"database/sql"
	"fmt"
	"testing"

	driver "github.com/nakamasato/test-db-driver"
)

func init() {
	driver.Register("testdbdriver", "testdb", "test_user@/test_db")
}

func TestMain(t *testing.T) {
	db, err := sql.Open("testdbdriver", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	res, err := db.Exec("SELECT * FROM user")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}
