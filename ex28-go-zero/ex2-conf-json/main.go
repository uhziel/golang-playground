package main

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
)

type ServerProperties struct {
	PVP bool `json:"pvp,optional"`
}

type ServerMinecraft struct {
	ServerProperties *ServerProperties `json:"serverProperties,optional"`
}

const confText = `{
  "serverProperties": {
  "pvp": true
  }
}`

func main() {
	serverMinecraft := &ServerMinecraft{}
	if err := conf.LoadFromJsonBytes([]byte(confText), serverMinecraft); err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", serverMinecraft)
	if serverMinecraft.ServerProperties.PVP == true {
		fmt.Println("pvp = true")
	}

	serverMinecraft2 := &ServerMinecraft{}
	if err := conf.LoadFromJsonBytes([]byte("{}"), serverMinecraft2); err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", serverMinecraft2)
	if serverMinecraft2.ServerProperties == nil {
		fmt.Println("ServerProperties = nil")
	}
}
