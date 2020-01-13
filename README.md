## go-pinger

This's a Simple tool to ping multiple hosts at once, provide command-line interface (CLI) and web-based overview, inspired by [gomphs](https://github.com/42wim/gomphs)

Usage 

```
# ./go-pinger
Usage: 
  -h	help
  -hostfile string
    	host file, one host per line. this option will disable hosts option
  -hosts string
    	ip addresses/hosts to ping, space seperated
  -port int
    	web listen port(default 8888) (default 8888)
  -web
    	enable webserver
```
#### example

##### command-line interface (CLI) 
* command：
```
./go-pinger -hosts="slashdot.org www.linkedin.com github.com 39.156.69.79"
```
* result
<img src="https://github.com/xiaoxuanzi/box/blob/master/go-pinger-example-1.gif"  width="80%"  height="80%"/>

<img src="https://github.com/xiaoxuanzi/box/blob/master/go-pinger-web.png"  width="80%"  height="80%"/>
