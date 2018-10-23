package main

import (
	"fmt"

	cnf "./src/config"
	rts "./src/routers"
	//"log"
	//"net/http"
)

func init() {
	fmt.Println("Start Init func: ")
}

func main() {
	fmt.Println("Start Main func: ")
	// fmt.Println(cnf.Conf.HTTPserver.Host)
	// fmt.Println(cnf.Conf.HTTPserver.Port)
	defer cnf.Conf.PutConf()
	rts.R.info

	// err = http.ListenAndServe(Mconfig.HTTPserver.Host+":"+Mconfig.HTTPserver.Port, r)
	// if err != nil {
	// 	log.Fatal("error starting http server : ", err)
	// 	return
	// }
	//cnf.
}
