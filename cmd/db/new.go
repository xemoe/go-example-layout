package db

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// NewCmd represents the new command
var NewCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new sqlite3 db",
	Long:  `Create new sqlite3 db with default backup existing db file (--backup=true).`,
	Run: func(cmd *cobra.Command, args []string) {
		//
		// @TODO
		// - [ ] read db filepath from configuration.
		// - [ ] check existing sqlite3 db file.
		// - [ ] do backup if flag --backup=true.
		//   - generate new backup filename.
		//   - copy db file with generated filename.
		//   - remove old db file.
		// - [ ] create new sqlite3 db.
		//

		//
		// @TODO
		// - [ ] read dbfile value from configuration
		// - [ ] read dbpath value from configuration
		// - [ ] read --backup flag value
		// - [ ] if backup then do db_backup(dbfile)
		//
		dbfile := "./test.db"

		//
		// return if db file exist
		//
		if fileExists(dbfile) {
			log.WithFields(log.Fields{
				"db.file":   dbfile,
				"db.exists": true,
			}).Errorf("DB file: %s is exist", dbfile)

			return
		}

		db, err := sql.Open("sqlite3", dbfile)
		if err != nil {
			log.WithFields(log.Fields{
				"db.file":    dbfile,
				"db.created": false,
			}).Errorf("DB file: %s cannot create", dbfile)
			log.Error(err)

			return
		}
		defer db.Close()

		err = db.Ping()
		if err != nil {
			log.Error("Failed to connect to the source database:", err)
		}

		log.WithFields(log.Fields{
			"db.file":    dbfile,
			"db.created": true,
		}).Infof("DB file: %s has been create", dbfile)
	},
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
