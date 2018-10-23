package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type configa struct {
	HTTPserver struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"httpserver"`
}

//Conf is configuration
var Conf configa

func init() {
	Conf.getConf()
	fmt.Println(Conf)
	// проверочный код
	fmt.Printf("Host: %#v\n", Conf.HTTPserver.Host)
	fmt.Printf("Port: %#v\n", Conf.HTTPserver.Port)
}

func main() {
	fmt.Println("Main config run: ")
	Conf.PutConf()
}

func (c *configa) getConf() *configa {
	// проверяес есть ли файл с конфигурацией
	_, errFile := os.Stat("./config/conf.yaml")
	if os.IsNotExist(errFile) {
		fmt.Printf("File configuration is not exist. Create configugation with default date.\n")
		c.HTTPserver.Host = "127.0.0.1"
		c.HTTPserver.Port = "80"
		return c
	}
	// если файл есть считіваем его
	yamlFile, err := ioutil.ReadFile("./config/conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	// парсим файл
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

func (c *configa) PutConf() bool {
	//filename, _ = filepath.Abs(filename)
	// проверяес есть ли файл с конфигурацией
	_, errFile := os.Stat("./config/conf.yaml")
	if errFile == nil {
		errFile = os.Rename("./config/conf.yaml", "./config/conf.yaml"+".bak")
	}

	yamlFile, err := yaml.Marshal(&c)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./config/conf.yaml", yamlFile, 0644)
	if err != nil {
		panic(err)
	}

	return true
}

//Config is stuct for storage data from config file
// type Config struct {
// 	HTTPserver struct {
// 		Host string `json:"host"`
// 		Port string `json:"port"`
// 	} `json:"httpserver"`
// 	Database struct {
// 		Host     string `json:"host"`
// 		Port     string `json:"port"`
// 		User     string `json:"user"`
// 		Password string `json:"password"`
// 		DBname   string `json:"dbname"`
// 	} `json:"database"`
// }

//Mconfig singleton
//var Mconfig Config
