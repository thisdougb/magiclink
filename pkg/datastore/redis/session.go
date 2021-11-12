package redis

import (
	"fmt"
	"github.com/thisdougb/magiclink/config"
)

func (d *Datastore) GetSessionOwner(sessionID string) (string, error) {

	var cfg *config.Config // dynamic config settings

	key := fmt.Sprintf("%s%s:%s", cfg.ValueAsStr("REDIS_KEY_PREFIX"), sessionIDsKey, sessionID)

	data, err := d.getValueAtKey(key)
	if err != nil {
		if err.Error() == "redis: nil" { // not found
			return "", nil
		}
		return "", err
	}

	return data, nil
}
