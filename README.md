## go-pinger

This's a Simple tool to ping multiple hosts at once, provide command-line interface (CLI) and web-based overview, inspired by [gomphs](https://github.com/42wim/gomphs)
### Installation
```
go get github.com/xiaoxuanzi/go-pinger
```
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
The web GUI is by default available via http://localhost:8888/pinger/web
Use -port to use another port.
<code>
$./go-pinger -hosts="slashdot.org www.linkedin.com github.com 39.156.69.79" -web
</code>
<br><br>
<img src="https://github.com/xiaoxuanzi/box/blob/master/go-pinger-web.png" />

#### 2. Command-line interface (CLI) 
The symbol '!' means host isn't reachable or down.

<code>
$./go-pinger -hosts="slashdot.org www.linkedin.com github.com 39.156.69.79"
</code>
<img src="https://github.com/xiaoxuanzi/box/blob/master/go-pinger-example-1.gif"/>

#### 3. Import host from file
<code>
$./go-pinger -hostfile hostfile.txt  -web
</code>

#### 4. Get result in json format
Via http://localhost:8888/pinger/json
```
{
    "39.156.69.79":{
        "max":4,
        "min":4,
        "Avg":4,
        "ploss":"0.00%"
    },
    "github.com":{
        "max":-1,
        "min":-1,
        "Avg":-1,
        "ploss":"100.00%"
    },
    "slashdot.org":{
        "max":205,
        "min":205,
        "Avg":205,
        "ploss":"0.00%"
    },
    "www.linkedin.com":{
        "max":22,
        "min":21,
        "Avg":21,
        "ploss":"0.00%"
    }
}
```
### License
MIT
