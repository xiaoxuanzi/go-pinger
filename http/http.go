package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/xiaoxuanzi/go-pinger/g"
	"github.com/xiaoxuanzi/gin-gonic/gin"
)

func pingerHostStats() map[string]g.PingStats{
	hostStats := make(map[string]g.PingStats)
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

		hostStats[host] = stats
	}

	return hostStats
}

func pingerWeb(c *gin.Context){

	hostStats := pingerHostStats()
	c.HTML(http.StatusOK, "web.html",  gin.H{"hostStat": hostStats})
}

func pingerJson(c *gin.Context){

	hostStats := pingerHostStats()
	c.JSON(http.StatusOK, hostStats )

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
		pinger.GET("/json",    pingerJson)
		pinger.GET("/web", pingerWeb)
	}

	return router
}

func Start(port int){

	listen := fmt.Sprintf("0.0.0.0:%d", port)
	r := Router()
	r.Run(listen)

}
