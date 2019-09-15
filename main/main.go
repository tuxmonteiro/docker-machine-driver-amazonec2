package main

import (
	"github.com/docker/machine/libmachine/drivers/plugin"
	"github.com/tuxmonteiro/docker-machine-driver-awsec2/awsec2"
)

func main() {
	plugin.RegisterDriver(awsec2.NewDriver("", ""))
}
