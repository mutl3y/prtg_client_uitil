# PRTG_Client_Util

Custom sensor for PRTG to allow you to check a clients view of

- DNS,  Can client resolve dns queries
- Ping, Ping hosts and measure response times and packet loss for multiple hosts
- NTP,  Measure time drift of client time to NTP server
- sshremote, run a version of this app remotely through a jumpbox/proxy
- genDocs,  Generate documentation

 

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

Please see individual docs for usage or use genDocs subcommand to generate man or rest format
prtg_util genDocs
* [prtg_client_util](docs/prtg_client_util.md)	 - simple prtg tests for remote nodes
* [prtg_client_util dns](docs/prtg_client_util_dns.md)	 - DNS resolve test
* [prtg_client_util genDocs](docs/prtg_client_util_genDocs.md)	 - Create documentation for app
* [prtg_client_util ntp](docs/prtg_client_util_ntp.md)	 - Check local time vs ntp server time, 
* [prtg_client_util ping](docs/prtg_client_util_ping.md)	 - Ping list of addresses, return avgRTT by default
* [prtg_client_util sshremote](docs/prtg_client_util_sshremote.md)	 - run command remotely through ssh tunnel via jumphost / proxy

If you feel like saying thanks    
        XMR: 49QA139gTEVMDV9LrTbx3qGKKEoYJucCtT4t5oUHHWfPBQbKc4MdktXfKSeT1ggoYVQhVsZcPAMphRS8vu8oxTf769NDTMu
	

With thanks to Jetbrains and their support of the open source community
![ https://www.jetbrains.com/?from=JJ-s-XMR-STAK-HashRate-Monitor-and-Restart-Tool](jetbrains-variant-3.png?v=4&s=200)
 
     

	
