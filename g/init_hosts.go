package g

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func InitHosts(hosts, hostfile string) {

	if hostfile == ""{
		InputHosts = strings.Fields( hosts )
		return
	}

	file, err := os.Open(hostfile)
	if err != nil {
		log.Fatalln("[ERROR]: ", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		h := scanner.Text()
		InputHosts = append(InputHosts, h)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln("[ERROR]: ", err)
	}

}
