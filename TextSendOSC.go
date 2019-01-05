package main

import (
	"os"
	"flag"
	"strings"
	"encoding/json"
	"io/ioutil"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/japanese"
	"github.com/hypebeast/go-osc/osc"
)

type Config struct {
	Host    string `json:"host"`
	Port    int    `json:"port"`
	Address string `json:"address"`
	IsYukariNet bool `json:"isYukariNet"`
}

func main() {
	_conf, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	var config Config
	json.Unmarshal(_conf, &config)

	flag.Parse()
	str := strings.Join(flag.Args(), " ")

	SendOSC(config, str)

	if(config.IsYukariNet){
		os.Stdout.WriteString(ConvertShiftJIS(str))
	} else {
		os.Stdout.WriteString(str)
	}
}

func SendOSC(config Config, str string) {
	client := osc.NewClient(config.Host, config.Port)
	msg := osc.NewMessage(config.Address)
	msg.Append(str)
	client.Send(msg)
}

func ConvertShiftJIS(str string) (string) {
	result, _, err := transform.String(japanese.ShiftJIS.NewEncoder(), str)
	if err != nil {
		panic(err)
	}
	return result
}