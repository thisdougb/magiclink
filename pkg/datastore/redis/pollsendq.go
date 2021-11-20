package redis

import (
	"fmt"
	"github.com/thisdougb/magiclink/config"
)

// only one pop for now
func (d *Datastore) GetNextTask() (string, error) {

	var cfg *config.Config // dynamic config settings

	key := fmt.Sprintf("%s%s", cfg.ValueAsStr("REDIS_KEY_PREFIX"), linkSendQueue)

	task, err := d.lpopFromtList(key)
	if err != nil {
		if err.Error() == "redis: nil" { // not found
			return "", nil
		}
		return "", err
	}
	return task, nil
}
