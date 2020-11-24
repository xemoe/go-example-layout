package db

import (
	"github.com/spf13/cobra"

	_ "github.com/mattn/go-sqlite3"
)

var dbName string

// NewCmd represents the new command
var NewCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new db",
	Long:  `Create new db with required db file name (--db-filename).`,
	Run: func(cmd *cobra.Command, args []string) {
		NewSqliteDB(toAbsolutePath(withSqliteFileExtension(dbName)))
	},
}

func init() {
	//
	// flags: --db-name
	//
	NewCmd.PersistentFlags().StringVar(
		&dbName,
		"db-name",
		DefaultDBName,
		"database name")
}
