package redis

import (
	"fmt"
)

import ()

func (d *Datastore) SubmitSendLinkRequest(data string) error {

	// push to FIFO queue, so the notify process will send an email
	return d.lpushToList(linkSendQueue, data)
}

func (d *Datastore) StoreAuthID(email string, id string, ttlSeconds int) error {

	// we set to the key with a ttl, so it auto-magically cleans up. only the email
	// is required.

	key := fmt.Sprintf("%s:%s", authIDsKey, id)

	return d.setWithExpiry(key, email, ttlSeconds)
}
