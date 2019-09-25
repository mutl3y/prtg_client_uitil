# PRTG_Client_Util

Custom sensor for PRTG to allow you to check a clients view of

- DNS,  Can client resolve dns queries
- Ping, Ping hosts and measure response times and packet loss for multiple hosts
- NTP,  Measure time drift of client time to NTP server
 

Tested with PRTG Version 19.3.51.2722
XMR-Stak Version 2.10.4

Place binary in PRTG folder on monitored host (this may be different on your install)
- Windows: C:\Program Files (x86)\PRTG Network Monitor\Custom Sensors\EXEXML
- Linux: /var/prtg/scriptsxml

This can also be compiled onto any Golang supported platform

Linux and windows versions will be found in the release pages of Github
## Note on Linux Support:

This library attempts to send an
"unprivileged" ping via UDP. 

On linux, this must be enabled on the client by setting

```
sudo sysctl -w net.ipv4.ping_group_range="0   2147483647"
```

See [this blog](https://sturmflut.github.io/linux/ubuntu/2015/01/17/unprivileged-icmp-sockets-on-linux/)
and [the Go icmp library](https://godoc.org/golang.org/x/net/icmp) for more details.

## To compile this yourself you need to...
-    install Golang
-    download or clone repo
-    run `go get` to download required packages
-    run `go build`
-    move the binary to the correct place
    
There are likely to be other small steps here as things may vary on your systems, If you need a OS binary and 
not in a rush drop me a request    

Add this to PRTG as an advanced custom exe / ssh script

```
./prtg_dns-linux-amd64 -h

simple dns resolve test for remote nodes

Usage:
  prtg_client_util dns [flags]

Flags:
  -a, --addr strings   up to 50 addresses (default [www.google.com,www.facebook.com])
  -h, --help           help for dns

Global Flags:
  -d, --debug              command line output
  -t, --timeout duration   timeout string eg 500ms (default 500ms)
```
```
./prtg_dns-linux-amd64 ping -h

Returns AvgRtt for list of addresses by default for PRTG

Uses default gateway if addr not specified

Examples:
        prtg_client_util-windows-amd64.exe ping -t 200ms  -a "192.168.0.1,8.8.8.8,8.8.4.4"

timeout will be adjusted to be (count * interval)+interval

response time will vary depending on interval timer,
10 * 1s interval = 10 seconds

Beware if you have IPS running in your network setting a low interval can be seen as packet loss

for example Fortigate firewalls drop udp pings that exceed 1 per second silently

Usage:
  prtg_client_util ping [flags]

Flags:
  -a, --addr strings        A help for foo (default [192.168.0.1])
  -c, --count int           how many pings (default 3)
  -h, --help                help for ping
  -i, --interval duration   timeout string eg 500ms, whole operation not per ping (default 500ms)
  -s, --size int            packet size k (default 32)
  -T, --type string         leave blank for average response times
                            loss         packet loss
                            everything   all stats for first ip

Global Flags:
  -d, --debug              command line output
  -t, --timeout duration   timeout string eg 500ms (default 500ms)
```


If you feel like saying thanks    
        XMR: 49QA139gTEVMDV9LrTbx3qGKKEoYJucCtT4t5oUHHWfPBQbKc4MdktXfKSeT1ggoYVQhVsZcPAMphRS8vu8oxTf769NDTMu
	

With thanks to Jetbrains and their support of the open source community
![ https://www.jetbrains.com/?from=JJ-s-XMR-STAK-HashRate-Monitor-and-Restart-Tool](jetbrains-variant-3.png?v=4&s=200)
 
     

	
