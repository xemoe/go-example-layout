package db

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

const (
	defaultDbFilename = "test.db"
)

var dbFilename string

// NewCmd represents the new command
var NewCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new sqlite3 db",
	Long:  `Create new sqlite3 db with required db file name (--db-filename).`,
	Run: func(cmd *cobra.Command, args []string) {

		//
		// return if db file exist
		//
		if fileExists(dbFilename) {
			log.WithFields(log.Fields{
				"db.file":   dbFilename,
				"db.exists": true,
			}).Errorf("DB file: %s is exist", dbFilename)

			return
		}

		db, err := gorm.Open(sqlite.Open(dbFilename), &gorm.Config{})
		if err != nil {
			log.WithFields(log.Fields{
				"db.file":    dbFilename,
				"db.created": false,
			}).Errorf("DB file: %s cannot create", dbFilename)
			log.Error(err)

			return
		}

		geDB, err := db.DB()
		if err != nil {
			log.Error("Failed to get generic DB():", err)
			return
		}

		defer geDB.Close()

		err = geDB.Ping()
		if err != nil {
			log.Error("Failed to connect to the source database:", err)
			return
		}

		log.WithFields(log.Fields{
			"db.file":    dbFilename,
			"db.created": true,
		}).Infof("DB file: %s has been create", dbFilename)
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

func init() {
	NewCmd.PersistentFlags().StringVar(
		&dbFilename,
		"db-filename",
		defaultDbFilename,
		"sqlite3 db filename (default is test.db)")
}
