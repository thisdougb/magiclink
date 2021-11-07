package redis

import (
	"fmt"
	"github.com/thisdougb/magiclink/config"
)

func (d *Datastore) SubmitSendLinkRequest(data string) error {

	var cfg *config.Config // dynamic config settings

	key := fmt.Sprintf("%s%s", cfg.REDIS_KEY_PREFIX(), linkSendQueue)

	// push to FIFO queue, so the notify process will send an email
	return d.lpushToList(key, data)
}

func (d *Datastore) StoreAuthID(email string, id string, ttlSeconds int) error {

	var cfg *config.Config // dynamic config settings

	// we set to the key with a ttl, so it auto-magically cleans up. only the email
	// is required.

	key := fmt.Sprintf("%s%s:%s", cfg.REDIS_KEY_PREFIX(), authIDsKey, id)

	return d.setWithExpiry(key, email, ttlSeconds)
}
