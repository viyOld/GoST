package config
// установка фремверка парсера
// go get gopkg.in/yaml.v2
import (
    "fmt"
    "io/ioutil"
    "path/filepath"
    "gopkg.in/yaml.v2"    
    "os"
//    "log"
//    "errors"
)

var cfg Config

type Config struct {
       Name     string  `yaml:"Name"`
       IP       string  `yaml:"IP"`
       Port     int     `yaml:"Port"`
       LogLevel int     `yaml:"LogLevel"`
       Language string  `yaml:"Language"` 
}

func TestConf() {
	fmt.Printf("Test Test Test\n")
}

//загружаем конфиг
func ReadConfig(filename string) {
    filename, _ = filepath.Abs(filename)
    // проверяес есть ли файл с конфигурацией
    _, errFile := os.Stat(filename)
    if os.IsNotExist(errFile) { 
    	fmt.Printf("File configuration is not exist. Create configugation with default date.\n")
    	cfg.Name = "Two"
    	cfg.IP = "127.0.0.1"
    	cfg.Port = 4321
    	cfg.LogLevel = 5
    	cfg.Language = "ua"
    	return 
    }
    // читаем файл в файловую переменную
    yamlFile, err := ioutil.ReadFile(filename)
    if err != nil {
        panic(err)
    }
    // распарсиваем файловую переменную в структуру
    err = yaml.Unmarshal(yamlFile, &cfg)
    if err != nil {
        panic(err)
    }
    // проверочный код
    fmt.Printf("Name: %#v\n", cfg.Name)
    fmt.Printf("IP: %#v\n", cfg.IP)
    fmt.Printf("Port: %#v\n", cfg.Port)
    fmt.Printf("LogLevel: %#v\n", cfg.LogLevel)
    fmt.Printf("Language: %#v\n", cfg.Language)
    fmt.Printf("******Value: %#v\n", yamlFile)
}

func  WriteConfig(filename string) {
	filename, _ = filepath.Abs(filename)
    // проверяес есть ли файл с конфигурацией
    _, errFile := os.Stat(filename)
    if errFile == nil { 
    	errFile =  os.Rename(filename, filename + ".bak")
    }

	// cfg.Port = 4321

	yamlFile, err := yaml.Marshal(&cfg)
	if err != nil {
        panic(err)
    }

    err = ioutil.WriteFile(filename, yamlFile, 0644)
    if err != nil {
        panic(err)
    }
}