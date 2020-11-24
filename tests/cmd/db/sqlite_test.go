package db_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/xemoe/go-example-layout/cmd/db"
)

func TestNewSqliteDB(t *testing.T) {

	dir, err := ioutil.TempDir("", "db-sqlite3-test")
	if err != nil {
		return
	}

	defer func() {
		if err := os.RemoveAll(dir); err != nil {
			t.Error(err)
		}
	}()

	dbfile := dir + "/testdb.db"
	db.NewSqliteDB(dbfile)
}

func TestMigrateSqliteDB(t *testing.T) {

	dir, err := ioutil.TempDir("", "db-sqlite3-test")
	if err != nil {
		return
	}

	defer func() {
		if err := os.RemoveAll(dir); err != nil {
			t.Error(err)
		}
	}()

	dbfile := dir + "/testdb.db"
	db.NewSqliteDB(dbfile)
	db.MigrateSqliteDB(dbfile, db.DefaultMigrationFilesDir)
}