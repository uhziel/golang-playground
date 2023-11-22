package main

func main() {
	client := &Client{}

	mac := &Mac{}
	client.InsertLightningConnector2Computer(mac)

	windows := &Windows{}
	windowsAdapter := &WindowsAdapter{
		Windows: windows,
	}
	client.InsertLightningConnector2Computer(windowsAdapter)
}
