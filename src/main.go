package main

import "MVC_DI/cmd"

func main() {
	defer cmd.Stop()
	cmd.Start()
}
