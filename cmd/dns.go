/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"github.com/mutl3y/prtg_client_util/sensor"

	"github.com/spf13/cobra"
)

// dnsCmd represents the dns command
var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "DNS resolve test",
	Long: `
simple dns resolve test for remote nodes

`,
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		a, err := flags.GetStringSlice("addr")
		if err != nil {
			fmt.Println(err)
		}

		t, err := flags.GetDuration("timeout")
		if err != nil {
			fmt.Println(err)
		}
		err = sensor.PrtgLookup(a, t)
	},
}

func init() {
	rootCmd.AddCommand(dnsCmd)
	dnsCmd.Flags().StringSliceP("addr", "a", []string{"www.google.com", "www.facebook.com"}, "up to 50 addresses")

}
