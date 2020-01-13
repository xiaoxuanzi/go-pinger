package pinger

import (
	"fmt"
	"log"
	"net"
	"time"
	"os"

	"github.com/xiaoxianzi/fatih/color"
	"github.com/xiaoxianzi/tatsushid/go-fastping"
	"github.com/xiaoxianzi/go-pinger/g"
)

var (
	Width          = "5"
	RowCounter     = 0
	DisplayCounter = 0
	Epoch          = 20
)

func InterruptHandler( c chan os.Signal ){

	for range c {
		pingerSummary()
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

func pingerSummary(){

	fmt.Printf("\n%-21s: %4s %4s %4s    %5s\n",
		"source", "min", "max", "avg", "ploss")
	hostIpMap := g.HostIpMap.GetAll()
	for host, ip := range hostIpMap {
		sl, ok := g.HistoryRttMap.Get(ip)
		if !ok {
			log.Println(ip + " not found in HostIpMap")
			continue
		}

		stats, err := sl.GetSummary()
		if err != nil {
			log.Println("GetSummary failed! host: ", host, " err: ", err)
			continue
		}

		fmt.Printf("+%-20s: %4d %4d %4d    %s\n",
			host, stats.Min, stats.Max, stats.Avg, stats.Ploss)
	}
}

func onIdle() func() {

	return func(){
		if DisplayCounter % Epoch == 0{
			displayHeader( g.InputHosts, Width )
		}

		fmt.Printf("%04d", RowCounter)

		for _, host := range( g.InputHosts ) {

			addr, ok := g.HostIpMap.Get(host)
			if !ok {
				continue
			}

			t, ok := g.LastRTT.Get(addr)
			if ok {
				color.Set(color.BgGreen, color.FgYellow, color.Bold)
				fmt.Printf("%" + Width + "d", t)
				color.Unset()

				continue

			}

			g.HistoryRttMap.PushFrontAndMaintain(addr, -1)
			color.Set(color.BgRed, color.FgYellow, color.Bold)
			fmt.Printf("%" + Width + "s", "!")
			color.Unset()
		}

		fmt.Println(" ")
		g.LastRTT = &g.SafeRttMap{M: make(map[string]int64)}
		
		RowCounter ++
		DisplayCounter ++

	}
}

func onRecv() func(addr *net.IPAddr, rtt time.Duration) {

	return func(addr *net.IPAddr, rtt time.Duration) {
		k := addr.String()
		t := rtt.Milliseconds()
		g.LastRTT.Set(k, t)
		g.HistoryRttMap.PushFrontAndMaintain(k, t)
	}
}

func Start(){

	p := fastping.NewPinger()

	for _, host := range g.InputHosts {

		ra, err := net.ResolveIPAddr("ip", host)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		p.AddIPAddr(ra)
		g.HostIpMap.Set(host, ra.String())

	}

	p.OnRecv = onRecv()
	p.OnIdle = onIdle()

	for{
		err := p.Run()
		if err != nil {
			log.Fatalln("[ERROR] pinger failed! ", err)
		}
	}
}
