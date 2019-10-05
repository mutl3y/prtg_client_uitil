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
a copy of the app must be placed in that folder with execute permissions for remote user

Be aware this effectively allows PRTG to perform remote code execution.

Only a basic user account should be created for this, no sudo rights etc.

RSA key authentication is preferred however authentication will fall back to password if supplied

`,
	Run: func(cmd *cobra.Command, args []string) {

		flags := cmd.Flags()
		run, err := flags.GetString("run")
		handleVarError(err)
		t_host, err := flags.GetString("t_host")
		handleVarError(err)
		t_port, err := flags.GetString("t_port")
		handleVarError(err)
		t_user, err := flags.GetString("t_user")
		handleVarError(err)
		t_pass, err := flags.GetString("t_pass")
		handleVarError(err)
		t_KeyFile, err := flags.GetString("t_KeyFile")
		handleVarError(err)
		timeout, err := flags.GetDuration("timeout")
		handleVarError(err)

		d := util.SshStruct{
			User:     t_user,
			Server:   t_host,
			KeyPath:  t_KeyFile,
			Port:     t_port,
			Password: t_pass,
			Timeout:  timeout,
		}

		j_host, err := flags.GetString("j_host")
		handleVarError(err)
		j_port, err := flags.GetString("j_port")
		handleVarError(err)
		j_user, err := flags.GetString("j_user")
		handleVarError(err)
		j_pass, err := flags.GetString("j_pass")
		handleVarError(err)
		j_KeyFile, err := flags.GetString("j_KeyFile")
		handleVarError(err)

		p := util.SshStruct{
			User:     j_user,
			Server:   j_host,
			KeyPath:  j_KeyFile,
			Port:     j_port,
			Password: j_pass,
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
	defaultSshKey := strings.Join([]string{home, ".ssh", "it_rsa"}, string(os.PathSeparator))

	sshremoteCmd.Flags().StringP("t_host", "I", "localhost", "target - ip")
	sshremoteCmd.Flags().StringP("t_port", "O", "22", "target - ssh port")
	sshremoteCmd.Flags().StringP("t_user", "U", "prtgUtil", "target - user")
	sshremoteCmd.Flags().StringP("t_pass", "P", "prtgUtil", "target - password")
	sshremoteCmd.Flags().StringP("t_KeyFile", "F", "", "target - private key file hint:"+defaultSshKey)

	sshremoteCmd.Flags().StringP("j_host", "i", "", "jumphost - ip")
	sshremoteCmd.Flags().StringP("j_port", "o", "22", "jumphost - ssh port")
	sshremoteCmd.Flags().StringP("j_user", "u", "prtgUtil", "jumphost - user")
	sshremoteCmd.Flags().StringP("j_pass", "p", "prtgUtil", "jumphost - password")
	sshremoteCmd.Flags().StringP("j_KeyFile", "f", "", "jumphost - private key file")

}

func handleVarError(err error) {
	if err == nil {
		return
	} else {
		log.Fatalf("failed to get variable %v", err)
	}
}
