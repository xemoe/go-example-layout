package db

import (
	"database/sql"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"

	_ "github.com/mattn/go-sqlite3"
)

const (
	defaultDbFileName = "test.db"
)

var dbFileName string

// NewCmd represents the new command
var NewCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new sqlite3 db",
	Long:  `Create new sqlite3 db with required db file name (--db-filename).`,
	Run: func(cmd *cobra.Command, args []string) {
		newDB(toAbsolutePath(dbFileName))
	},
}

func toAbsolutePath(filename string) string {
	p, err := filepath.Abs(filename)
	if err != nil {
		panic(err)
	}
	return p
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

func init() {
	//
	// flags: --db-filename
	//
	NewCmd.PersistentFlags().StringVar(
		&dbFileName,
		"db-filename",
		defaultDbFileName,
		"sqlite3 db filename")
}

func newDB(dbfile string) {

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
