/*
Copyright © 2019 mutl3y

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
	"github.com/jackpal/gateway"
	"github.com/mutl3y/prtg_dns/sensor"
	"net"

	"github.com/spf13/cobra"
)

// lossCmd represents the loss command
var lossCmd = &cobra.Command{
	Use:   "loss",
	Short: "Returns packet loss for list of addresses / IP for use with PRTG",
	Long: `
Returns packet loss for list of addresses / IP for use with PRTG

Uses default gateway if addr not specified

Examples:
	prtg_dns-windows-amd64.exe loss -t 200ms  
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
		c, err := flags.GetInt("count")
		if err != nil {
			fmt.Println(err)
		}

		err = sensor.PrtgPacketLoss(a, c, t)
	},
}

func init() {
	rootCmd.AddCommand(lossCmd)
	defgw, err := gateway.DiscoverGateway()
	if err != nil {
		defgw = net.ParseIP("127.0.0.1")

	}
	lossCmd.Flags().StringSliceP("addr", "a", []string{defgw.String()}, "help for loss")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	lossCmd.Flags().IntP("count", "c", 3, "how many pings")
}
