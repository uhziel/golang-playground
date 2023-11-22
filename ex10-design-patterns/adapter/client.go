package main

import "fmt"

type Client struct {
}

func (c *Client) InsertLightningConnector2Computer(computer Computer) {
	fmt.Println("client inserts Lightning connector into computer")
	computer.Insert2LightningPort()
}
