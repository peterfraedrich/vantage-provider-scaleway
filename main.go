package main

var CONFIG *Config

func main() {
	CONFIG = loadConfig("config.yaml")

}
