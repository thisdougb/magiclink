package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/thisdougb/magiclink/config"
	"log"
	"strconv"
	"time"
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

func (d *Datastore) GetLoginAttempts(email string, ttlMinutes int) ([]string, error) {

	var cfg *config.Config // dynamic config settings

	var logins []string

	key := fmt.Sprintf(LoginRequestsKeyFormat, email)
	key = fmt.Sprintf("%s%s", cfg.REDIS_KEY_PREFIX(), key)

	now := time.Now().UnixNano() / 1e6             // convert to milliseconds
	since := now - int64((ttlMinutes * 60 * 1000)) // convert to milliseconds

	// begin with clean out old login attempts
	since = since - 1
	_, err := d.client.ZRemRangeByScore(d.ctx, key, strconv.FormatInt(0, 10), strconv.FormatInt(since, 10)).Result()
	if err != nil {
		return logins, err
	}

	// now read the logins
	values, err := d.client.ZRangeByScore(d.ctx, key, &redis.ZRangeBy{
		Min:    strconv.FormatInt(since, 10),
		Max:    strconv.FormatInt(now, 10),
		Offset: 0,
		Count:  -1,
	}).Result()
	if err != nil {
		log.Println("data.GetOwnerLogins(): ", err)
		return logins, err
	}

	for _, value := range values {
		logins = append(logins, value)
	}

	return logins, nil
}

func (d *Datastore) LogLoginAttempt(email string) error {

	var cfg *config.Config // dynamic config settings

	key := fmt.Sprintf(LoginRequestsKeyFormat, email)
	key = fmt.Sprintf("%s%s", cfg.REDIS_KEY_PREFIX(), key)

	timestamp := time.Now().UnixNano() / 1e6 // convert to milliseconds

	_, err := d.client.ZAdd(d.ctx, key, &redis.Z{
		Score:  float64(timestamp),
		Member: time.Now().Format("Mon Jan _2 15:04:05 2006"),
	}).Result()
	if err != nil {
		return err
	}

	return nil
}
