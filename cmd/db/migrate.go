package db

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	log "github.com/sirupsen/logrus"
	"os"

	"github.com/spf13/cobra"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

const (
	defaultMigrationFileDir = "db/migrations"
)

var migrationFileDir string
var migrationDbFileName string

// migrateCmd represents the migrate command
var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate sqlite3 table",
	Long:  `Migrate sqlite3 table with required migrate dir and db name`,
	Run: func(cmd *cobra.Command, args []string) {
		doMigrate(
			toAbsolutePath(migrationDbFileName),
			toAbsolutePath(migrationFileDir))
	},
}

func init() {
	//
	// flags: --migration-file-dir
	//
	MigrateCmd.PersistentFlags().StringVar(
		&migrationFileDir,
		"migration-file-dir",
		defaultMigrationFileDir,
		"sqlite3 db migration file directory")

	//
	// flags: --db-filename
	//
	MigrateCmd.PersistentFlags().StringVar(
		&migrationDbFileName,
		"db-filename",
		defaultDbFileName,
		"sqlite3 db filename")
}

func doMigrate(dbfile string, migrationDir string) {

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
	// migrations
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
				"db.file": dbfile,
				"db.migration": migrationDir,
			}).Info("Migration no change")
			return
		}

		log.Error(err)

		return
	}
}