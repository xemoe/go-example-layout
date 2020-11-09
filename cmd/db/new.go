package db

import (
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
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

		filename := toAbsolutePath(dbFileName)

		//
		// return if db file exist
		//
		if fileExists(filename) {
			log.WithFields(log.Fields{
				"db.file":   filename,
				"db.exists": true,
			}).Errorf("DB file is exist")

			return
		}

		db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})
		if err != nil {
			log.WithFields(log.Fields{
				"db.file":    filename,
				"db.created": false,
			}).Errorf("DB file  cannot create")
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
			"db.file":    filename,
			"db.created": true,
		}).Infof("DB file  has been create")
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
	NewCmd.PersistentFlags().StringVar(
		&dbFileName,
		"db-filename",
		defaultDbFileName,
		"sqlite3 db filename (default is test.db)")
}
