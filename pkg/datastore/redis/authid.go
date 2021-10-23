package redis

import (
	"fmt"
)

func (d *Datastore) GetExpireAccountFromID(magiclinkid string) (string, error) {

	key := fmt.Sprintf("%s:%s", authIDsKey, magiclinkid)

	data, err := d.getExpire(key)
	if err != nil {
		return "", err
	}

	return data, nil
}

func (d *Datastore) StoreSessionID(accountID string, sessionID string, ttlSeconds int) error {

	key := fmt.Sprintf("%s:%s", sessionIDsKey, sessionID)

	return d.setWithExpiry(key, accountID, ttlSeconds)

}
