## prtg_client_util ping

Ping list of addresses, return avgRTT by default

### Synopsis


Returns AvgRtt for list of addresses by default

Uses default gateway if addr not specified

timeout will be adjusted to be (count * interval)+interval

response time will vary depending on interval timer, 
10 * 1s interval = 10 seconds

Beware if you have IPS running in your network setting a low interval can be seen as packet loss

for example Fortigate firewalls drop udp pings that exceed 1 per second


```
prtg_client_util ping [flags]
```

### Options

```
  -a, --addr strings        comma separated hostname's or ip's (default [192.168.0.1])
  -c, --count int           how many pings (default 3)
  -h, --help                help for ping
  -i, --interval duration   timeout string eg 500ms, whole operation not per ping (default 500ms)
  -s, --size int            packet size k (default 32)
  -T, --type string         leave blank for average response times
                            loss	 packet loss
                            everything	 all stats for first ip
```

### Options inherited from parent commands

```
  -d, --debug              command line output
  -t, --timeout duration   timeout string eg 500ms (default 500ms)
```

### SEE ALSO

* [prtg_client_util](prtg_client_util.md)	 - simple prtg tests for remote nodes

