package redis

import (
	"time"
)

func (d *Datastore) lpushToList(key string, data string) error {

	_, err := d.client.LPush(d.ctx, key, data).Result()
	if err != nil {
		return err
	}

	return nil
}

func (d *Datastore) setWithExpiry(key string, data string, ttlSeconds int) error {

	ttl := time.Duration(ttlSeconds) * time.Second

	_, err := d.client.Set(d.ctx, key, data, ttl).Result()
	if err != nil {
		return err
	}

	return nil
}
