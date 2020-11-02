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
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.WithFields(log.Fields{
			"flag.port": v.Get("port"),
		}).Debugf("Flag port: %d", v.Get("port"))
	},
}

func init() {
	apiCmd.Flags().IntP("port", "p", 8088, "Api Bind port address")
	v.BindPFlag("port", apiCmd.Flags().Lookup("port"))

	rootCmd.AddCommand(apiCmd)
}