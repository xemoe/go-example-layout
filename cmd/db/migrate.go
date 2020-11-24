package db

import (
	"github.com/spf13/cobra"
)

var migrationFilesDir string
var migrationDBName string

// migrateCmd represents the migrate command
var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate database table",
	Long:  `Migrate database table with optional migrate dir and db name`,
	Run: func(cmd *cobra.Command, args []string) {
		MigrateSqliteDB(
			toAbsolutePath(withSqliteFileExtension(migrationDBName)),
			toAbsolutePath(migrationFilesDir))
	},
}

func init() {
	//
	// flags: --migration-files-dir
	//
	MigrateCmd.PersistentFlags().StringVar(
		&migrationFilesDir,
		"migration-files-dir",
		DefaultMigrationFilesDir,
		"database migration files directory")

	//
	// flags: --db-name
	//
	MigrateCmd.PersistentFlags().StringVar(
		&migrationDBName,
		"db-name",
		DefaultDBName,
		"database name")
}
