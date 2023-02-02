package main

import (
	"douyin/routers"
	"log"
)

func main() {
	server := routers.InitRouters()
	err := server.Run(":8080")
	if err != nil {
		log.Fatalln("run server error", err)
	}

}
