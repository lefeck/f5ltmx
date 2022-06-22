package main

import (
	"f5ltmx/ltm"
)

func main() {
	client, _ := ltm.NewF5Client()
	vs := ltm.VirtualServer{}
	vs.Exec(client)
}
