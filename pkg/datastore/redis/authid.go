package redis

import (
	"fmt"
	"github.com/thisdougb/magiclink/config"
)

func (d *Datastore) GetExpireAccountFromID(magiclinkid string) (string, error) {

	var cfg *config.Config // dynamic config settings

	key := fmt.Sprintf("%s%s:%s", cfg.ValueAsStr("REDIS_KEY_PREFIX"), authIDsKey, magiclinkid)

	data, err := d.getExpire(key)
	if err != nil {
		if err.Error() == "redis: nil" { // not found
			return "", nil
		}
		return "", err
	}

	return data, nil
}

func (d *Datastore) StoreSessionID(email string, sessionID string, ttlSeconds int) error {

	var cfg *config.Config // dynamic config settings

	key := fmt.Sprintf("%s%s:%s", cfg.ValueAsStr("REDIS_KEY_PREFIX"), sessionIDsKey, sessionID)

	return d.setWithExpiry(key, email, ttlSeconds)

}
