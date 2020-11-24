package db

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DefaultDBName       = "test"
	DefaultMigrationFilesDir = "db/migrations"
	sqliteFileExtension = "db"
)

func toAbsolutePath(filename string) string {
	p, err := filepath.Abs(filename)
	if err != nil {
		panic(err)
	}
	return p
}

func withSqliteFileExtension(dbname string) string {
	ext := filepath.Ext(dbname)
	if ext != sqliteFileExtension {
		dbname = dbname + "." + sqliteFileExtension
	}

	return dbname
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// NewSqliteDB for create empty sqlite3 file
func NewSqliteDB(dbfile string) {

	if fileExists(dbfile) {
		log.WithFields(log.Fields{
			"db.file":   dbfile,
			"db.exists": true,
		}).Errorf("DB file is exist")

		return
	}

	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		log.WithFields(log.Fields{
			"db.file": dbfile,
			"db.open": false,
		}).Errorf("DB file  cannot be open")
		log.Error(err)

		return
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}
	}()

	err = db.Ping()
	if err != nil {
		log.Error("Failed to connect to the source database:", err)
		return
	}

	log.WithFields(log.Fields{
		"db.file":    dbfile,
		"db.created": true,
	}).Infof("DB file  has been create")
}

// MigrateSqliteDB for migrate sqlite3 db
func MigrateSqliteDB(dbfile string, migrationDir string) {

	log.WithFields(log.Fields{
		"db.migration": migrationDir,
	}).Infof("Read DB Migration file	")

	if !fileExists(dbfile) {
		log.WithFields(log.Fields{
			"db.file":   dbfile,
			"db.exists": false,
		}).Errorf("DB file is not exist")

		return
	}

	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		log.WithFields(log.Fields{
			"db.file": dbfile,
			"db.open": false,
		}).Errorf("DB file  cannot be open")
		log.Error(err)

		return
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}
	}()

	//
	// migrate
	//
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Error(err)
		return
	}

	m, err := migrate.NewWithDatabaseInstance("file://"+migrationDir, "ql", driver)
	if err != nil {
		log.Error(err)
		return
	}

	err = m.Up()
	if err != nil {

		if err == migrate.ErrNoChange {
			log.WithFields(log.Fields{
				"db.file":      dbfile,
				"db.migration": migrationDir,
			}).Info("Migration no change")
			return
		}

		log.Error(err)

		return
	}
}
