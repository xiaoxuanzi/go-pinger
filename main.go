package main

import (
	"fmt"
	"os"
	"os/signal"
	"flag"
	"strings"

	"github.com/go-pinger/g"
	"github.com/go-pinger/pinger"
	"github.com/go-pinger/http"
)

func main() {

	hosts := flag.String("hosts", "", "ip addresses/hosts to ping, space seperated")
	help  := flag.Bool("h", false, "help") 
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

	if *enableWeb {
		go http.Start( *port )
	}

	g.InputHosts = strings.Fields( *hosts )

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go pinger.InterruptHandler( c )

	go pinger.Start()

	select {}
	
}
