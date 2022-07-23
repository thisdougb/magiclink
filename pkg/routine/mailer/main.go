package mailer

import (
	"encoding/json"
	"github.com/thisdougb/magiclink/config"
	"github.com/thisdougb/magiclink/pkg/datastore/redis"
	"github.com/thisdougb/magiclink/pkg/usecase/poll"
	"github.com/thisdougb/magiclink/pkg/usecase/smtpsend"
	"log"
	"time"
)

const (
	loopDelaySeconds = 3
)

// Start starts
func Poll() {

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

	pollService := poll.NewService(ds)
	smtpService := smtpsend.NewService(ds)

	for {

		time.Sleep(time.Second * loopDelaySeconds)

		nextTask := pollService.GetNextTask()
		if nextTask == "" {
			//log.Println("routine.Poll: task queue empty")
			continue
		}

		type Task struct {
			Timestamp    int64
			Email        string
			MagicLinkURL string
		}

		var task *Task = &Task{}

		err := json.Unmarshal([]byte(nextTask), task)
		if err != nil {
			log.Println("routine.Poll() queue entry is invalid:", err.Error())
			continue
		}

		log.Println("routine.Poll() sending magiclink to", task.Email)
		err = smtpService.SendMagicLink(task.Email, task.MagicLinkURL)
		if err != nil {
			log.Println(err.Error())
			// do not requeue message
		}
	}
}
