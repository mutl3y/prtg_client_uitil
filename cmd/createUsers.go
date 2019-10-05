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
	"github.com/mutl3y/prtg_client_util/util"
	"log"
	"path/filepath"

	"github.com/spf13/cobra"
)

// createUsersCmd represents the createUsers command
var createUsersCmd = &cobra.Command{
	Use:   "createUsers",
	Short: "create users on target / jumpphost",
	Long: `create users on target / jumpphost

This command requires root privileges

if you prefer to do this yourself

	create $user on target and jumpohost if required
	mkdir /var/prtg/scriptsxml
	chown -r $user:$user /var/prtg/scriptsxml
	chmod 755 /var/prtg/scriptsxml
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
		tuser, err := flags.GetString("super_t")
		handleVarError(err)
		tpass, err := flags.GetString("super_tpass")
		handleVarError(err)
		juser, err := flags.GetString("super_j")
		handleVarError(err)
		jpass, err := flags.GetString("super_jpass")
		if err != nil {
			jpass = j_pass
		}
		handleVarError(err)
		debug("create users \tTarget %v:%q Jumphost %v:%q\n", tuser, tpass, juser, jpass)
		err = rem.CreateUsers(tuser, tpass, juser, jpass)
		if err != nil {
			log.Fatal(err)
		}
		err = rem.Deploy(dir)
		if err != nil {
			fmt.Println(err)
		}

	},
}

func init() {
	deployCmd.AddCommand(createUsersCmd)

	// target creds for creating new user
	createUsersCmd.Flags().String("super_t", "root", "admin user on jumphost system")
	createUsersCmd.Flags().String("super_tpass", "", "password for super-t")

	// jumphost creds for creating new user
	createUsersCmd.Flags().String("super_j", "root", "admin user on jumphost system")
	createUsersCmd.Flags().String("super_jpass", "", "password for super-j")
}
