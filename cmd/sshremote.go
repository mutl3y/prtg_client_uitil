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
	"github.com/mutl3y/prtg_client_util/util"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

// sshremoteCmd represents the sshremote command
var sshremoteCmd = &cobra.Command{
	Use:   "sshremote",
	Short: "run command remotely through ssh tunnel via jumphost / proxy",
	Long: `run command remotely through ssh tunnel

Functionality is restricted to running prtg_client_util remotely from /var/prtg/scriptsxml.
a copy of the app must be p[laced in that folder with execute permissions for remote user

Be aware this effectively allows PRTG to perform remote code execution.

Only a basic user account should be created for this, no sudo rights etc.

RSA key authentication is preferred however authentication will fall back to password if supplied

`,
	Run: func(cmd *cobra.Command, args []string) {

		flags := cmd.Flags()
		run, err := flags.GetString("run")
		handleVarError(err)
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

		p := util.SshStruct{
			User:     p_user,
			Server:   p_host,
			KeyPath:  p_KeyFile,
			Port:     p_port,
			Password: p_pass,
			Timeout:  timeout,
		}

		rem := util.NewCon(d, p)
		_ = rem.Remote(run, timeout)
	},
}

func init() {
	rootCmd.AddCommand(sshremoteCmd)
	sshremoteCmd.Flags().StringP("run", "R", "ping", "command to run on remote host")

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("failed to get users home directory %v", err)
	}
	defaultSshKey := strings.Join([]string{home, ".ssh", "id_rsa"}, string(os.PathSeparator))

	sshremoteCmd.Flags().StringP("d_host", "D", "localhost", "destination host ip")
	sshremoteCmd.Flags().StringP("d_port", "O", "22", "ssh port on dest")
	sshremoteCmd.Flags().StringP("d_user", "U", "prtgUtil", "user on destination")
	sshremoteCmd.Flags().StringP("d_pass", "P", "prtgUtil", "destination user password")
	sshremoteCmd.Flags().StringP("d_KeyFile", "F", defaultSshKey, "destination private keyfile")

	sshremoteCmd.Flags().StringP("p_host", "H", "", "proxy host ip")
	sshremoteCmd.Flags().StringP("p_port", "o", "22", "ssh port on proxy")
	sshremoteCmd.Flags().StringP("p_user", "u", "prtgUtil", "user on proxy")
	sshremoteCmd.Flags().StringP("p_pass", "p", "prtgUtil", "proxy user password")
	sshremoteCmd.Flags().StringP("p_KeyFile", "f", defaultSshKey, "proxy private keyfile")
	_ = sshremoteCmd.MarkFlagRequired("dest")
	_ = sshremoteCmd.MarkFlagRequired("proxy")

}

func handleVarError(err error) {
	if err == nil {
		return
	} else {
		log.Fatalf("failed to get variable", err)
	}
}
