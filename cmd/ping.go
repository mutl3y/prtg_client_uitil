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
	"time"

	"github.com/spf13/cobra"
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping list of addresses, return avgRTT by default",
	Long: `
Returns AvgRtt for list of addresses by default

Uses default gateway if addr not specified

timeout will be adjusted to be (count * interval)+interval

response time will vary depending on interval timer, 
10 * 1s interval = 10 seconds

Beware if you have IPS running in your network setting a low interval can be seen as packet loss

for example Fortigate firewalls drop udp pings that exceed 1 per second
`,
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		addr, err := flags.GetStringSlice("addr")
		if err != nil {
			fmt.Println(err)
		}

		timeout, err := flags.GetDuration("timeout")
		if err != nil {
			fmt.Println(err)
		}
		interval, err := flags.GetDuration("interval")
		if err != nil {
			fmt.Println(err)
		}
		count, err := flags.GetInt("count")
		if err != nil {
			fmt.Println(err)
		}
		size, err := flags.GetInt("size")
		if err != nil {
			fmt.Println(err)
		}
		d, err := flags.GetBool("debug")
		if err != nil {
			fmt.Println(err)
		}
		statsType, err := flags.GetString("type")
		if d {
			sensor.Debug = true
		}
		err = sensor.PrtgPing(addr, count, size, timeout, interval, statsType)
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	defgw, err := gateway.DiscoverGateway()
	if err != nil {
		defgw = net.ParseIP("127.0.0.1")

	}
	pingCmd.Flags().StringSliceP("addr", "a", []string{defgw.String()}, "comma separated hostname's or ip's")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	pingCmd.Flags().IntP("count", "c", 3, "how many pings")
	pingCmd.Flags().IntP("size", "s", 32, "packet size k")
	pingCmd.Flags().DurationP("interval", "i", 500*time.Millisecond, "timeout string eg 500ms, whole operation not per ping")
	pingCmd.Flags().StringP("type", "T", "", "leave blank for average response times\nloss\t packet loss\neverything\t all stats for first ip")
}
