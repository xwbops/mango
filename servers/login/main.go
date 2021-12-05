package main

import (
	"xlddz/core"
	"xlddz/servers/login/business"
)

func main() {
	core.Run(
		new(business.Gate),
		new(business.Module),
	)
}
