package main

import (
	"github.com/docker/machine/libmachine/drivers/plugin"
	"github.com/tuxmonteiro/machine"
)

func main() {
	plugin.RegisterDriver(amazonec2.NewDriver("", ""))
}
