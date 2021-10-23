package redis

import (
	"fmt"
)

func (d *Datastore) GetExpireAccountFromID(magiclinkid string) (string, error) {

	key := fmt.Sprintf("%s:%s", authIDsKey, magiclinkid)

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

	key := fmt.Sprintf("%s:%s", sessionIDsKey, sessionID)

	return d.setWithExpiry(key, email, ttlSeconds)

}
