package main

import (
	//"log"
	"net/http"

	log "github.com/sirupsen/logrus"

	cnf "./src/config"
	rts "./src/routers"
)

func init() {
	log.Println("Start Init func: ")
}

func main() {
	log.Println("Start Main func: ")
	defer cnf.Conf.PutConf()
	// fmt.Println(cnf.Conf.HTTPserver.Host)
	// fmt.Println(cnf.Conf.HTTPserver.Port)

	rts.Status()

	err := http.ListenAndServe(cnf.Conf.HTTPserver.Host+":"+cnf.Conf.HTTPserver.Port, rts.R)
	if err != nil {
		log.Fatal("Error starting http server : ", err)
		return
	}

}
