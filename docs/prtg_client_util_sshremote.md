## prtg_client_util sshremote

run command remotely through ssh tunnel via jumphost / proxy

### Synopsis

run command remotely through ssh tunnel

Functionality is restricted to running prtg_client_util remotely from /var/prtg/scriptsxml.
a copy of the app must be placed in that folder with execute permissions for remote user

Be aware this effectively allows PRTG to perform remote code execution.

Only a basic user account should be created for this, no sudo rights etc.

RSA key authentication is preferred however authentication will fall back to password if supplied



```
prtg_client_util sshremote [flags]
```

### Options

```
  -F, --d_KeyFile string   destination private keyfile (default "C:\\Users\\mark\\.ssh\\id_rsa")
  -D, --d_host string      destination host ip (default "localhost")
  -P, --d_pass string      destination user password (default "prtgUtil")
  -O, --d_port string      ssh port on dest (default "22")
  -U, --d_user string      user on destination (default "prtgUtil")
  -h, --help               help for sshremote
  -f, --p_KeyFile string   proxy private keyfile (default "C:\\Users\\mark\\.ssh\\id_rsa")
  -H, --p_host string      proxy host ip
  -p, --p_pass string      proxy user password (default "prtgUtil")
  -o, --p_port string      ssh port on proxy (default "22")
  -u, --p_user string      user on proxy (default "prtgUtil")
  -R, --run string         command to run on remote host (default "ping")
```

### Options inherited from parent commands

```
  -d, --debug              command line output
  -t, --timeout duration   timeout string eg 500ms (default 500ms)
```

### SEE ALSO

* [prtg_client_util](prtg_client_util.md)	 - simple prtg tests for remote nodes

