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
	"github.com/mutl3y/prtg_client_util/util"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// deployCmd represents the deploy command

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "deploy prtg_client_util remotely through ssh tunnel to jumphost or endpoint",
	Long: `deploy prtg_client_util remotely through ssh tunnel to jumphost or endpoint

Functionality is restricted to copying prtg_client_util to /var/prtg/scriptsxml

Only a basic user account should be created for this, no sudo rights etc but will need permissions 
to add files and set executable bit in /var/prtg/scriptsxml

sets files to perm 755 with user specified as owner

RSA key authentication is preferred however authentication will fall back to password if supplied
`,
	Run: func(cmd *cobra.Command, args []string) {

		flags := cmd.Flags()
		d_host, err := flags.GetString("d_host")
		handleVarError(err)
		d_port, err := flags.GetString("d_port")
		handleVarError(err)
		d_user, err := flags.GetString("d_user")
		handleVarError(err)
		d_pass, err := flags.GetString("d_pass")
		handleVarError(err)
		d_KeyFile, err := flags.GetString("d_KeyFile")
		handleVarError(err)
		timeout, err := flags.GetDuration("timeout")
		handleVarError(err)

		d := util.SshStruct{
			User:     d_user,
			Server:   d_host,
			KeyPath:  d_KeyFile,
			Port:     d_port,
			Password: d_pass,
			Timeout:  timeout,
		}

		p_host, err := flags.GetString("p_host")
		handleVarError(err)
		p_port, err := flags.GetString("p_port")
		handleVarError(err)
		p_user, err := flags.GetString("p_user")
		handleVarError(err)
		p_pass, err := flags.GetString("p_pass")
		handleVarError(err)
		p_KeyFile, err := flags.GetString("p_KeyFile")
		handleVarError(err)

		dir, err := flags.GetString("dir")
		dir = filepath.Base(dir)
		handleVarError(err)

		p := util.SshStruct{
			User:     p_user,
			Server:   p_host,
			KeyPath:  p_KeyFile,
			Port:     p_port,
			Password: p_pass,
			Timeout:  timeout,
		}

		rem := util.NewCon(d, p)
		_ = rem.Deploy(dir)

	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("failed to get present working directory")
		return
	}
	deployCmd.Flags().StringP("dir", "R", strings.Join([]string{pwd, "releases"}, string(os.PathSeparator)), "releases directory, defaults to $pwd/releases")

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("failed to get users home directory %v", err)
	}
	defaultSshKey := strings.Join([]string{home, ".ssh", "id_rsa"}, string(os.PathSeparator))

	deployCmd.Flags().StringP("d_host", "D", "localhost", "destination host ip")
	deployCmd.Flags().StringP("d_port", "O", "22", "ssh port on dest")
	deployCmd.Flags().StringP("d_user", "U", "prtgUtil", "user on destination")
	deployCmd.Flags().StringP("d_pass", "P", "prtgUtil", "destination user password")
	deployCmd.Flags().StringP("d_KeyFile", "F", defaultSshKey, "destination private keyfile")

	deployCmd.Flags().StringP("p_host", "H", "", "proxy host ip")
	deployCmd.Flags().StringP("p_port", "o", "22", "ssh port on proxy")
	deployCmd.Flags().StringP("p_user", "u", "prtgUtil", "user on proxy")
	deployCmd.Flags().StringP("p_pass", "p", "prtgUtil", "proxy user password")
	deployCmd.Flags().StringP("p_KeyFile", "f", defaultSshKey, "proxy private keyfile")
	_ = deployCmd.MarkFlagRequired("dest")
	_ = deployCmd.MarkFlagRequired("proxy")

}
