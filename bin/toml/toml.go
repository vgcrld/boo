package main

import (
	_ "embed"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pelletier/go-toml/v2"
)

//go:embed samples/sample.toml
var sampleConfigFile string
var sampleFile string = "sample.toml"

type SwitchConfig struct {
	User     string
	Password string
	Servers  map[string]SwichServers
}

type SwichServers struct {
	IP   string
	Role string
}

func init() {
	if _, err := os.Stat(sampleFile); errors.Is(err, os.ErrNotExist) {
		ioutil.WriteFile(sampleFile, []byte(sampleConfigFile), 0644)
		fmt.Println("Wrote sample file: sample.toml", sampleFile)
	}
}

func main() {

	fmt.Println("\n\nFrom File:\n====================")
	cfg := tomlFromFile()
	fmt.Printf("User: %s\n", cfg.User)
	fmt.Printf("Password: %s\n", cfg.Password)
	fmt.Printf("Total Servers: %v\n", len(cfg.Servers))

	for serv, servConfig := range cfg.Servers {
		fmt.Printf("%s: IP=%s, Role=%s\n", serv, servConfig.IP, servConfig.IP)

	}
}

func tomlFromFile() SwitchConfig {
	config, err :=
		ioutil.ReadFile(sampleFile)
	if err != nil {
		panic("can read the config")
	}
	var cfg SwitchConfig
	err = toml.Unmarshal([]byte(config), &cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
