package main

import (
	"github.com/thisdougb/magiclink/api/handlers"
	"github.com/thisdougb/magiclink/config"
	"github.com/thisdougb/magiclink/pkg/datastore/redis"
	"github.com/thisdougb/magiclink/pkg/usecase/auth"
	"github.com/thisdougb/magiclink/pkg/usecase/send"
	"log"
	"net/http"
	"os"
)

func main() {

	ds := redis.NewRedisDatastore(config.DB_HOST, config.DB_PORT)

	result := ds.Connect()
	if !result {
		log.Println("Datasore connection failed, exiting...")
		os.Exit(1)
	}
	defer ds.Disconnect()

	env := &handlers.Env{
		SendService: send.NewService(ds),
		AuthService: auth.NewService(ds),
	}

	http.HandleFunc("/send/", env.Send)
	http.HandleFunc("/auth/", env.Auth)

	log.Println("webserver.Start(): listening on port", config.API_PORT)
	log.Fatal(http.ListenAndServe(":"+config.API_PORT, nil))
}
