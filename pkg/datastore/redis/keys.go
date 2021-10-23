package redis

// define key path constants

const (
	linkSendQueue = "magiclink:queue:send" // queue of objects for external notify process to send
	authIDsKey    = "magiclink:id"         // magic link ids used for authentication
	sessionIDsKey = "magiclink:session"
)
