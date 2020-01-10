package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-pinger/g"
	"github.com/gin-gonic/gin"
)

/*
//TODO
//CREATE STREAM WIHT dygraphs 
func pingerLatestRtt(c *gin.Context){
	lastRtt := g.LastRTT.GetAll()
	
	var keys   []string
	var values []int64

	keys = append(keys, "time")
	for k, v := range lastRtt{
		keys   = append(keys, k)
		values = append(values, v)
	}
	
	index := c.Query("index")
	if index == "labels" {
		c.JSON(http.StatusOK, map[string]interface{}{"data" : keys} )
		return 
	}

	c.JSON(http.StatusOK, map[string]interface{}{"data" : values} )
}

func pingerRttStream(c *gin.Context){
	c.HTML(http.StatusOK, "web.html",  gin.H{})
}
*/

func pingerResult(c *gin.Context){

	items := make(map[string]g.PingStats)
	hostIpMap := g.HostIpMap.GetAll()
	for host, ip := range hostIpMap {
		sl, ok := g.HistoryRttMap.Get(ip)
		if !ok {
			continue
		}

		stats, err := sl.GetSummary()
		if err != nil {
			log.Println("[ERRROR] ", err)
			continue
		}

		items[host] = stats
	}

	log.Println("[INFO] pingerResult: ", items)
	c.JSON(http.StatusOK, items )

}

func Router() *gin.Engine {

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context){
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "welcome to Pinger's website" })

	})

	pinger := router.Group("/pinger")
	{
		pinger.GET("/",  pingerResult)
		//TODO
		//pinger.GET("/latest",  pingerLatestRtt)
		//pinger.GET("/stream",  pingerRttStream)
	}

	return router
}

func Start(port int){

	listen := fmt.Sprintf("0.0.0.0:%d", port)
	r := Router()
	r.Run(listen)

}
