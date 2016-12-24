package main

import (
	"gameproject/global/config"
)

func main() {
	config.GetConfig().Show()
	//config.ReloadConfig()
	//config.GetConfig().Show()
}
