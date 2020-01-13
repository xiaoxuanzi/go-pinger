package main

import (
	"fmt"
	"os"
	"flag"
	"os/signal"

	"github.com/xiaoxianzi/go-pinger/g"
	"github.com/xiaoxianzi/go-pinger/pinger"
	"github.com/xiaoxianzi/go-pinger/http"
)

func main() {

	help  := flag.Bool("h", false, "help") 
	hosts := flag.String("hosts", "", "ip addresses/hosts to ping, space seperated")
	hostfile  := flag.String("hostfile", "", "host file, one host per line. this option will disable hosts option")
	enableWeb := flag.Bool("web", false, "enable webserver") 
	port := flag.Int("port", 8888, "web listen port(default 8888)")

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	if flag.NFlag() == 0 {
		fmt.Println("Usage: ")
		flag.PrintDefaults()
		os.Exit(2)
	}

	g.InitHosts(*hosts, *hostfile)

	// hosts reading from stdinput.
	// run command-line interface
	if *hosts != "" {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go pinger.InterruptHandler( c )
	}

	if *enableWeb {
		go http.Start( *port )
	}

	go pinger.Start()

	select {}
	
}
