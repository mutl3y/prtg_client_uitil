###### PRTG_dns

Custom sensor for PRTG to allow you to check a host can resolce dns queries

Tested with PRTG Version 19.3.51.2722
XMR-Stak Version 2.10.4

This sensor can be run from any PRTG probe however can also be run on the target machine or another in the same 
network as it only supports http

If running on a PRTG server itself copy the windows binary into the appropriate folder for PRTG
on a default install this would be Directory of C:\Program Files (x86)\PRTG Network Monitor\Custom Sensors\EXEXML

This can also be compiled onto any Golang supported platform, Linux and windows versions will be found in 
the release pages of Github

To compile this yourself you need to...
    1, install Golang
    2, download or clone repo
    3, run go get to download required packages
    4, run go build
    5, move the binary to the correct place
    
There are likely to be other small steps here as things may vary on your systems, If you need a OS binary and 
not in a rush drop me a request    

C:\prtg_dns.exe -h
simple dns resolve test for remote nodes

Usage:
  prtg_dns [flags]

Flags:
  -a, --addresses strings   1 to 50 domains (default [www.google.com,www.facebook.com])
  -h, --help                help for prtg_dns
  -t, --timeout duration    timeout string eg 500ms (default 500ms)


If you feel like saying thanks    
        XMR: 49QA139gTEVMDV9LrTbx3qGKKEoYJucCtT4t5oUHHWfPBQbKc4MdktXfKSeT1ggoYVQhVsZcPAMphRS8vu8oxTf769NDTMu
	

With thanks to Jetbrains and their support of the open source community
![ https://www.jetbrains.com/?from=JJ-s-XMR-STAK-HashRate-Monitor-and-Restart-Tool](jetbrains-variant-3.png?v=4&s=200)
 
     

	
