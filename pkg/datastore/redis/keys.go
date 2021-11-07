package redis

// define key path constants

const (
	linkSendQueue = "queue:send" // queue of objects for external notify process to send
	authIDsKey    = "id"         // magic link ids used for authentication
	sessionIDsKey = "session"

	LoginRequestsKeyFormat = "loginRequests:%s"
)
