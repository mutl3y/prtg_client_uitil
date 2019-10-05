## prtg_client_util deploy

deploy prtg_client_util via ssh tunnel to target

### Synopsis


Deploy prtg_client_util via ssh tunnel to target

Functionality is restricted to the follwing unless --createUsers is called 
copy prtg_client_util compiled for destination OS to /var/prtg/scriptsxml
chmod 755 /var/prtg/scriptsxml

createUsers:
create user on target/jumpbox according to t_var and j_var parameters
adds keyfile to authoriset_keys if supplied

RSA key authentication is preferred however authentication will fall back to password if supplied


```
prtg_client_util deploy [flags]
```

### Options

```
  -h, --help               help for deploy
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
```

### Options inherited from parent commands

```
  -d, --debug              command line output
  -t, --timeout duration   timeout string eg 500ms (default 500ms)
```

### SEE ALSO

* [prtg_client_util](prtg_client_util.md)	 - simple prtg tests for remote nodes
* [prtg_client_util deploy createUsers](prtg_client_util_deploy_createUsers.md)	 - create users on target / jumpphost

