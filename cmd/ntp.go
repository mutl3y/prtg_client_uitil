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
	"github.com/mutl3y/prtg_dns/sensor"
	"github.com/spf13/cobra"
	"time"
)

// ntpCmd represents the ntp command
var ntpCmd = &cobra.Command{
	Use:   "ntp",
	Short: "Check localtime vs ntp server time, ",
	Long: `ntp client check for PRTG
returns error on too much drift

`,
	Run: func(cmd *cobra.Command, args []string) {

		flags := cmd.Flags()
		drift, err := flags.GetDuration("maxdrift")
		if err != nil {
			fmt.Println("drift", err)
		}

		timeout, err := flags.GetDuration("timeout")
		if err != nil {
			fmt.Println("timeout", err)
		}
		ntpHost, err := flags.GetString("ntphost")
		if err != nil {
			fmt.Println("ntpHost", err)
		}
		_ = sensor.PrtgNtp(ntpHost, timeout, drift)

	},
}

func init() {
	rootCmd.AddCommand(ntpCmd)
	ntpCmd.Flags().DurationP("maxdrift", "m", time.Duration(0), "max drift allowed, default 0 (disabled)")
	ntpCmd.Flags().StringP("ntphost", "n", "time.google.com", "ntp server to compare against local time")
}
