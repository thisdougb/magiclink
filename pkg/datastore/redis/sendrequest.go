package redis

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/idthings/alphanum"
	"github.com/thisdougb/magiclink/config"
	"strconv"
	"time"
)

func (d *Datastore) SubmitSendLinkRequest(data string) error {

	var cfg *config.Config // dynamic config settings

	key := fmt.Sprintf("%s%s", cfg.ValueAsStr("REDIS_KEY_PREFIX"), linkSendQueue)

	// push to FIFO queue, so the notify process will send an email
	return d.lpushToList(key, data)
}

func (d *Datastore) StoreAuthID(email string, id string, ttlSeconds int) error {

	var cfg *config.Config // dynamic config settings

	// we set to the key with a ttl, so it auto-magically cleans up. only the email
	// is required.

	key := fmt.Sprintf("%s%s:%s", cfg.ValueAsStr("REDIS_KEY_PREFIX"), authIDsKey, id)

	return d.setWithExpiry(key, email, ttlSeconds)
}

func (d *Datastore) GetLoginAttempts(email string, ttlMinutes int) ([]string, error) {

	var cfg *config.Config // dynamic config settings
	var logins []string

	key := fmt.Sprintf(LoginRequestsKeyFormat, email)
	key = fmt.Sprintf("%s%s", cfg.ValueAsStr("REDIS_KEY_PREFIX"), key)

	now := time.Now().UnixNano() / 1e6 // convert to milliseconds
	nowStr := strconv.FormatInt(now, 10)

	since := now - int64((ttlMinutes * 60 * 1000)) // convert to milliseconds
	sinceStr := strconv.FormatInt(since, 10)

	// batch Redis commands
	pipe := d.client.TxPipeline()

	// housekeeping, remove old login attempts - do this here simply for general efficiency
	d.client.ZRemRangeByScore(d.ctx, key, "-inf", sinceStr)

	// An initial check for denial of service attempts. If there already too many
	// login requests don't continue. Don't add more set members, which would just
	// consume more memory.
	result := pipe.ZRangeByScore(d.ctx, key, &redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Offset: 0,
		Count:  -1,
	})

	if len(result.Val()) > cfg.ValueAsInt("RATE_LIMIT_MAX_SEND_REQUESTS") {
		return logins, errors.New("too many requests")
	}

	_, err := pipe.Exec(d.ctx)
	if err != nil {
		return logins, err
	}

	// now read the recent login attempts
	values := pipe.ZRangeByScore(d.ctx, key, &redis.ZRangeBy{
		Min:    sinceStr,
		Max:    nowStr,
		Offset: 0,
		Count:  -1,
	})

	// record this login attempt
	pipe.ZAdd(d.ctx, key, &redis.Z{
		Score:  float64(now),
		Member: alphanum.New(5), // ensure uniqueness, so every login counts
	})

	_, err = pipe.Exec(d.ctx)
	if err != nil {
		return logins, err
	}

	for _, value := range values.Val() {
		logins = append(logins, value)
	}

	return logins, nil
}
