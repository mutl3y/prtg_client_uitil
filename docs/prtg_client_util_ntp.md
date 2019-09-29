## prtg_client_util ntp

Check local time vs ntp server time, 

### Synopsis

ntp client check for PRTG
returns error on too much drift



```
prtg_client_util ntp [flags]
```

### Options

```
  -h, --help                help for ntp
  -m, --maxdrift duration   max drift allowed, default 0 (disabled)
  -n, --ntphost string      ntp server to compare against local time (default "time.google.com")
```

### Options inherited from parent commands

```
  -d, --debug              command line output
  -t, --timeout duration   timeout string eg 500ms (default 500ms)
```

### SEE ALSO

* [prtg_client_util](prtg_client_util.md)	 - simple prtg tests for remote nodes

