This tools is for scanning port to scan website public port. This is really useful for network security. Inspired by the book  [Black Hat Go](https://www.amazon.com/Black-Hat-Go-Programming-Pentesters/dp/1593278659)

## Commands

    ports_scan
    or
    go run ports_scan.go

## Options

- ``-host``: host name, default is "scanme.nmap.org"
- ``-min``: minimum port number to scan, default is "0"
- ``-max``: maximum port number to scan, default is "100"
- ``-worker``: number of worker, default is 5
 
 ## Example
 
    >> ports_scan -host=scanme.nmap.org  -min=70 -max=500
    
    #output
    80 open
    443 open

