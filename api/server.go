package main

import (
	"github.com/thisdougb/magiclink/api/handlers"
	"github.com/thisdougb/magiclink/config"
	"github.com/thisdougb/magiclink/pkg/datastore/redis"
	"github.com/thisdougb/magiclink/pkg/routine/mailer"
	"github.com/thisdougb/magiclink/pkg/usecase/auth"
	"github.com/thisdougb/magiclink/pkg/usecase/owner"
	"github.com/thisdougb/magiclink/pkg/usecase/send"
	"log"
	"net/http"
	"time"
)

func main() {

	var cfg *config.Config // dynamic config settings

	ds := redis.NewRedisDatastore(
		cfg.ValueAsStr("REDIS_HOST"),
		cfg.ValueAsStr("REDIS_PORT"),
		cfg.ValueAsStr("REDIS_USERNAME"),
		cfg.ValueAsStr("REDIS_PASSWORD"),
		cfg.ValueAsBool("REDIS_TLS"))

	for {
		log.Printf("Datastore connecting, host: %s:%s, username: %s\n",
			cfg.ValueAsStr("REDIS_HOST"),
			cfg.ValueAsStr("REDIS_PORT"),
			cfg.ValueAsStr("REDIS_USERNAME"))

		err := ds.Connect()
		if err == nil {
			log.Println("Datastore connected.")
			break
		}
		log.Println("Datastore connect failed:", err.Error())
		time.Sleep(5 * time.Second)
	}
	defer ds.Disconnect()

	if cfg.ValueAsBool("SMTP_ENABLED") {
		log.Println("SMTP mailer enabled")
		go mailer.Poll()

	} else {
		log.Println("SMTP mailer disabled")
	}

	env := &handlers.Env{
		SendService:  send.NewService(ds),
		AuthService:  auth.NewService(ds),
		OwnerService: owner.NewService(ds),
	}

	urlPrefix := env.GetURLPrefix()

	http.HandleFunc(urlPrefix+"/send/", env.Send)
	http.HandleFunc(urlPrefix+"/auth/", env.Auth)

	// we only expose this endpoint if required. alternative method to get session
	// owner is directly via Redis, in the caller app. by default this is turned off.
	sessionOwnerURL := cfg.ValueAsStr("SESSION_OWNER_PROTECTED_URL")
	if sessionOwnerURL != "" {
		log.Println("Adding handler for session owner endpoint:", sessionOwnerURL)
		http.HandleFunc(urlPrefix+sessionOwnerURL, env.Owner)
	}

	log.Println("listening on port", cfg.ValueAsStr("API_PORT"))
	log.Fatal(http.ListenAndServe(":"+cfg.ValueAsStr("API_PORT"), nil))
}
