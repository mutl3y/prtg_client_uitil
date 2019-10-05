## prtg_client_util deploy createUsers

create users on target / jumpphost

### Synopsis

create users on target / jumpphost

This command requires root privileges

if you prefer to do this yourself

	create $user on target and jumpohost if required
	mkdir /var/prtg/scriptsxml
	chown -r $user:$user /var/prtg/scriptsxml
	chmod 755 /var/prtg/scriptsxml


```
prtg_client_util deploy createUsers [flags]
```

### Options

```
  -h, --help                 help for createUsers
      --super_j string       admin user on jumphost system (default "root")
      --super_jpass string   password for super-j
      --super_t string       admin user on jumphost system (default "root")
      --super_tpass string   password for super-t
```

### Options inherited from parent commands

```
  -d, --debug              command line output
  -f, --j_KeyFile string   jumphost - private key file hint:C:\Users\mark\.ssh\id_rsa
  -i, --j_host string      jumphost - ip
  -p, --j_pass string      jumphost - password (default "prtgUtil")
  -o, --j_port string      jumphost - ssh port (default "22")
  -u, --j_user string      jumphost - user (default "prtgUtil")
  -R, --releases string    releases directory, defaults to $pwd/releases (default "D:\\goland\\prtg_dns\\releases")
  -F, --t_KeyFile string   target - private key file hint:C:\Users\mark\.ssh\id_rsa
  -I, --t_host string      target - ip (default "localhost")
  -P, --t_pass string      target - password (default "prtgUtil")
  -O, --t_port string      target - ssh port (default "22")
  -U, --t_user string      target - user (default "prtgUtil")
  -t, --timeout duration   timeout string eg 500ms (default 500ms)
```

### SEE ALSO

* [prtg_client_util deploy](prtg_client_util_deploy.md)	 - deploy prtg_client_util via ssh tunnel to target

