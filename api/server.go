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

	var cfg *config.Config // dynamic config settings

	ds := redis.NewRedisDatastore(cfg.ValueAsStr("REDIS_HOST"), cfg.ValueAsStr("REDIS_PORT"))

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

	log.Println("magiclink.Start(): listening on port", cfg.ValueAsStr("API_PORT"))
	log.Fatal(http.ListenAndServe(":"+cfg.ValueAsStr("API_PORT"), nil))
}
