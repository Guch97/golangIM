package main

import (
	"ginchat/router"
	"ginchat/utils"
)

func main() {
	utils.InitMySQl()

	r := router.Router()

	r.Run(":8084") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
