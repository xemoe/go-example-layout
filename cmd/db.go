/*
Copyright © 2020 Teerapong Ladlee <blckpearl.sheeper@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xemoe/go-example-layout/cmd/db"
)

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	//
	// @TODO
	// - [ ] newCmd `db new --backup-db true`
	// - [ ] migrateCmd `db migrate`
	// - [ ] backupCmd `db backup --backup-dir ./backup`
	// - [ ] restoreCmd `db restore --from-file backup.db`
	// - [ ] dumpCmd `db dump`
	//
	dbCmd.AddCommand(db.NewCmd)
	rootCmd.AddCommand(dbCmd)
}
