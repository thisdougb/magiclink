package mailer

import (
	"encoding/json"
	"github.com/thisdougb/magiclink/config"
	"github.com/thisdougb/magiclink/pkg/datastore/redis"
	"github.com/thisdougb/magiclink/pkg/usecase/poll"
	"log"
	"os"
	"time"
)

const (
	loopDelaySeconds = 3
)

// Start starts
func Poll() {

	var cfg *config.Config // dynamic config settings

	ds := redis.NewRedisDatastore(cfg.ValueAsStr("REDIS_HOST"), cfg.ValueAsStr("REDIS_PORT"))

	result := ds.Connect()
	if !result {
		log.Println("Datasore connection failed, exiting...")
		os.Exit(1)
	}
	defer ds.Disconnect()

	pollService := poll.NewService(ds)

	//	smtpService := smtpsend.NewService(ds)

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
			log.Println("routine.Poll: queue entry is invalid:", err.Error())
			continue
		}

		log.Println("routine.Poll: Processing task for: ", task.Email)
		//err = smtpService.SendMagicLink(task.Email, task.MagicLinkURL)
		//if err != nil {
		//	log.Println(err.Error())
		// requeue message
		//}

	}
}
