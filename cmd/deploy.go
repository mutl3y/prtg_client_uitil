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

var Debug bool

// deployCmd represents the deploy command

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "deploy prtg_client_util via ssh tunnel to target",
	Long: `
Deploy prtg_client_util via ssh tunnel to target

Functionality is restricted to the follwing unless --createUsers is called 
copy prtg_client_util compiled for destination OS to /var/prtg/scriptsxml
chmod 755 /var/prtg/scriptsxml

createUsers:
create user on target/jumpbox according to t_var and j_var parameters
adds keyfile to authoriset_keys if supplied

RSA key authentication is preferred however authentication will fall back to password if supplied
`,
	Run: func(cmd *cobra.Command, args []string) {

		flags := cmd.Flags()
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

		t := util.SshStruct{
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

		dir, err := flags.GetString("releases")
		dir = filepath.Base(dir)
		handleVarError(err)

		j := util.SshStruct{
			User:     j_user,
			Server:   j_host,
			KeyPath:  j_KeyFile,
			Port:     j_port,
			Password: j_pass,
			Timeout:  timeout,
		}
		Debug, err = flags.GetBool("debug")
		handleVarError(err)
		rem := util.NewCon(t, j)
		debug("connection \t%v\n", *rem)
		err = rem.Deploy(dir)
		if err != nil {
			fmt.Println(err)
		}

	},
}

func debug(format string, a ...interface{}) {
	if Debug {
		fmt.Printf(format+"\n", a...)
	}
}

func init() {
	rootCmd.AddCommand(deployCmd)
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("failed to get present working directory")
		return
	}
	deployCmd.PersistentFlags().StringP("releases", "R", strings.Join([]string{pwd, "releases"}, string(os.PathSeparator)), "releases directory, defaults to $pwd/releases")

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("failed to get users home directory %v", err)
	}
	defaultSshKey := strings.Join([]string{home, ".ssh", "id_rsa"}, string(os.PathSeparator))

	deployCmd.PersistentFlags().StringP("t_host", "I", "localhost", "target - ip")
	deployCmd.PersistentFlags().StringP("t_port", "O", "22", "target - ssh port")
	deployCmd.PersistentFlags().StringP("t_user", "U", "prtgUtil", "target - user")
	deployCmd.PersistentFlags().StringP("t_pass", "P", "prtgUtil", "target - password")
	deployCmd.PersistentFlags().StringP("t_KeyFile", "F", "", "target - private key file hint:"+defaultSshKey)

	deployCmd.PersistentFlags().StringP("j_host", "i", "", "jumphost - ip")
	deployCmd.PersistentFlags().StringP("j_port", "o", "22", "jumphost - ssh port")
	deployCmd.PersistentFlags().StringP("j_user", "u", "prtgUtil", "jumphost - user")
	deployCmd.PersistentFlags().StringP("j_pass", "p", "prtgUtil", "jumphost - password")
	deployCmd.PersistentFlags().StringP("j_KeyFile", "f", "", "jumphost - private key file hint:"+defaultSshKey)

}
