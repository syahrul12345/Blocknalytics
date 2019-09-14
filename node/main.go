package main

import (
	"github.com/syahrul12345/Blocknalytics/node/modules/database"
	"github.com/syahrul12345/Blocknalytics/node/modules/server"
)

func main() {
	database.Start()
	server.Start()

}
