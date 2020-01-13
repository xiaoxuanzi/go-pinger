## go-pinger

This's a Simple tool to ping multiple hosts at once, provide command-line interface (CLI) and web-based overview, inspired by [gomphs](https://github.com/42wim/gomphs)

### Usage 
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

### Example
#### 1. Using web
<code>
$./go-pinger -hosts="slashdot.org www.linkedin.com github.com 39.156.69.79" -web
</code>
<br>
<img src="https://github.com/xiaoxuanzi/box/blob/master/go-pinger-web.png" />

#### 2. Command-line interface (CLI) 
<code>
$./go-pinger -hosts="slashdot.org www.linkedin.com github.com 39.156.69.79"
</code>
<img src="https://github.com/xiaoxuanzi/box/blob/master/go-pinger-example-1.gif"/>

#### 3. Import host from file
<code>
$./go-pinger -hostfile hostfile.txt  -web
</code>
