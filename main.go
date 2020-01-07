package main

import (
	"fmt"
	//"log"
	"net"
	//"sync"
	//"net/http"
	"time"
	"os"
	"os/signal"
	"flag"
	"strings"


	"github.com/fatih/color"
	"github.com/tatsushid/go-fastping"
	//"github.com/go-pinger/g"
)

/*
type SafeHostIpMap struct{
	sync.RWMutex
	M map[string]string
}

type SafeHostIpMap struct{
	sync.RWMutex
	M map[string]string
}
*/

var (

	width = "5"
	rowCounter = 0
	displayCounter = 0
	epoch = 20

	inputHosts []string

	hostIpMap  map[string]string
	historyRTT map[string][]int64
	lastRTT    map[string]int64
)

func printStat(){
	fmt.Println("chosed")
}

func interruptHandler( c chan os.Signal ){

	for range c {
		printStat()
		os.Exit(0)
	}
}

func displayHeader(hosts []string, width string){

	fmt.Println(" ")
	for i, host := range( hosts ) {
		fmt.Printf("  %d = %s\n", i, host)
	}

	fmt.Printf("%4s", "")
	for i, _ := range( hosts ) {
		fmt.Printf("%" + width + "d", i)
	}

	fmt.Println(" ")
}

func onRecv() func(addr *net.IPAddr, rtt time.Duration) {

	return func(addr *net.IPAddr, rtt time.Duration) {
		k := addr.String()
		t := rtt.Milliseconds()
		lastRTT[ k ] = t
		historyRTT[ k ] = append(historyRTT[ k ], t)
	}

}


func onIdle() func() {

	return func(){
		if displayCounter % epoch == 0{
			displayHeader( inputHosts, width )
		}

		fmt.Printf("%04d", rowCounter)

		for _, host := range( inputHosts ) {

			addr := hostIpMap[ host ]

			t, ok := lastRTT[ addr ]
			if ok {
				color.Set(color.BgGreen, color.FgYellow, color.Bold)
				fmt.Printf("%" + width + "d", t)
				color.Unset()

				continue

			}

			historyRTT[ addr ] = append(historyRTT[ addr ], -1)
			color.Set(color.BgRed, color.FgYellow, color.Bold)
			fmt.Printf("%" + width + "s", "!")
			color.Unset()
		}

		fmt.Println(" ")
		lastRTT = make(map[string]int64)
		
		rowCounter ++
		displayCounter ++

	}
}

func main() {

	hosts := flag.String("hosts", "", "ip addresses/hosts to ping, space seperated")
	help  := flag.Bool("h", false, "help") 

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

	inputHosts = strings.Fields( *hosts )

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go interruptHandler( c )
	
	hostIpMap  = make(map[string]string)
	historyRTT = make(map[string][]int64)
	lastRTT    = make(map[string]int64)

	p := fastping.NewPinger()

	for _, host := range inputHosts {

		ra, err := net.ResolveIPAddr("ip", host)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		p.AddIPAddr(ra)
		hostIpMap[ host ] = ra.String()

	}

	p.OnRecv = onRecv()
	p.OnIdle = onIdle()

	for{
		err := p.Run()
		if err != nil {
			fmt.Println(err)
		}
	}

}
